package usecase

import (
	"encoding/json"
	"errors"
	"payment-gateway-backend/pkg/aesfront"
	"payment-gateway-backend/pkg/logruslogger"
	"time"

	"database/sql"
	"payment-gateway-backend/pkg/aes"
	"payment-gateway-backend/pkg/jwe"
	"payment-gateway-backend/pkg/jwt"
	"payment-gateway-backend/pkg/mandrill"
	"payment-gateway-backend/pkg/wavecell"
	"payment-gateway-backend/usecase/viewmodel"

	"github.com/go-redis/redis/v7"
	"github.com/minio/minio-go/v6"
	"github.com/streadway/amqp"
)

const (
	// DefaultLimit ...
	DefaultLimit = 10
	// MaxLimit ...
	MaxLimit = 50
	// DefaultLocation ...
	DefaultLocation = "Asia/Jakarta"
	// DefaultTimezone ...
	DefaultTimezone = "+07:00"
	// MaxAPISendOTP ...
	MaxAPISendOTP = 10
	// MaxOtpSubmitRetry ...
	MaxOtpSubmitRetry = 3
	// OtpLifetime ...
	OtpLifetime = "2m"
	// LoginAction ...
	LoginAction = "login"
	//RegisterAction ...
	RegisterAction = "register"
	//SetPinAction ...
	SetPinAction = "set_pin"
	//VerifyOtpAction ...
	VerifyOtpAction = "verify_otp"
)

// AmqpConnection ...
var AmqpConnection *amqp.Connection

// AmqpChannel ...
var AmqpChannel *amqp.Channel

// ContractUC ...
type ContractUC struct {
	ReqID       string
	DB          *sql.DB
	Tx          *sql.Tx
	AmqpConn    *amqp.Connection
	AmqpChannel *amqp.Channel
	Redis       *redis.Client
	EnvConfig   map[string]string
	Jwt         jwt.Credential
	Jwe         jwe.Credential
	Aes         aes.Credential
	AesFront    aesfront.Credential
	Wavecell    wavecell.Connection
	Mandrill    mandrill.Credential
	Minio       *minio.Client
}

// StoreToRedis save data to redis with key key
func (uc ContractUC) StoreToRedis(key string, val interface{}) error {
	ctx := "ContractUC.StoreToRedis"

	b, err := json.Marshal(val)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "json_marshal", uc.ReqID)
		return err
	}

	err = uc.Redis.Set(key, string(b), 0).Err()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_set", uc.ReqID)
		return err
	}

	return err
}

// StoreToRedisExp save data to redis with key and exp time
func (uc ContractUC) StoreToRedisExp(key string, val interface{}, duration string) error {
	ctx := "ContractUC.StoreToRedisExp"

	dur, err := time.ParseDuration(duration)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "parse_duration", uc.ReqID)
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "json_marshal", uc.ReqID)
		return err
	}

	err = uc.Redis.Set(key, string(b), dur).Err()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_set", uc.ReqID)
		return err
	}

	return err
}

// GetFromRedis get value from redis by key
func (uc ContractUC) GetFromRedis(key string, cb interface{}) error {
	ctx := "ContractUC.GetFromRedis"

	res, err := uc.Redis.Get(key).Result()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_get", uc.ReqID)
		return err
	}

	if res == "" {
		logruslogger.Log(logruslogger.WarnLevel, "", ctx, "redis_empty", uc.ReqID)
		return errors.New("[Redis] Value of " + key + " is empty.")
	}

	err = json.Unmarshal([]byte(res), &cb)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "json_unmarshal", uc.ReqID)
		return err
	}

	return err
}

// GetAllStringFromRedis get all value from redis by key
func (uc ContractUC) GetAllStringFromRedis(key string) (res []viewmodel.RedisStringValueVM, err error) {
	ctx := "ContractUC.GetAllStringFromRedis"

	keyList, err := uc.Redis.Keys(key).Result()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_get_all", uc.ReqID)
		return res, err
	}

	for _, k := range keyList {
		var value string
		uc.GetFromRedis(k, &value)

		if value != "" {
			data := viewmodel.RedisStringValueVM{
				Key:   k,
				Value: value,
			}
			res = append(res, data)
		}
	}

	return res, err
}

// RemoveFromRedis remove a key from redis
func (uc ContractUC) RemoveFromRedis(key string) error {
	ctx := "ContractUC.RemoveFromRedis"

	err := uc.Redis.Del(key).Err()
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_delete", uc.ReqID)
		return err
	}

	return err
}

// PaginationPageOffset Calculate offset and limit by inputed page and limit
func (uc ContractUC) PaginationPageOffset(page, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 || limit > MaxLimit {
		limit = DefaultLimit
	}
	offset := (page - 1) * limit

	return limit, offset
}

// LimitMax Set limit to default if value > max
func (uc ContractUC) LimitMax(limit int) int {
	if limit <= 0 || limit > MaxLimit {
		limit = DefaultLimit
	}

	return limit
}

// PaginationRes pagination viewmodel helper
func PaginationRes(page, count, limit int) viewmodel.PaginationVM {
	lastPage := count / limit
	if count%limit > 0 {
		lastPage = lastPage + 1
	}

	pagination := viewmodel.PaginationVM{
		CurrentPage:   page,
		LastPage:      lastPage,
		Count:         count,
		RecordPerPage: limit,
	}
	return pagination
}

// AddCounterRedis add counter to redis by key
func (uc ContractUC) AddCounterRedis(key, duration string) (err error) {
	ctx := "ContractUC.AddCounterRedis"

	var (
		cb      interface{}
		counter float64 = 1
	)
	err = uc.GetFromRedis(key, &cb)
	// If exist add up the counter
	if err == nil {
		counter = cb.(float64) + 1
	}

	err = uc.StoreToRedisExp(key, counter, duration)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "redis_set", uc.ReqID)
		return err
	}

	return err
}

// LimitByKey global function to limit action using redis
func (uc ContractUC) LimitByKey(key string, limit float64, errMsg string) (err error) {
	var count float64
	resRedis := map[string]interface{}{}
	err = uc.GetFromRedis(key, &resRedis)
	if err != nil {
		err = nil
		resRedis = map[string]interface{}{
			"count": count,
		}
	}

	count = resRedis["count"].(float64) + 1
	if count > limit {
		return errors.New(errMsg)
	}

	resRedis["count"] = count
	uc.StoreToRedisExp(key, resRedis, "24h")

	return err
}

// ResetByKey global function to reset counter using redis
func (uc ContractUC) ResetByKey(key string) (err error) {
	var count float64
	resRedis := map[string]interface{}{
		"count": count,
	}
	uc.StoreToRedisExp(key, resRedis, "24h")

	return err
}
