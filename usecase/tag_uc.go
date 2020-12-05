package usecase

import (
	"database/sql"
	"errors"
	"time"
	"xettle-backend/helper"
	"xettle-backend/model"
	"xettle-backend/pkg/logruslogger"
	"xettle-backend/server/request"
	"xettle-backend/usecase/viewmodel"
)

// TagUC ...
type TagUC struct {
	*ContractUC
}

// FindByID ...
func (uc TagUC) FindByID(id, userID string) (res viewmodel.TagVM, err error) {
	ctx := "TagUC.FindByID"

	tagModel := model.NewTagModel(uc.DB)
	tag, err := tagModel.FindByID(id, userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.TagVM{
		ID:            tag.ID,
		TransactionID: tag.TransactionID,
		UserID:        tag.UserID,
		Hashtag:       tag.Hashtag,
	}

	return res, err
}

// BulkStore ...
func (uc TagUC) BulkStore(UserID string, data []request.DetailTag, TransID string, tx *sql.Tx) (res []viewmodel.DetailTagVM, err error) {
	ctx := "TagUC.BulkStore"
	for _, t := range data {
		storing, err := uc.Store(TransID, UserID, t.Name, tx)
		if err != nil {
			logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
			return res, err
		}
		res = append(res, storing)
	}
	return res, err
}

// Store ..
func (uc TagUC) Store(TransID, UserID, tagName string, tx *sql.Tx) (res viewmodel.DetailTagVM, err error) {
	ctx := "TagUC.Store"
	TagModel := model.NewTagModel(uc.DB)
	now := time.Now().UTC()

	inputVM := viewmodel.TagVM{
		TransactionID: TransID,
		UserID:        UserID,
		Hashtag:       tagName,
	}

	output, err := TagModel.Store(inputVM, now, tx)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, errors.New(helper.StoreFailed)
	}

	res = viewmodel.DetailTagVM{
		ID:      output.ID,
		Hashtag: output.Hashtag,
	}

	return res, err
}

// FindByUserID ...
func (uc TagUC) FindByUserID(userID string) (res []viewmodel.DetailTagVM, err error) {
	ctx := "TagUC.FindByUserID"
	tagModel := model.NewTagModel(uc.DB)

	tag, err := tagModel.FindByUserID(userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	for _, data := range tag {
		res = append(res, viewmodel.DetailTagVM{
			Hashtag: data.Hashtag,
		})
	}

	return res, err
}

// BulkUpdate ...
func (uc TagUC) BulkUpdate(UserID, TransID string, data []request.DetailTag, tx *sql.Tx) (res []viewmodel.DetailTagVM, err error) {
	ctx := "TagUC.BulkUpdate"
	now := time.Now().UTC()

	// Delete existing tags
	tagModel := model.NewTagModel(uc.DB)
	_, err = tagModel.DestroyByTransactionID(UserID, TransID, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
	}

	// Create new tags
	res, err = uc.BulkStore(UserID, data, TransID, tx)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}
	return res, err
}

// Update ...
func (uc TagUC) Update(id, userID string, data request.DetailTag) (res viewmodel.DetailTagVM, err error) {
	ctx := "TagUC.Update"
	now := time.Now().UTC()
	tagModel := model.NewTagModel(uc.DB)

	_, err = tagModel.FindByID(id, userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "not_found", uc.ReqID)
		return res, errors.New(helper.TagNotFound)
	}

	datavm := viewmodel.DetailTagVM{
		Hashtag: data.Name,
	}

	tag, err := tagModel.Update(id, datavm, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.DetailTagVM{
		ID:        tag.ID,
		Hashtag:   tag.Hashtag,
		UpdatedAt: now.Format(time.RFC3339),
	}

	return res, err
}

// Delete ...
func (uc TagUC) Delete(id, userID string) (res viewmodel.DetailTagVM, err error) {
	ctx := "TagUC.Delete"

	now := time.Now().UTC()
	tagModel := model.NewTagModel(uc.DB)

	_, err = tagModel.FindByID(id, userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, id, ctx, "not_found", uc.ReqID)
		return res, errors.New(helper.TagNotFound)
	}

	tag, err := tagModel.Destroy(id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.DetailTagVM{
		ID:        tag.ID,
		UpdatedAt: now.Format(time.RFC3339),
		DeletedAt: now.Format(time.RFC3339),
	}

	return res, err
}

// FindByTransactionID ...
func (uc TagUC) FindByTransactionID(transID string) (res []viewmodel.DetailTagVM, err error) {
	ctx := "TagUC.FindByTransactionID"
	TagModel := model.NewTagModel(uc.DB)

	tags, err := TagModel.FindByTransactionID(transID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	for _, data := range tags {
		res = append(res, viewmodel.DetailTagVM{
			ID:      data.ID,
			Hashtag: data.Hashtag,
		})
	}

	return res, err
}
