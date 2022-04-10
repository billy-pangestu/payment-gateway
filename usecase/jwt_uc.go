package usecase

import (
	"errors"
	"payment-gateway-backend/helper"
	"payment-gateway-backend/pkg/logruslogger"
	"payment-gateway-backend/usecase/viewmodel"

	"github.com/rs/xid"
)

// JwtUC ...
type JwtUC struct {
	*ContractUC
}

// GenerateToken ...
func (uc JwtUC) GenerateToken(payload map[string]interface{}, res *viewmodel.JwtVM) (err error) {
	ctx := "JwtUC.GenerateToken"

	deviceID := xid.New().String()
	payload["device_id"] = deviceID
	err = uc.StoreToRedisExp("userDeviceID"+payload["id"].(string), deviceID, uc.EnvConfig["TOKEN_EXP_SECRET"]+"h")
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "device_id", uc.ReqID)
		return errors.New(helper.InternalServer)
	}

	jwePayload, err := uc.ContractUC.Jwe.Generate(payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwe", uc.ReqID)
		return errors.New(helper.JWT)
	}
	res.Token, res.ExpiredDate, err = uc.ContractUC.Jwt.GetToken(jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwt", uc.ReqID)
		return errors.New(helper.JWT)
	}
	res.RefreshToken, res.RefreshExpiredDate, err = uc.ContractUC.Jwt.GetRefreshToken(jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "refresh_jwt", uc.ReqID)
		return errors.New(helper.JWT)
	}

	return err
}
