package handler

import (
	"net/http"
	"strconv"

	"payment-gateway-backend/model"
	"payment-gateway-backend/usecase"
)

// MerchantHandler ...
type MerchantHandler struct {
	Handler
}

// FindAllHandler ...
func (h *MerchantHandler) FindAllHandler(w http.ResponseWriter, r *http.Request) {
	api := "v1/api/merchant"

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		SendBadRequest(w, "Invalid page value")
		return
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		SendBadRequest(w, "Invalid limit value")
		return
	}

	tx := model.SQLDBTx{DB: h.DB}
	dbTx, err := tx.TxBegin()
	h.ContractUC.Tx = dbTx.DB
	if err != nil {
		_, _ = usecase.HistoryUC{ContractUC: h.ContractUC}.SendToAmqp("", api, "failed", err.Error())
		SendBadRequest(w, "Transaction")
		return
	}

	merchantUc := usecase.MerchantUC{ContractUC: h.ContractUC}
	res, p, err := merchantUc.FindAll(page, limit, false)
	if err != nil {
		_, _ = usecase.HistoryUC{ContractUC: h.ContractUC}.SendToAmqp("", api, "failed", err.Error())
		h.ContractUC.Tx.Rollback()
		SendBadRequest(w, err.Error())
		return
	}
	_, _ = usecase.HistoryUC{ContractUC: h.ContractUC}.SendToAmqp("", api, "successed", "")
	h.ContractUC.Tx.Commit()

	SendSuccess(w, res, p)
	return
}
