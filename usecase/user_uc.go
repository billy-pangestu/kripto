package usecase

import (
	"errors"
	"time"
	"xettle-backend/helper"
	"xettle-backend/model"
	"xettle-backend/pkg/logruslogger"
	"xettle-backend/pkg/str"
	"xettle-backend/server/request"
	"xettle-backend/usecase/viewmodel"
)

// UserUC ...
type UserUC struct {
	*ContractUC
}

// FindByID ...
func (uc UserUC) FindByID(id string) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByID"

	userModel := model.NewUserModel(uc.DB)
	data, err := userModel.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.UserVM{
		ID:           data.ID,
		Name:         data.Name.String,
		PhoneNumber:  data.PhoneNumber.String,
		PhoneValidAt: data.PhoneValidatedAt.String,
		CreatedAt:    data.CreatedAt.String,
		UpdatedAt:    data.UpdatedAt.String,
		DeletedAt:    data.DeletedAt.String,
	}

	return res, err
}

// FindByPhoneInactiveNumber ...
func (uc UserUC) FindByPhoneInactiveNumber(PhoneNumber string) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByPhoneInactiveNumber"

	userModel := model.NewUserModel(uc.DB)
	data, err := userModel.FindByPhoneInactiveNumber(PhoneNumber)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.UserVM{
		ID:           data.ID,
		Name:         data.Name.String,
		PhoneNumber:  data.PhoneNumber.String,
		PhoneValidAt: data.PhoneValidatedAt.String,
		CreatedAt:    data.CreatedAt.String,
		UpdatedAt:    data.UpdatedAt.String,
		DeletedAt:    data.DeletedAt.String,
	}

	return res, err
}

//FindByPhoneNumber ...
func (uc UserUC) FindByPhoneNumber(phone string) (res viewmodel.UserVM, err error) {
	ctx := "UserUC.FindByPhoneNumber"
	userModel := model.NewUserModel(uc.DB)

	datas, err := userModel.FindByPhoneNumber(phone)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_not_found", uc.ReqID)
		return res, err
	}

	res = viewmodel.UserVM{
		ID:           datas.ID,
		Name:         datas.Name.String,
		PhoneNumber:  datas.PhoneNumber.String,
		Pin:          datas.Pin.String,
		PhoneValidAt: datas.PhoneValidatedAt.String,
		CreatedAt:    datas.CreatedAt.String,
		UpdatedAt:    datas.UpdatedAt.String,
		DeletedAt:    datas.DeletedAt.String,
	}
	return res, err
}

//ChangePin ...
func (uc UserUC) ChangePin(id string, data request.UserChangePinRequest) (res string, err error) {
	ctx := "UserUC.ChangePin"

	now := time.Now().UTC()
	userModel := model.NewUserModel(uc.DB)

	authUc := AuthUC{ContractUC: uc.ContractUC}
	oldPin, err := authUc.DecryptedOnly(data.OldPin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	newPininput, err := authUc.DecryptedOnly(data.NewPin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	PinDB, err := userModel.GrepPin(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}
	NewPinDB, err := authUc.DecryptedDBPin(PinDB)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	//Check OldPin and DB Pin (verifikasi)
	if oldPin != NewPinDB {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "old_Pin_invalid", uc.ReqID)
		return res, errors.New(helper.InvalidOldPin)
	}

	//Check Apakah newPininput sama dengan DB Pin
	if newPininput == NewPinDB {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "Pin_same", uc.ReqID)
		return res, errors.New(helper.PinSame)
	}

	//Encrypt newPin
	newpin, err := uc.Aes.Encrypt(newPininput)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "encrypt_Pin", uc.ReqID)
		return res, err
	}

	res, err = userModel.ChangePin(id, newpin, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	return res, err
}

//VerifyPin ...
func (uc UserUC) VerifyPin(id string, data request.UserVerifyPinRequest) (res string, err error) {
	ctx := "UserUC.VerifyPin"

	userModel := model.NewUserModel(uc.DB)

	authUc := AuthUC{ContractUC: uc.ContractUC}

	newPininput, err := authUc.DecryptedOnly(data.Pin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	PinDB, err := userModel.GrepPin(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, errors.New(helper.PinUnset)
	}
	NewPinDB, err := authUc.DecryptedDBPin(PinDB)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	//Check Apakah Pininput sama dengan DB Pin (verifiksi)
	if newPininput != NewPinDB {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "pin_not_match", uc.ReqID)
		return res, errors.New(helper.InvalidPin)
	}

	return res, err
}

// Register ...
func (uc UserUC) Register(data request.UserRegisterRequest) (res viewmodel.UserRegisterVM, err error) {
	ctx := "UserUC.Register"

	UserModel := model.NewUserModel(uc.DB)

	// Check duplicate phone number
	authUc := AuthUC{ContractUC: uc.ContractUC}
	_, err = authUc.AuthRegister(data)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "auth_invalid", uc.ReqID)
		return res, err
	}

	now := time.Now().UTC()
	datavm := viewmodel.UserStoreVM{
		Name:        data.Name,
		PhoneNumber: data.PhoneNumber,
		CreatedAt:   now.Format(time.RFC3339),
		UpdatedAt:   now.Format(time.RFC3339),
	}

	results, err := UserModel.Store(datavm)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "store_failed", uc.ReqID)
		return res, errors.New(helper.StoreFailed)
	}

	//Generate OTP
	ResOtp := str.RandomNumericString(4)

	//Save To Redis
	SaveRedis := uc.StoreToRedisExp(results.ID, &ResOtp, OtpLifetime)
	if SaveRedis != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "invalid_save_to_redis", uc.ReqID)
		return res, err
	}

	res = viewmodel.UserRegisterVM{
		ID:          results.ID,
		Name:        results.Name.String,
		PhoneNumber: results.PhoneNumber.String,
		OTP:         ResOtp,
	}
	return res, err
}

//VerifyOTP ...
func (uc UserUC) VerifyOTP(data request.UserVerifyOTP) (res string, err error) {
	ctx := "UserUC.VerifyOTP"

	//CheckOTP is true or not
	authUC := AuthUC{ContractUC: uc.ContractUC}
	checkOTP, err := authUC.VerifyOTP(data)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "verify_otp", uc.ReqID)
		return res, err
	}

	now := time.Now().UTC()

	userModel := model.NewUserModel(uc.DB)
	_, err = userModel.ValidatedOTP(checkOTP, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "query", uc.ReqID)
		return res, err
	}
	return res, err
}

//GenerateJwtToken ...
func (uc UserUC) GenerateJwtToken(id, role string) (token, exp string, err error) {
	ctx := "UserUC.GenerateJwtToken"

	payload := map[string]interface{}{
		"id":   id,
		"role": role,
	}
	jwePayload, err := uc.ContractUC.Jwe.Generate(payload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwe", uc.ReqID)
		return token, exp, errors.New(helper.JWT)
	}
	token, exp, err = uc.ContractUC.Jwt.GetToken(jwePayload)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "jwt", uc.ReqID)
		return token, exp, errors.New(helper.JWT)
	}

	return token, exp, err
}

//Login ...
func (uc UserUC) Login(PhoneNumber, Pin string) (res viewmodel.UserLoginVM, err error) {
	ctx := "UserUC.Login"

	userModel := model.NewUserModel(uc.DB)
	results, err := userModel.FindByPhoneNumber(PhoneNumber)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, PhoneNumber, ctx, "phone_or_password_not_match", uc.ReqID)
		return res, err

	}
	if results.PhoneValidatedAt.String == "" {
		logruslogger.Log(logruslogger.WarnLevel, PhoneNumber, ctx, "phone_not_verified", uc.ReqID)
		return res, errors.New(helper.PhoneNotValidated)
	}
	if results.Pin.String == "" {
		logruslogger.Log(logruslogger.WarnLevel, PhoneNumber, ctx, "pin_unset", uc.ReqID)
		return res, errors.New(helper.PinUnset)
	}

	authUc := AuthUC{ContractUC: uc.ContractUC}
	// passwordInput, err := authUc.DecryptedOnly(Pin)
	// if err != nil {
	// 	logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
	// 	return res, err
	// }

	NewPinDB, err := authUc.DecryptedDBPin(results.Pin.String)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	if Pin != NewPinDB {
		logruslogger.Log(logruslogger.WarnLevel, results.ID, ctx, "phone_or_password_not_match", uc.ReqID)
		return res, errors.New(helper.InvalidLogin)
	}

	// Jwe the payload & Generate jwt token
	res.Token, res.ExpiredDate, err = uc.GenerateJwtToken(results.ID, "user")
	return res, err
}

//SetPin ...
func (uc UserUC) SetPin(data request.UserSetPin) (res viewmodel.UserLoginVM, err error) {
	ctx := "UserUC.SetPin"

	userModel := model.NewUserModel(uc.DB)
	UserData, err := userModel.FindByPhoneNumber(data.PhoneNumber)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "user_not_found", uc.ReqID)
		return res, errors.New(helper.UserNotFound)
	}
	if UserData.Pin.String != "" {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "pin_already_set", uc.ReqID)
		return res, errors.New(helper.PinAlreadySet)
	}
	if UserData.PhoneValidatedAt.String == "" {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "phone_not_validated", uc.ReqID)
		return res, errors.New(helper.PhoneNotValidated)
	}

	now := time.Now().UTC()

	//encrypt pin
	authUc := AuthUC{ContractUC: uc.ContractUC}
	newpin, err := authUc.Encrypted(data.Pin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "encrypt", uc.ReqID)
		return res, err
	}

	_, err = userModel.SetPin(UserData.ID, newpin, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "query", uc.ReqID)
		return res, err
	}

	res.Token, res.ExpiredDate, err = uc.GenerateJwtToken(UserData.ID, "user")

	return res, err
}

//SetPinWOToken ...
func (uc UserUC) SetPinWOToken(phonenumber, newpin string) (res viewmodel.NewpinVM, err error) {
	ctx := "UserUC.SetPinWOToken"

	userModel := model.NewUserModel(uc.DB)
	UserData, err := userModel.FindByPhoneNumber(phonenumber)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "user_not_found", uc.ReqID)
		return res, errors.New(helper.UserNotFound)
	}

	now := time.Now().UTC()

	//encrypt pin
	authUc := AuthUC{ContractUC: uc.ContractUC}
	newpinDB, err := authUc.Encrypted(newpin)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "encrypt", uc.ReqID)
		return res, err
	}

	epin, err := userModel.SetPin(UserData.ID, newpinDB, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.NewpinVM{
		EPin: epin,
	}
	return res, err
}

//ShowEncryptPin ...
func (uc UserUC) ShowEncryptPin(phonenumber string) (res viewmodel.EpinVM, err error) {
	ctx := "UserUC.ShowEncryptPin"

	userModel := model.NewUserModel(uc.DB)
	UserData, err := userModel.FindByPhoneNumber(phonenumber)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "user_not_found", uc.ReqID)
		return res, errors.New(helper.UserNotFound)
	}

	res = viewmodel.EpinVM{
		EPin: UserData.Pin.String,
	}
	return res, err
}

//ShowPlainPin ...
func (uc UserUC) ShowPlainPin(phonenumber string) (res viewmodel.PpinVM, err error) {
	ctx := "UserUC.ShowPlainPin"

	userModel := model.NewUserModel(uc.DB)
	UserData, err := userModel.FindByPhoneNumber(phonenumber)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "user_not_found", uc.ReqID)
		return res, errors.New(helper.UserNotFound)
	}

	authUc := AuthUC{ContractUC: uc.ContractUC}

	pin, err := authUc.DecryptedDBPin(UserData.Pin.String)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "decrypt", uc.ReqID)
		return res, err
	}

	res = viewmodel.PpinVM{
		PPin: pin,
	}
	return res, err
}
