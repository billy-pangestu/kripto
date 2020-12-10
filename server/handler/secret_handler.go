package handler

import (
	"net/http"

	"xettle-backend/server/request"
	"xettle-backend/usecase"

	"github.com/go-chi/chi"
	validator "gopkg.in/go-playground/validator.v9"
)

// SecretHandler ...
type SecretHandler struct {
	Handler
}

//GetNotes ...
func (h *SecretHandler) GetNotes(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	SecretUC := usecase.SecretUC{ContractUC: h.ContractUC}
	res, err := SecretUC.FindByUserID(id)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}

//GetNote ...
func (h *SecretHandler) GetNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "idnote")
	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	SecretUC := usecase.SecretUC{ContractUC: h.ContractUC}
	res, err := SecretUC.FindByID(id)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}

//GetDecriptNote ...
func (h *SecretHandler) GetDecriptNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "idnote")
	pass := chi.URLParam(r, "password")
	if id == "" || pass == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	SecretUC := usecase.SecretUC{ContractUC: h.ContractUC}
	res, err := SecretUC.GetDecriptNote(id, pass)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}

//DeleteNote ...
func (h *SecretHandler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	SecretUC := usecase.SecretUC{ContractUC: h.ContractUC}
	res, err := SecretUC.DeleteByID(id)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}

//InsertNote ...
func (h *SecretHandler) InsertNote(w http.ResponseWriter, r *http.Request) {

	req := request.SecretInsertNotes{}
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

	SecretUC := usecase.SecretUC{ContractUC: h.ContractUC}
	res, err := SecretUC.InsertNote(req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}

//UpdateNote ...
func (h *SecretHandler) UpdateNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	req := request.SecretUpdateNote{}
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

	SecretUC := usecase.SecretUC{ContractUC: h.ContractUC}
	res, err := SecretUC.UpdateNote(id, req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}
	SendSuccess(w, res, nil)
	return
}
