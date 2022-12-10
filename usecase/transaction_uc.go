package usecase

import (
	"errors"
	"payment-gateway-backend/helper"
	"payment-gateway-backend/model"
	"payment-gateway-backend/pkg/logruslogger"
	"payment-gateway-backend/server/request"
	"payment-gateway-backend/usecase/viewmodel"
	"time"
)

// TransactionUC ...
type TransactionUC struct {
	*ContractUC
}

// Store ..
func (uc TransactionUC) Store(UserID string, data request.TransactionRequest) (res viewmodel.TransactionVM, err error) {
	ctx := "TransactionUC.Store"

	userData, err := UserUC{ContractUC: uc.ContractUC}.FindByID(UserID, false)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_not_found", uc.ReqID)
		return res, errors.New(helper.UserNotFound)
	}
	if userData.Amount < data.Amount {
		return res, errors.New("Insufficient Fund")
	}

	now := time.Now().UTC()
	res = viewmodel.TransactionVM{
		UserID:    UserID,
		MerchID:   data.MerchandID,
		Amount:    data.Amount,
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}

	_, err = MerchantUC{ContractUC: uc.ContractUC}.AddFund(data.MerchandID, res.Amount)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "store_to_merchant", uc.ReqID)
		return res, err
	}

	_, err = UserUC{ContractUC: uc.ContractUC}.SubFund(UserID, res.Amount)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "store_to_merchant", uc.ReqID)
		return res, err
	}

	transactionModel := model.NewTransactionModel(uc.Tx)
	res.ID, err = transactionModel.Store(res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	return res, err
}
