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

// UserUC ...
type UserUC struct {
	*ContractUC
}

// BuildBody ...
func (uc UserUC) BuildBody(data *model.UserEntity, res *viewmodel.UserVM, showPass bool) {
	res.ID = data.ID
	res.FirstName = data.FirstName.String
	res.LastName = data.LastName.String
	res.UniqueID = data.UniqueID.String
	res.Amount = data.Amount
	res.CreatedAt = data.CreatedAt.String
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String

	if showPass {
		res.Password = data.Password.String
	}

	res.RoleName = data.UserRoleName.String
}

// FindByID ...
func (uc UserUC) FindByID(id string, showPass bool) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByID"

	userModel := model.NewUserModel(uc.Tx)
	data, err := userModel.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&data, &res, showPass)

	return res, err
}

// FindByIDWithoutTX ...
func (uc UserUC) FindByIDWithoutTX(id string, showPass bool) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByIDWithoutTX"

	userModel := model.NewUserModelWithoutTX(uc.DB)
	data, err := userModel.FindByIDWithoutTX(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&data, &res, showPass)

	return res, err
}

//FindByUniqueID ...
func (uc UserUC) FindByUniqueID(uniqueID string, showPass bool) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByUniqueID"
	userModel := model.NewUserModel(uc.Tx)

	data, err := userModel.FindByUniqueID(uniqueID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_not_found", uc.ReqID)
		return res, err
	}

	uc.BuildBody(&data, &res, showPass)

	return res, err
}

// Create ...
func (uc UserUC) Create(data request.UserRequest) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.Create"

	// Decrypt password dari frontend karena frontend implementasi aes untuk menencrypsi password yang akan dikirimkan
	// authUc := AuthUC{ContractUC: uc.ContractUC}
	// passwordInput, err := authUc.DecryptedOnly(data.Password)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
	// 	return res, err
	// }

	//password diubah dalam bentuk hash
	password, err := helper.HashPassword(data.Password)
	if err != nil {
		return res, err
	}

	// Error Checking Role ID
	roleUc := RoleUC{ContractUC: uc.ContractUC}
	_, err = roleUc.FindByID(data.RoleID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "find_role_by_id", uc.ReqID)
		return res, errors.New("Role Not Found")
	}

	now := time.Now().UTC()
	res = viewmodel.UserVM{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		UniqueID:  data.UniqueID,
		Password:  password,
		RoleID:    data.RoleID,
	}
	userModel := model.NewUserModel(uc.Tx)
	res.ID, err = userModel.Store(res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_store_agent", uc.ReqID)
		return res, err
	}

	res.Password = ""

	return res, err
}

// AddFund ...
func (uc UserUC) AddFund(userID string, data request.UserAddFundRequest) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.AddFund"

	_, err = uc.FindByID(userID, false)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_not_found", uc.ReqID)
		return res, errors.New(helper.UserNotFound)
	}

	now := time.Now().UTC()
	res = viewmodel.UserVM{
		Amount: data.Amount,
	}
	userModel := model.NewUserModel(uc.Tx)
	res.ID, err = userModel.AddFund(userID, res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_store_agent", uc.ReqID)
		return res, err
	}

	res.Password = ""

	return res, err
}

// SubFund ...
func (uc UserUC) SubFund(userID string, amount float64) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.SubFund"

	_, err = uc.FindByID(userID, false)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_not_found", uc.ReqID)
		return res, errors.New(helper.UserNotFound)
	}

	now := time.Now().UTC()
	res = viewmodel.UserVM{
		Amount: amount,
	}
	userModel := model.NewUserModel(uc.Tx)
	res.ID, err = userModel.SubFund(userID, res, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query_store_agent", uc.ReqID)
		return res, err
	}

	res.Password = ""

	return res, err
}

//Login ...
func (uc UserUC) Login(data request.UserLoginRequest) (res viewmodel.JwtVM, err error) {
	ctx := "UserUC.Login"

	user, err := uc.FindByUniqueID(data.UniqueID, true)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.UniqueID, ctx, "id_or_password_not_match", uc.ReqID)
		return res, errors.New("Id or Password Not Match")

	}

	// Decrypt password dari frontend jika frontend implementasi aes untuk menencrypsi password yang akan dikirimkan
	// authUc := AuthUC{ContractUC: uc.ContractUC}
	// passwordInput, err := authUc.DecryptedOnly(data.Password)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
	// 	return res, err
	// }

	passwordInput := data.Password

	// mengecheck apakah password yang diinput sama dengan password yang disimpan di dalam database
	// note: password yang disimpan dalam db dalam bentuk hash
	match := helper.CheckPasswordHash(passwordInput, user.Password)
	if !match {
		logruslogger.Log(logruslogger.WarnLevel, "invalid_password", ctx, "invalid_password", uc.ReqID)
		return res, errors.New(helper.InvalidCredentials)
	}

	// Jwe the payload & Generate jwt token
	payload := map[string]interface{}{
		"id":        user.ID,
		"role_name": user.RoleName,
		"unique_id": user.UniqueID,
	}

	// generate jwt token and store to redis
	jwtUc := JwtUC{ContractUC: uc.ContractUC}
	err = jwtUc.GenerateToken(payload, &res)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwe", uc.ReqID)
		return res, errors.New(helper.InternalServer)
	}

	return res, err
}

// Logout ...
func (uc UserUC) Logout(token, userID string) (res viewmodel.JwtVM, err error) {

	err = uc.RemoveFromRedis("userDeviceID" + userID)
	if err != nil {
		return res, err
	}

	return res, err
}
