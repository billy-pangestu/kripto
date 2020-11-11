package handler

import (
	"net/http"

	"xettle-backend/server/request"
	"xettle-backend/usecase"

	validator "gopkg.in/go-playground/validator.v9"
)

// UserAuthHandler ...
type UserAuthHandler struct {
	Handler
}

//CheckPhoneNumber ...
func (h *UserAuthHandler) CheckPhoneNumber(w http.ResponseWriter, r *http.Request) {

	req := request.UserAuthCheckPhoneHandler{}
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

	userAuthUC := usecase.UserAuthUC{ContractUC: h.ContractUC}
	res, err := userAuthUC.CheckPhoneNumber(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}
