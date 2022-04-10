package handler

// TransactionHandler ...
type TransactionHandler struct {
	Handler
}

// // GetByIDHandler ...
// func (h *TransactionHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	user := requestIDFromContextInterface(r.Context(), "user")
// 	userID := user["id"].(string)

// 	id := r.URL.Query().Get("id")

// 	if id == "" {
// 		SendBadRequest(w, "Parameter must be filled")
// 		return
// 	}

// 	// Get logrus request ID
// 	h.ContractUC.ReqID = getHeaderReqID(r)

// 	transactionUc := usecase.TransactionUC{ContractUC: h.ContractUC}
// 	res, err := transactionUc.FindByID(id, userID)

// 	if err != nil {
// 		SendBadRequest(w, err.Error())
// 		return
// 	}

// 	SendSuccess(w, res, nil)
// 	return
// }

// // GetByTagHandler ...
// func (h *TransactionHandler) GetByTagHandler(w http.ResponseWriter, r *http.Request) {
// 	user := requestIDFromContextInterface(r.Context(), "user")
// 	userID := user["id"].(string)

// 	tag := r.URL.Query().Get("tag")

// 	if tag == "" {
// 		SendBadRequest(w, "Parameter must be filled")
// 		return
// 	}

// 	// Get logrus request ID
// 	h.ContractUC.ReqID = getHeaderReqID(r)

// 	transactionUc := usecase.TransactionUC{ContractUC: h.ContractUC}

// 	res, err := transactionUc.FindByTag(userID, tag)

// 	if err != nil {
// 		SendBadRequest(w, err.Error())
// 		return
// 	}

// 	SendSuccess(w, res, nil)
// 	return
// }

// // StoreHandler ...
// func (h *TransactionHandler) StoreHandler(w http.ResponseWriter, r *http.Request) {
// 	user := requestIDFromContextInterface(r.Context(), "user")
// 	userID := user["id"].(string)

// 	req := request.TransactionRequest{}

// 	if err := h.Handler.Bind(r, &req); err != nil {

// 		SendBadRequest(w, err.Error())
// 		return
// 	}
// 	if err := h.Handler.Validate.Struct(req); err != nil {

// 		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
// 		return
// 	}

// 	// Get logrus request ID
// 	h.ContractUC.ReqID = getHeaderReqID(r)

// 	transactionUc := usecase.TransactionUC{ContractUC: h.ContractUC}

// 	tx, err := transactionUc.DB.Begin()
// 	if err != nil {
// 		SendBadRequest(w, err.Error())
// 	}

// 	res, err := transactionUc.Store(userID, req, tx)

// 	if err != nil {
// 		tx.Rollback()
// 		SendBadRequest(w, err.Error())
// 		return
// 	}

// 	tx.Commit()

// 	SendSuccess(w, res, nil)
// 	return
// }

// // GetByUserIDHandler ...
// func (h *TransactionHandler) GetByUserIDHandler(w http.ResponseWriter, r *http.Request) {
// 	user := requestIDFromContextInterface(r.Context(), "user")
// 	userID := user["id"].(string)

// 	// Get logrus request ID
// 	h.ContractUC.ReqID = getHeaderReqID(r)

// 	transactionUc := usecase.TransactionUC{ContractUC: h.ContractUC}
// 	res, err := transactionUc.FindByUserID(userID)
// 	if err != nil {
// 		SendBadRequest(w, err.Error())
// 		return
// 	}

// 	SendSuccess(w, res, nil)
// 	return
// }

// // GetTotalAmountHandler ...
// func (h *TransactionHandler) GetTotalAmountHandler(w http.ResponseWriter, r *http.Request) {
// 	user := requestIDFromContextInterface(r.Context(), "user")
// 	userID := user["id"].(string)

// 	// Get logrus request ID
// 	h.ContractUC.ReqID = getHeaderReqID(r)

// 	transactionUc := usecase.TransactionUC{ContractUC: h.ContractUC}
// 	res, err := transactionUc.FindTotalAmount(userID)

// 	if err != nil {
// 		SendBadRequest(w, err.Error())
// 		return
// 	}

// 	SendSuccess(w, res, nil)
// 	return
// }

// // UpdateHandler ...
// func (h *TransactionHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
// 	user := requestIDFromContextInterface(r.Context(), "user")
// 	userID := user["id"].(string)

// 	id := r.URL.Query().Get("id")

// 	if id == "" {
// 		SendBadRequest(w, "Parameter must be filled")
// 		return
// 	}

// 	req := request.TransactionRequest{}

// 	if err := h.Handler.Bind(r, &req); err != nil {

// 		SendBadRequest(w, err.Error())
// 		return
// 	}
// 	if err := h.Handler.Validate.Struct(req); err != nil {

// 		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
// 		return
// 	}

// 	// Get logrus request ID
// 	h.ContractUC.ReqID = getHeaderReqID(r)

// 	transactionUc := usecase.TransactionUC{ContractUC: h.ContractUC}

// 	tx, err := transactionUc.DB.Begin()
// 	if err != nil {
// 		SendBadRequest(w, err.Error())
// 	}

// 	res, err := transactionUc.Update(id, userID, req, tx)

// 	if err != nil {
// 		tx.Rollback()
// 		SendBadRequest(w, err.Error())
// 		return
// 	}

// 	tx.Commit()

// 	SendSuccess(w, res, nil)
// 	return
// }

// // DeleteHandler ...
// func (h *TransactionHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
// 	user := requestIDFromContextInterface(r.Context(), "user")
// 	userID := user["id"].(string)

// 	id := r.URL.Query().Get("id")

// 	if id == "" {
// 		SendBadRequest(w, "Parameter must be filled")
// 		return
// 	}

// 	// Get logrus request ID
// 	h.ContractUC.ReqID = getHeaderReqID(r)

// 	transactionUc := usecase.TransactionUC{ContractUC: h.ContractUC}
// 	res, err := transactionUc.Delete(id, userID)

// 	if err != nil {
// 		SendBadRequest(w, err.Error())
// 		return
// 	}

// 	SendSuccess(w, res, nil)
// 	return
// }
