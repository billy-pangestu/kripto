package handler

import (
	"net/http"
	"xettle-backend/server/request"
	"xettle-backend/usecase"

	validator "gopkg.in/go-playground/validator.v9"
)

// TagHandler ...
type TagHandler struct {
	Handler
}

// GetByIDHandler ...
func (h *TagHandler) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	id := r.URL.Query().Get("id")

	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	tagUc := usecase.TagUC{ContractUC: h.ContractUC}
	res, err := tagUc.FindByID(id, userID)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

// GetByUserIDHandler ...
func (h *TagHandler) GetByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	tagUc := usecase.TagUC{ContractUC: h.ContractUC}
	res, err := tagUc.FindByUserID(userID)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

// UpdateHandler ...
func (h *TagHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	id := r.URL.Query().Get("id")

	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	req := request.DetailTag{}

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

	tagUc := usecase.TagUC{ContractUC: h.ContractUC}
	res, err := tagUc.Update(id, userID, req)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}

// DeleteHandler ...
func (h *TagHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	user := requestIDFromContextInterface(r.Context(), "user")
	userID := user["id"].(string)

	id := r.URL.Query().Get("id")

	if id == "" {
		SendBadRequest(w, "Parameter must be filled")
		return
	}

	// Get logrus request ID
	h.ContractUC.ReqID = getHeaderReqID(r)

	tagUc := usecase.TagUC{ContractUC: h.ContractUC}

	res, err := tagUc.Delete(id, userID)
	if err != nil {
		SendBadRequest(w, err.Error())
		return
	}

	SendSuccess(w, res, nil)
	return
}
