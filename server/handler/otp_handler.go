package handler

import (
	"net/http"
	"xettle-backend/server/request"
	"xettle-backend/usecase"

	validator "gopkg.in/go-playground/validator.v9"
)

// OTPHandler ...
type OTPHandler struct {
	Handler
}

// RequestOTP ...
func (h *OTPHandler) RequestOTP(w http.ResponseWriter, r *http.Request) {
	req := request.OTPRequest{}

	if err := h.Handler.Bind(r, &req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	otpUc := usecase.OTPUC{ContractUC: h.ContractUC}
	res, err := otpUc.RequestOTP(req)

	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}
