package handler

import (
	"net/http"
	"xettle-backend/server/request"
	"xettle-backend/usecase"

	validator "gopkg.in/go-playground/validator.v9"
)

// BankHandler ...
type BankHandler struct {
	Handler
}

// GetByIDHandler ...
func (h *BankHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	id := r.URL.Query().Get("id")
	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	bankUc := usecase.BankUC{ContractUC: h.ContractUC}
	res, err := bankUc.FindByID(id, userID)

	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

// GetByUserIDHandler ...
func (h *BankHandler) GetByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	bankUc := usecase.BankUC{ContractUC: h.ContractUC}
	res, err := bankUc.FindByUserID(userID)

	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

// StoreHandler ...
func (h *BankHandler) StoreHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	req := request.BankRequest{}

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

	bankUc := usecase.BankUC{ContractUC: h.ContractUC}

	res, err := bankUc.Store(userID, req)

	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

// UpdateHandler ...
func (h *BankHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	id := r.URL.Query().Get("id")
	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	req := request.BankRequest{}

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

	bankUc := usecase.BankUC{ContractUC: h.ContractUC}
	res, err := bankUc.Update(id, userID, req)

	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

// DeleteHandler ...
func (h *BankHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	id := r.URL.Query().Get("id")
	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	bankUc := usecase.BankUC{ContractUC: h.ContractUC}

	res, err := bankUc.Delete(id, userID)

	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}
