package usecase

import (
	"errors"
	"xettle-backend/helper"
	"xettle-backend/model"
	"xettle-backend/pkg/logruslogger"
	"xettle-backend/server/request"
)

// AuthUC ...
type AuthUC struct {
	*ContractUC
}

//AuthRegister ...
func (uc AuthUC) AuthRegister(data request.UserRegisterRequest) (res string, err error) {
	ctx := "AuthUC.AuthRegister"

	authModel := model.NewAuthModel(uc.DB)

	// Check duplicate phone number
	authphonenum, _ := authModel.FindByPhoneNum(data.PhoneNumber)
	if authphonenum.ID != "" {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "duplicate_phone_number", uc.ReqID)
		return res, errors.New(helper.DuplicatePhoneNumber)
	}

	return res, err
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

	// Encrypt password
	// newpin, err := uc.Aes.Encrypt(pin)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "encrypt_password", uc.ReqID)
	// 	return res, err
	// }

	res = pin

	return res, err
}

//DecryptedDBPin ...
func (uc AuthUC) DecryptedDBPin(pin string) (res string, err error) {
	ctx := "AuthUC.DecryptedDBPin"

	//Decrypt password input
	pin, err = uc.Aes.Decrypt(pin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "decrypt", uc.ReqID)
		return res, err
	}
	// Encrypt password
	// newpin, err := uc.Aes.Encrypt(pin)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "encrypt_password", uc.ReqID)
	// 	return res, err
	// }

	res = pin

	return res, err
}

//VerifyOTP ...
func (uc AuthUC) VerifyOTP(data request.UserVerifyOTP) (res string, err error) {
	ctx := "AuthUC.VerifyOTP"

	authModel := model.NewAuthModel(uc.DB)

	// Check phoneNumber exist
	authphonenum, _ := authModel.FindByPhoneNum(data.PhoneNumber)
	if authphonenum.ID == "" {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "duplicate_phone_number", uc.ReqID)
		return res, errors.New(helper.InvalidPhoneNumber)
	}

	//Get OTP from Redis
	var GetResRedis string
	GetRedis := uc.GetFromRedis(authphonenum.ID, &GetResRedis)
	if GetRedis != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "invalid_get_redis", uc.ReqID)
		return res, errors.New(helper.ExpOTP)
	}

	if GetResRedis != data.OTP {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "otp_failed", uc.ReqID)
		return res, errors.New(helper.OTP)
	}
	res = authphonenum.ID
	return res, err
}
