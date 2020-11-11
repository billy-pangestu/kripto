package handler

import (
	"net/http"

	"xettle-backend/server/request"
	"xettle-backend/usecase"

	validator "gopkg.in/go-playground/validator.v9"
)

// UserHandler ...
type UserHandler struct {
	Handler
}

// GetByTokenHandler ...
func (h *UserHandler) GetByTokenHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.FindByID(userID)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

//ChangePin ...
func (h *UserHandler) ChangePin(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	req := request.UserChangePinRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.ChangePin(userID, req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

//VerifyPin ...
func (h *UserHandler) VerifyPin(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	req := request.UserVerifyPinRequest{}
	if err := h.Handler.Bind(r, &req); err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	if err := h.Handler.Validate.Struct(req); err != nil {
		h.SendRequestValidationError(w, err.(validator.ValidationErrors))
		return
	}

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.VerifyPin(userID, req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

//Register ...
func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	req := request.UserRegisterRequest{}
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

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Register(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}

//VerifyOTP ...
func (h *UserHandler) VerifyOTP(w http.ResponseWriter, r *http.Request) {

	req := request.UserVerifyOTP{}
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

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.VerifyOTP(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}

//Login ...
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	req := request.UserLoginRequest{}
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

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.Login(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}

//SetPin ...
func (h *UserHandler) SetPin(w http.ResponseWriter, r *http.Request) {

	req := request.UserSetPin{}
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

	userUc := usecase.UserUC{ContractUC: h.ContractUC}
	res, err := userUc.SetPin(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}
