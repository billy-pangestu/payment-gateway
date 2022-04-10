package handler

import (
	"net/http"
	"payment-gateway-backend/model"
	"payment-gateway-backend/server/request"
	"payment-gateway-backend/usecase"

	validator "gopkg.in/go-playground/validator.v9"
)

// TransactionHandler ...
type TransactionHandler struct {
	Handler
}

// StoreHandler ...
func (h *TransactionHandler) StoreHandler(w http.ResponseWriter, r *http.Request) {
	api := "v1/api/payment/"

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	req := request.TransactionRequest{}

	if err := h.Handler.Bind(r, &req); err != nil {
		_, _ = usecase.HistoryUC{ContractUC: h.ContractUC}.SendToAmqp(req, api, "failed", err.Error())
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		_, _ = usecase.HistoryUC{ContractUC: h.ContractUC}.SendToAmqp(req, api, "failed", err.Error())
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	tx := model.SQLDBTx{DB: h.DB}
	dbTx, err := tx.TxBegin()
	h.ContractUC.Tx = dbTx.DB
	if err != nil {
		SendBadRequest(w, "Transaction")
		return
	}

	transactionUc := usecase.TransactionUC{ContractUC: h.ContractUC}
	res, err := transactionUc.Store(userID, req)

	if err != nil {
		_, _ = usecase.HistoryUC{ContractUC: h.ContractUC}.SendToAmqp(req, api, "failed", err.Error())
		h.ContractUC.Tx.Rollback()
		SendBadRequest(w, err.Error())
		return
	}
	_, _ = usecase.HistoryUC{ContractUC: h.ContractUC}.SendToAmqp(req, api, "successed", "")
	h.ContractUC.Tx.Commit()
	SendSuccess(w, res, nil)
	return
}
