package usecase

import (
	"errors"
	"payment-gateway-backend/model"
	"payment-gateway-backend/pkg/logruslogger"
	"payment-gateway-backend/server/request"
	"payment-gateway-backend/usecase/viewmodel"
	"time"
)

// MerchantUC ...
type MerchantUC struct {
	*ContractUC
}

// BuildBody ...
func (uc MerchantUC) BuildBody(data *model.MerchantEntity, res *viewmodel.MerchantVM, showAmount bool) {
	res.ID = data.ID
	res.Name = data.Name.String
	res.CreatedAt = data.CreatedAt.String
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String

	if showAmount {
		res.Amount = data.Amount
	}
}

// FindByAll ...
func (uc MerchantUC) FindAll(page, limit int, showAmount bool) (res []viewmodel.MerchantVM, pagination viewmodel.PaginationVM, err error) {
	ctx := "MerchantUC.FindByID"

	limit = uc.LimitMax(limit)
	limit, offset := uc.PaginationPageOffset(page, limit)

	merchantModel := model.NewMerchantModel(uc.Tx)
	data, count, err := merchantModel.FindAll(offset, limit)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, pagination, err
	}

	pagination = PaginationRes(page, count, limit)

	for _, r := range data {
		temp := viewmodel.MerchantVM{}
		uc.BuildBody(&r, &temp, showAmount)

		res = append(res, temp)
	}

	return res, pagination, err
}

// FindByID ...
func (uc MerchantUC) FindByID(id string, showAmount bool) (res viewmodel.MerchantVM, err error) {
	ctx := "MerchantUC.FindByID"

	merchantModel := model.NewMerchantModel(uc.Tx)
	data, err := merchantModel.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&data, &res, showAmount)

	return res, err
}

// Create ...
func (uc MerchantUC) Create(data request.MerchantRequest) (res viewmodel.MerchantVM, err error) {
	ctx := "MerchantUC.Create"

	now := time.Now().UTC()
	res = viewmodel.MerchantVM{
		Name: data.Name,
	}
	merchantModel := model.NewMerchantModel(uc.Tx)
	res.ID, err = merchantModel.Store(res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_store_agent", uc.ReqID)
		return res, err
	}

	return res, err
}

// AddFund ...
func (uc MerchantUC) AddFund(merchantID string, amount float64) (res viewmodel.MerchantVM, err error) {
	ctx := "MerchantUC.AddFund"

	_, err = uc.FindByID(merchantID, false)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "metchant_not_found", uc.ReqID)
		return res, errors.New("merchant_id_not_found")
	}

	now := time.Now().UTC()
	res = viewmodel.MerchantVM{
		Amount: amount,
	}
	merchantModel := model.NewMerchantModel(uc.Tx)
	res.ID, err = merchantModel.AddFund(merchantID, res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_update", uc.ReqID)
		return res, err
	}

	return res, err
}
