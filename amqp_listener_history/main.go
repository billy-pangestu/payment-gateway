package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"payment-gateway-backend/model"
	"payment-gateway-backend/pkg/interfacepkg"
	"runtime"

	amqpPkg "payment-gateway-backend/pkg/amqp"
	"payment-gateway-backend/pkg/amqpconsumer"
	"payment-gateway-backend/pkg/env"
	"payment-gateway-backend/pkg/logruslogger"
	"payment-gateway-backend/pkg/pg"
	"payment-gateway-backend/pkg/str"
	"payment-gateway-backend/usecase"

	"github.com/streadway/amqp"
)

var (
	uri          *string
	formURL      = flag.String("form_url", "http://localhost", "The URL that requests are sent to")
	logFile      = flag.String("log_file", "system.log", "The file where errors are logged")
	threads      = flag.Int("threads", 1, "The max amount of go routines that you would like the process to use")
	maxprocs     = flag.Int("max_procs", 1, "The max amount of processors that your application should use")
	paymentsKey  = flag.String("payments_key", "secret", "Access key")
	exchange     = flag.String("exchange", amqpPkg.HistoryExchange, "The exchange we will be binding to")
	exchangeType = flag.String("exchange_type", "direct", "Type of exchange we are binding to | topic | direct| etc..")
	queue        = flag.String("queue", amqpPkg.History, "Name of the queue that you would like to connect to")
	routingKey   = flag.String("routing_key", amqpPkg.HistoryDeadLetter, "queue to route messages to")
	workerName   = flag.String("worker_name", "worker.name", "name to identify worker by")
	verbosity    = flag.Bool("verbos", false, "Set true if you would like to log EVERYTHING")

	// Hold consumer so our go routine can listen to
	// it's done error chan and trigger reconnects
	// if it's ever returned
	conn      *amqpconsumer.Consumer
	envConfig map[string]string
)

func init() {
	flag.Parse()
	runtime.GOMAXPROCS(*maxprocs)
	envConfig = env.NewEnvConfig("../.env")
	uri = flag.String("uri", envConfig["AMQP_URL"], "The rabbitmq endpoint")
}

func main() {
	file := false
	// Open a system file to start logging to
	if file {
		f, err := os.OpenFile(*logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		defer f.Close()
		if err != nil {
			log.Printf("error opening file: %v", err.Error())
		}
		log.SetOutput(f)
	}

	conn := amqpconsumer.NewConsumer(*workerName, *uri, *exchange, *exchangeType, *queue)

	if err := conn.Connect(); err != nil {
		log.Printf("Error: %v", err)
	}

	deliveries, err := conn.AnnounceQueue(*queue, *routingKey)
	if err != nil {
		log.Printf("Error when calling AnnounceQueue(): %v", err.Error())
	}

	// setup database connection
	dbInfo := pg.Connection{
		Host:    envConfig["DATABASE_HOST"],
		DB:      envConfig["DATABASE_DB"],
		User:    envConfig["DATABASE_USER"],
		Pass:    envConfig["DATABASE_PASSWORD"],
		Port:    str.StringToInt(envConfig["DATABASE_PORT"]),
		SslMode: envConfig["DATABASE_SSL_MODE"],
	}
	db, err := dbInfo.Connect()
	if err != nil {
		handleError(err, "Can't connect to DB")
	}
	defer db.Close()

	cUC := usecase.ContractUC{
		DB:        db,
		EnvConfig: envConfig,
	}

	// Create static folder for file uploading
	filePath := envConfig["FILE_STATIC_FILE"]
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.Mkdir(filePath, os.ModePerm)
	}

	conn.Handle(deliveries, handler, *threads, *queue, *routingKey, cUC)
}

func handler(deliveries <-chan amqp.Delivery, uc *usecase.ContractUC) {
	ctx := "ReplyListener"

	for d := range deliveries {
		var formData map[string]interface{}

		err := json.Unmarshal(d.Body, &formData)
		if err != nil {
			log.Printf("Error unmarshaling data: %s", err.Error())
		}

		logruslogger.Log(logruslogger.InfoLevel, interfacepkg.Marshall(formData), ctx, "begin", formData["qid"].(string))

		uc.ReqID = formData["qid"].(string)
		tx := model.SQLDBTx{DB: uc.DB}
		dbTx, err := tx.TxBegin()
		uc.Tx = dbTx.DB
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "transaction", formData["qid"].(string))
			d.Reject(false)
		}

		// Get Data
		var payload interface{}
		interfacepkg.UnmarshallCbInterface(formData["payload"].(interface{}), &payload)
		payloadBody, _ := json.Marshal(payload)

		api := formData["api"].(string)
		status := formData["status"].(string)
		errorString := formData["error_string"].(string)

		broadcastUc := usecase.HistoryUC{ContractUC: uc}
		err = broadcastUc.Store(formData["qid"].(string), string(payloadBody), api, status, errorString)
		if err != nil {
			uc.Tx.Rollback()
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "err", formData["qid"].(string))

		} else {
			uc.Tx.Commit()
			logruslogger.Log(logruslogger.InfoLevel, string(d.Body), ctx, "success", formData["qid"].(string))
			d.Ack(false)
		}
	}

	return
}

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
