package usecase

import (
	"errors"
	"xettle-backend/helper"
	"xettle-backend/pkg/logruslogger"
	"xettle-backend/pkg/str"
	"xettle-backend/server/request"
	"xettle-backend/usecase/viewmodel"
)

// OTPUC ...
type OTPUC struct {
	*ContractUC
}

// RequestOTP ...
func (uc OTPUC) RequestOTP(data request.OTPRequest) (res viewmodel.UserRegisterVM, err error) {
	ctx := "OTPUC.RequestOTP"

	//check ID is available
	userUc := UserUC{ContractUC: uc.ContractUC}
	userdata, err := userUc.FindByPhoneNumber(data.PhoneNumber)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "user_not_found", uc.ReqID)
		return res, errors.New(helper.UserNotFound)
	}

	if userdata.PhoneValidAt != "" {
		logruslogger.Log(logruslogger.WarnLevel, data.PhoneNumber, ctx, "user_already_validated", uc.ReqID)
		return res, errors.New(helper.UserAlreadyValidated)
	}

	//check otp still running
	var NowOTP string
	err = uc.GetFromRedis(userdata.ID, &NowOTP)
	if NowOTP != "" {
		logruslogger.Log(logruslogger.WarnLevel, userdata.ID, ctx, "otp", uc.ReqID)
		return res, errors.New(helper.OtpStillRunning)
	}

	//Generate OTP
	ResOtp := str.RandomNumericString(4)

	//Save To Redis
	err = uc.StoreToRedisExp(userdata.ID, &ResOtp, OtpLifetime)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, userdata.ID, ctx, "invalid_save_to_redis", uc.ReqID)
		return res, err
	}

	res = viewmodel.UserRegisterVM{
		ID:  userdata.ID,
		OTP: ResOtp,
	}
	return res, err
}
