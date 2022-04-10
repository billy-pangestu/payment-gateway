package main

import (
	"encoding/json"
	"flag"
	"go-halobro-backend/model"
	"go-halobro-backend/pkg/dialog360"
	"go-halobro-backend/pkg/dialogflow"
	"go-halobro-backend/pkg/facebook"
	"go-halobro-backend/pkg/gmail"
	"go-halobro-backend/pkg/interfacepkg"
	"go-halobro-backend/pkg/line"
	"go-halobro-backend/pkg/nexmo"
	"go-halobro-backend/pkg/pusher"
	"go-halobro-backend/pkg/telegram"
	"log"
	"os"
	"runtime"
	"strconv"

	amqpPkg "go-halobro-backend/pkg/amqp"
	"go-halobro-backend/pkg/amqpconsumer"
	"go-halobro-backend/pkg/env"
	"go-halobro-backend/pkg/logruslogger"
	"go-halobro-backend/pkg/minio"
	"go-halobro-backend/pkg/pg"
	"go-halobro-backend/pkg/str"
	"go-halobro-backend/usecase"
	"go-halobro-backend/usecase/viewmodel"

	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

var (
	uri          *string
	formURL      = flag.String("form_url", "http://localhost", "The URL that requests are sent to")
	logFile      = flag.String("log_file", "system.log", "The file where errors are logged")
	threads      = flag.Int("threads", 1, "The max amount of go routines that you would like the process to use")
	maxprocs     = flag.Int("max_procs", 1, "The max amount of processors that your application should use")
	paymentsKey  = flag.String("payments_key", "secret", "Access key")
	exchange     = flag.String("exchange", amqpPkg.BroadcastExchange, "The exchange we will be binding to")
	exchangeType = flag.String("exchange_type", "direct", "Type of exchange we are binding to | topic | direct| etc..")
	queue        = flag.String("queue", amqpPkg.Broadcast, "Name of the queue that you would like to connect to")
	routingKey   = flag.String("routing_key", amqpPkg.BroadcastDeadLetter, "queue to route messages to")
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

	// Setup redis connection
	redisClient := redis.NewClient(&redis.Options{
		Addr:     envConfig["REDIS_HOST"],
		Password: envConfig["REDIS_PASSWORD"],
		DB:       0,
	})
	_, err = redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}

	// setup database connection
	dbInfo := pg.Connection{
		Host:            envConfig["DATABASE_HOST"],
		DB:              envConfig["DATABASE_DB"],
		User:            envConfig["DATABASE_USER"],
		Pass:            envConfig["DATABASE_PASSWORD"],
		Port:            str.StringToInt(envConfig["DATABASE_PORT"]),
		SslMode:         envConfig["DATABASE_SSL_MODE"],
		MaxIdleConns:    str.StringToInt(envConfig["DATABASE_MAX_IDLE_CONN"]),
		MaxOpenConns:    str.StringToInt(envConfig["DATABASE_MAX_OPEN_CONN"]),
		ConnMaxLifetime: str.StringToInt(envConfig["DATABASE_MAX_LIFETIME"]),
	}
	db, err := dbInfo.Connect()
	if err != nil {
		handleError(err, "Can't connect to DB")
	}
	defer db.Close()

	// Minio connection
	minioInfo := minio.Connection{
		Endpoint:        envConfig["MINIO_ENDPOINT"],
		AccessKeyID:     envConfig["MINIO_ACCESS_KEY_ID"],
		SecretAccessKey: envConfig["MINIO_SECRET_ACCESS_KEY"],
		UseSSL:          str.StringToBool(envConfig["MINIO_USE_SSL"]),
		BaseURL:         envConfig["MINIO_BASE_URL"],
		DefaultBucket:   envConfig["MINIO_DEFAULT_BUCKET"],
	}
	minioConn, err := minioInfo.Connect()
	if err != nil {
		panic(err)
	}

	// telegramCredential
	telegramCredential := telegram.Credential{
		BaseURL: envConfig["TELEGRAM_URL"],
	}

	// facebookCredential
	facebookCredential := facebook.Credential{
		BaseURL: envConfig["FACEBOOK_URL"],
	}

	// lineCredential
	lineCredential := line.Credential{
		BaseURL: envConfig["LINE_URL"],
		Path:    envConfig["FILE_STATIC_FILE"],
	}

	// gmailCredential
	gmailCredential := gmail.Credential{
		UserID:     envConfig["GMAIL_USER_ID"],
		GmailKey:   envConfig["GMAIL_CREDENTIAL_KEY"],
		GmailPath:  envConfig["GMAIL_CREDENTIAL_URL"],
		GmailToken: envConfig["GMAIL_TOKEN"],
		BaseURL:    envConfig["GMAIL_URL"],
	}

	// Pusher
	pusher := pusher.Credential{
		AppID:   envConfig["PUSHER_APP_ID"],
		Key:     envConfig["PUSHER_KEY"],
		Secret:  envConfig["PUSHER_SECRET"],
		Cluster: envConfig["PUSHER_CLUSTER"],
	}

	// Nexmo
	nexmoCredential := nexmo.Credential{
		AppID:          envConfig["NEXMO_APP_ID"],
		AppKey:         envConfig["NEXMO_APP_KEY"],
		AppSecret:      envConfig["NEXMO_APP_SECRET"],
		Key:            envConfig["NEXMO_KEY"],
		BaseURL:        envConfig["NEXMO_BASE_URL"],
		SandboxBaseURL: envConfig["NEXMO_SANDBOX_BASE_URL"],
	}

	// Dialog360
	dialog360Credential := dialog360.Credential{
		BaseURL: envConfig["DIALOG360_BASE_URL"],
		Path:    envConfig["FILE_STATIC_FILE"],
	}

	// Dialogflow
	dialogflowCredential := dialogflow.Credential{
		ProjectID: envConfig["GOOGLE_DIALOGFLOW_PROJECT_ID"],
		Language:  envConfig["GOOGLE_DIALOGFLOW_LANGUAGE"],
	}

	cUC := usecase.ContractUC{
		DB:         db,
		Redis:      redisClient,
		Minio:      minioConn,
		EnvConfig:  envConfig,
		Pusher:     pusher,
		Telegram:   telegramCredential,
		Facebook:   facebookCredential,
		Line:       lineCredential,
		Gmail:      gmailCredential,
		Nexmo:      nexmoCredential,
		Dialog360:  dialog360Credential,
		Dialogflow: dialogflowCredential,
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
		var channel viewmodel.ChannelVM
		interfacepkg.UnmarshallCbInterface(formData["channel"].(interface{}), &channel)
		var client viewmodel.ClientVM
		interfacepkg.UnmarshallCbInterface(formData["client"].(interface{}), &client)
		var template viewmodel.TemplateVM
		interfacepkg.UnmarshallCbInterface(formData["template"].(interface{}), &template)
		userID := formData["userid"].(string)
		userEmail := formData["useremail"].(string)
		message := formData["message"].(string)

		broadcastUc := usecase.BroadcastUC{ContractUC: uc}
		err = broadcastUc.ListenerBroadcast(&channel, &client, &template, userID, userEmail, message)
		if err != nil {
			uc.Tx.Rollback()
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "err", formData["qid"].(string))

			// Get fail counter from redis
			failCounter := amqpconsumer.FailCounter{}
			err = uc.GetFromRedis("amqpFail"+formData["qid"].(string), &failCounter)
			if err != nil {
				failCounter = amqpconsumer.FailCounter{
					Counter: 1,
				}
			}

			if failCounter.Counter > amqpconsumer.MaxFailCounter {
				logruslogger.Log(logruslogger.WarnLevel, strconv.Itoa(failCounter.Counter), ctx, "rejected", formData["qid"].(string))
				d.Reject(false)
			} else {
				// Save the new counter to redis
				failCounter.Counter++
				err = uc.StoreToRedisExp("amqpFail"+formData["qid"].(string), failCounter, "10m")

				logruslogger.Log(logruslogger.WarnLevel, strconv.Itoa(failCounter.Counter), ctx, "failed", formData["qid"].(string))
				d.Nack(false, true)
			}
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
