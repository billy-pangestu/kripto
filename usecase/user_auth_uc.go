package usecase

import (
	"errors"
	"xettle-backend/helper"
	"xettle-backend/pkg/logruslogger"
	"xettle-backend/pkg/str"
	"xettle-backend/server/request"
	"xettle-backend/usecase/viewmodel"
)

// UserAuthUC ...
type UserAuthUC struct {
	*ContractUC
}

const (
	actionlogin    = "login"
	actionregister = "register"
)

//CheckPhoneNumber ...
func (uc UserAuthUC) CheckPhoneNumber(data request.UserAuthCheckPhoneHandler) (res viewmodel.UserAuthCheckPhoneVM, err error) {
	ctx := "UserUC.CheckPhoneNumber"

	checkisnumber := str.CheckNumeric(data.PhoneNumber, true)
	if checkisnumber == false {
		return res, errors.New(helper.InvalidPhoneNumber)
	}

	userUC := UserUC{ContractUC: uc.ContractUC}
	UserData, err := userUC.FindByPhoneNumber(data.PhoneNumber)
	if err != nil {
		res.Action = actionregister
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "user_register", uc.ReqID)
		return res, nil
	}

	if UserData.PhoneValidAt == "" {
		res.Action = VerifyOtpAction
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "verify_otp", uc.ReqID)
		return res, nil
	}

	if UserData.Pin == "" {
		res.Action = SetPinAction
		logruslogger.Log(logruslogger.WarnLevel, UserData.ID, ctx, "pin_unset", uc.ReqID)
		return res, nil
	}

	res.Action = actionlogin

	return res, err
}
