package usecase

// // TransactionUC ...
// type TransactionUC struct {
// 	*ContractUC
// }

// // FindByID ...
// func (uc TransactionUC) FindByID(id, userID string) (res viewmodel.TransactionVM, err error) {
// 	ctx := "TransactionUC.FindByID"

// 	transactionModel := model.NewTransactionModel(uc.DB)
// 	transaction, err := transactionModel.FindByID(id, userID)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}
// 	tagName := str.Unique(strings.Split(transaction.Tags.String, ","))
// 	tagTemp := []viewmodel.DetailTagVM{}
// 	for i := range tagName {
// 		if tagName[i] != "" {
// 			tagTemp = append(tagTemp, viewmodel.DetailTagVM{
// 				Hashtag: tagName[i],
// 			})
// 		}
// 	}

// 	res = viewmodel.TransactionVM{
// 		ID:        transaction.ID,
// 		UserID:    transaction.UserID,
// 		MoneyIn:   transaction.MoneyIn.Float64,
// 		MoneyOut:  transaction.MoneyOut.Float64,
// 		Notes:     transaction.Notes,
// 		Hashtag:   tagTemp,
// 		CreatedAt: transaction.CreatedAt,
// 		UpdatedAt: transaction.UpdatedAt.String,
// 		DeletedAt: transaction.DeletedAt.String,
// 	}

// 	return res, err
// }

// // FindByUserID ...
// func (uc TransactionUC) FindByUserID(userID string) (res viewmodel.TransactionWithTotalAmount, err error) {
// 	ctx := "TransactionUC.FindByUserID"

// 	transactionModel := model.NewTransactionModel(uc.DB)

// 	TotalAmount, err := uc.FindTotalAmount(userID)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	transaction, err := transactionModel.FindByUserID(userID)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	var datasTrans []viewmodel.TransactionVM
// 	for i, data := range transaction {

// 		tagName := str.Unique(strings.Split(transaction[i].Tags.String, ","))
// 		tagTemp := []viewmodel.DetailTagVM{}
// 		for i := range tagName {
// 			if tagName[i] != "" {
// 				tagTemp = append(tagTemp, viewmodel.DetailTagVM{
// 					Hashtag: tagName[i],
// 				})
// 			}
// 		}

// 		datasTrans = append(datasTrans, viewmodel.TransactionVM{
// 			ID:        data.ID,
// 			UserID:    data.UserID,
// 			MoneyIn:   data.MoneyIn.Float64,
// 			MoneyOut:  data.MoneyOut.Float64,
// 			Notes:     data.Notes,
// 			Hashtag:   tagTemp,
// 			CreatedAt: data.CreatedAt,
// 			UpdatedAt: data.UpdatedAt.String,
// 		})
// 	}

// 	res = viewmodel.TransactionWithTotalAmount{
// 		TotalMoneyIn:  TotalAmount.MoneyIn,
// 		TotalMoneyOut: TotalAmount.MoneyOut,
// 		Transactions:  datasTrans,
// 	}

// 	return res, err
// }

// // FindByTag ...
// func (uc TransactionUC) FindByTag(userID, tag string) (res viewmodel.TransactionWithTotalAmount, err error) {
// 	ctx := "TransactionUC.FindByTag"

// 	transactionModel := model.NewTransactionModel(uc.DB)

// 	transaction, err := transactionModel.FindByTag(userID, tag)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	totalAmount, err := transactionModel.FindTotalMoneyByTag(userID, tag)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	tagUC := TagUC{ContractUC: uc.ContractUC}
// 	var datasTrans []viewmodel.TransactionVM
// 	for _, data := range transaction {

// 		tags, err := tagUC.FindByTransactionID(data.ID)
// 		if err != nil {
// 			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 			return res, err
// 		}

// 		datasTrans = append(datasTrans, viewmodel.TransactionVM{
// 			ID:        data.ID,
// 			MoneyIn:   data.MoneyIn.Float64,
// 			MoneyOut:  data.MoneyOut.Float64,
// 			Notes:     data.Notes,
// 			Hashtag:   tags,
// 			CreatedAt: data.CreatedAt,
// 			UpdatedAt: data.UpdatedAt.String,
// 			DeletedAt: data.DeletedAt.String,
// 		})
// 	}

// 	res = viewmodel.TransactionWithTotalAmount{
// 		TotalMoneyIn:  totalAmount.MoneyIn.Float64,
// 		TotalMoneyOut: totalAmount.MoneyOut.Float64,
// 		Transactions:  datasTrans,
// 	}

// 	return res, err
// }

// // FindTotalAmount ...
// func (uc TransactionUC) FindTotalAmount(userID string) (res viewmodel.TransactionVM, err error) {
// 	ctx := "TransactionUC.TotalAmount"
// 	transactionModel := model.NewTransactionModel(uc.DB)

// 	transaction, err := transactionModel.FindTotalAmount(userID)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	res = viewmodel.TransactionVM{
// 		ID:       transaction.ID,
// 		UserID:   transaction.UserID,
// 		MoneyIn:  transaction.MoneyIn.Float64,
// 		MoneyOut: transaction.MoneyOut.Float64,
// 	}

// 	return res, err
// }

// // Store ..
// func (uc TransactionUC) Store(UserID string, data request.TransactionRequest, tx *sql.Tx) (res viewmodel.TransactionVM, err error) {
// 	ctx := "TransactionUC.Store"
// 	transactionModel := model.NewTransactionModel(uc.DB)
// 	now := time.Now().UTC()

// 	if len(data.Tags) > 3 {
// 		logruslogger.Log(logruslogger.WarnLevel, UserID, ctx, "too_many_tags", uc.ReqID)
// 		return res, errors.New(helper.TagsTooManyInput)
// 	}

// 	if (data.MoneyIn != 0 && data.MoneyOut != 0) || (data.MoneyIn == 0 && data.MoneyOut == 0) {
// 		logruslogger.Log(logruslogger.WarnLevel, UserID, ctx, "query", uc.ReqID)
// 		return res, errors.New(helper.InvalidAmount)
// 	}

// 	datas := viewmodel.TransactionVM{
// 		UserID:   UserID,
// 		MoneyIn:  data.MoneyIn,
// 		MoneyOut: data.MoneyOut,
// 		Notes:    data.Notes,
// 	}

// 	res.ID, err = transactionModel.Store(datas, now, tx)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	tagReq := request.TagRequest{
// 		UserID:        UserID,
// 		TransactionID: res.ID,
// 		Tags:          data.Tags,
// 	}

// 	tagUC := TagUC{ContractUC: uc.ContractUC}
// 	TagsRes, err := tagUC.BulkStore(tagReq, tx)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	res = viewmodel.TransactionVM{
// 		ID:        res.ID,
// 		UserID:    UserID,
// 		MoneyIn:   data.MoneyIn,
// 		MoneyOut:  data.MoneyOut,
// 		Notes:     data.Notes,
// 		Hashtag:   TagsRes,
// 		CreatedAt: now.Format(time.RFC3339),
// 		UpdatedAt: now.Format(time.RFC3339),
// 	}

// 	return res, err
// }

// // Update ...
// func (uc TransactionUC) Update(id, userID string, data request.TransactionRequest, tx *sql.Tx) (res viewmodel.TransactionVM, err error) {
// 	ctx := "TransactionUC.Update"
// 	now := time.Now().UTC()
// 	transactionModel := model.NewTransactionModel(uc.DB)

// 	_, err = transactionModel.FindByID(id, userID)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "not_found", uc.ReqID)
// 		return res, errors.New(helper.TransactionNotFound)
// 	}

// 	datavm := viewmodel.TransactionVM{
// 		MoneyIn:  data.MoneyIn,
// 		MoneyOut: data.MoneyOut,
// 		Notes:    data.Notes,
// 	}

// 	transaction, err := transactionModel.Update(id, datavm, now, tx)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	tagReq := request.TagRequest{
// 		UserID:        userID,
// 		TransactionID: id,
// 		Tags:          data.Tags,
// 	}

// 	tagUC := TagUC{ContractUC: uc.ContractUC}
// 	TagsRes, err := tagUC.BulkUpdate(tagReq, tx)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
// 		return res, err
// 	}

// 	res = viewmodel.TransactionVM{
// 		ID:        transaction.ID,
// 		UserID:    userID,
// 		MoneyIn:   transaction.MoneyIn.Float64,
// 		MoneyOut:  transaction.MoneyOut.Float64,
// 		Notes:     transaction.Notes,
// 		Hashtag:   TagsRes,
// 		CreatedAt: transaction.CreatedAt,
// 		UpdatedAt: now.Format(time.RFC3339),
// 	}

// 	return res, err
// }

// // Delete ...
// func (uc TransactionUC) Delete(id, userID string) (res viewmodel.TransactionVM, err error) {
// 	ctx := "TransactionUC.Delete"

// 	now := time.Now().UTC()
// 	transactionModel := model.NewTransactionModel(uc.DB)

// 	_, err = transactionModel.FindByID(id, userID)
// 	if err != nil {
// 		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "not_found", uc.ReqID)
// 		return res, errors.New(helper.TransactionNotFound)
// 	}

// 	transaction, err := transactionModel.Destroy(id, now)

// 	res = viewmodel.TransactionVM{
// 		ID:        transaction.ID,
// 		UpdatedAt: now.Format(time.RFC3339),
// 		DeletedAt: now.Format(time.RFC3339),
// 	}

// 	return res, err
// }
