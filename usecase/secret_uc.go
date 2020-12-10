package usecase

import (
	"errors"
	"time"
	"xettle-backend/helper"
	"xettle-backend/model"
	"xettle-backend/pkg/logruslogger"
	"xettle-backend/server/request"
	"xettle-backend/usecase/viewmodel"
)

// SecretUC ...
type SecretUC struct {
	*ContractUC
}

// FindByID ...
func (uc SecretUC) FindByID(id string) (res viewmodel.SecretVM, err error) {
	ctx := "SecretUC.FindByID"

	secretModel := model.NewSecretModel(uc.DB)
	data, err := secretModel.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.SecretVM{
		ID:        data.ID,
		UserID:    data.UserID.String,
		Subject:   data.Subject.String,
		Notes:     data.Notes.String,
		Password:  data.Password.String,
		CreatedAt: data.CreatedAt.String,
		UpdatedAt: data.UpdatedAt.String,
		DeletedAt: data.DeletedAt.String,
	}

	return res, err
}

// GetDecriptNote ...
func (uc SecretUC) GetDecriptNote(id, pass string) (res viewmodel.SecretVM, err error) {
	ctx := "SecretUC.GetDecriptNote"

	data, err := uc.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "query", uc.ReqID)
		return res, err
	}
	DPass, err := uc.Aes.Decrypt(data.Password)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "decript_pass", uc.ReqID)
		return res, err
	}

	if DPass != pass {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "password_not_match", uc.ReqID)
		return res, errors.New(helper.PasswordNotMatch)
	}

	DNote, err := uc.Aes.Decrypt(data.Notes)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.SecretVM{
		ID:      data.ID,
		UserID:  data.UserID,
		Subject: data.Subject,
		Notes:   DNote,
	}

	return res, err
}

// FindByUserID ...
func (uc SecretUC) FindByUserID(id string) (res []viewmodel.SecretVM, err error) {
	ctx := "SecretUC.FindByUserID"

	SecretModel := model.NewSecretModel(uc.DB)
	datas, err := SecretModel.FindByUserID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	for _, data := range datas {
		res = append(res, viewmodel.SecretVM{
			ID:      data.ID,
			Subject: data.Subject.String,
			Notes:   data.Notes.String,
		})
	}

	return res, err
}

// DeleteByID ...
func (uc SecretUC) DeleteByID(id string) (res string, err error) {
	ctx := "SecretUC.DeleteByID"

	SecretModel := model.NewSecretModel(uc.DB)

	now := time.Now().UTC()

	res, err = SecretModel.DeleteByID(id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	return res, err
}

// InsertNote ...
func (uc SecretUC) InsertNote(data request.SecretInsertNotes) (res string, err error) {
	ctx := "SecretUC.InsertNote"

	SecretModel := model.NewSecretModel(uc.DB)
	EncryptNotes, err := uc.Aes.Encrypt(data.Notes)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "encrypt_notes", uc.ReqID)
		return res, err
	}

	EncryptPassword, err := uc.Aes.Encrypt(data.Password)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "encrypt_password", uc.ReqID)
		return res, err
	}

	now := time.Now().UTC()
	datavm := viewmodel.SecretInsertVM{
		UserID:    data.UserID,
		Subject:   data.Subject,
		Notes:     EncryptNotes,
		Password:  EncryptPassword,
		CreatedAt: now.Format(time.RFC3339),
		UpdatedAt: now.Format(time.RFC3339),
	}

	res, err = SecretModel.InsertNotes(datavm)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, data.UserID, ctx, "store_failed", uc.ReqID)
		return res, errors.New(helper.StoreFailed)
	}
	return res, err
}

//UpdateNote ...
func (uc SecretUC) UpdateNote(id string, data request.SecretUpdateNote) (res string, err error) {
	ctx := "SecretUC.UpdateNote"

	SecretModel := model.NewSecretModel(uc.DB)

	result, err := uc.FindByID(id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, result.UserID, ctx, "encrypt_notes", uc.ReqID)
		return res, err
	}

	DecryptDBPass, err := uc.Aes.Decrypt(result.Password)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, result.UserID, ctx, "decrypt_db_pass", uc.ReqID)
		return res, err
	}

	if DecryptDBPass != data.Password {
		logruslogger.Log(logruslogger.WarnLevel, result.UserID, ctx, "pass_not_match", uc.ReqID)
		return res, errors.New(helper.PasswordNotMatch)
	}

	EncryptNotes, err := uc.Aes.Encrypt(data.Notes)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, result.UserID, ctx, "encrypt_notes", uc.ReqID)
		return res, err
	}

	now := time.Now().UTC()
	datavm := viewmodel.SecretUpdateVM{
		Subject:   data.Subject,
		Notes:     EncryptNotes,
		UpdatedAt: now.Format(time.RFC3339),
	}

	res, err = SecretModel.UpdateNote(id, datavm)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, result.UserID, ctx, "update_failed", uc.ReqID)
		return res, errors.New(helper.StoreFailed)
	}
	return res, err
}
