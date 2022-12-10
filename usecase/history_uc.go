package usecase

import (
	"errors"
	"payment-gateway-backend/helper"
	"payment-gateway-backend/model"
	"payment-gateway-backend/pkg/amqp"
	"payment-gateway-backend/pkg/logruslogger"
	"time"
)

// HistoryUC ...
type HistoryUC struct {
	*ContractUC
}

// Store ...
func (uc HistoryUC) Store(qid, payload, api, status, errorstring string) (err error) {
	ctx := "HistoryUC.Store"

	roleModel := model.NewHistoryModel(uc.Tx)
	now := time.Now().UTC()
	_, err = roleModel.Store(qid, payload, api, status, errorstring, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return err
	}

	return err
}

func (uc HistoryUC) SendToAmqp(payload interface{}, api, status, errorString string) (res string, err error) {
	ctx := "HistoryUC.SendToAmqp"

	mqueue := amqp.NewQueue(AmqpConnection, AmqpChannel)
	queueBody := map[string]interface{}{
		"qid":          uc.ContractUC.ReqID,
		"payload":      payload,
		"api":          api,
		"status":       status,
		"error_string": errorString,
	}
	AmqpConnection, AmqpChannel, err = mqueue.PushQueueReconnect(uc.ContractUC.EnvConfig["AMQP_URL"], queueBody, amqp.History, amqp.HistoryDeadLetter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "amqp", uc.ReqID)
		return res, errors.New(helper.InternalServer)
	}

	return res, err
}
