package usecase

import (
	"payment-gateway-backend/pkg/logruslogger"
)

// AuthUC ...
type AuthUC struct {
	*ContractUC
}

//Encrypted ...
func (uc AuthUC) Encrypted(pin string) (res string, err error) {
	ctx := "AuthUC.Encrypted"

	//Decrypt password input
	pin, err = uc.AesFront.Decrypt(pin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "decrypt", uc.ReqID)
		return res, err
	}

	// Encrypt password
	newpin, err := uc.Aes.Encrypt(pin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "encrypt_password", uc.ReqID)
		return res, err
	}

	res = newpin

	return res, err
}

//DecryptedOnly ...
func (uc AuthUC) DecryptedOnly(pin string) (res string, err error) {
	ctx := "AuthUC.DecryptedOnly"

	//Decrypt password input
	pin, err = uc.AesFront.Decrypt(pin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "decrypt", uc.ReqID)
		return res, err
	}

	res = pin

	return res, err
}
