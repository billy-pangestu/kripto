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

// BankUC ...
type BankUC struct {
	*ContractUC
}

// FindByID ...
func (uc BankUC) FindByID(id, userID string) (res viewmodel.BankVM, err error) {
	ctx := "BankUC.FindByID"

	bankModel := model.NewBankModel(uc.DB)

	bank, err := bankModel.FindByID(id, userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, errors.New(helper.BankNotFound)
	}

	res = viewmodel.BankVM{
		ID:      bank.ID,
		Name:    bank.Name,
		Balance: bank.Balance,
	}

	return res, err
}

// FindByUserID ...
func (uc BankUC) FindByUserID(userID string) (res []viewmodel.BankVM, err error) {
	ctx := "BankUC.FindByUserID"

	bankModel := model.NewBankModel(uc.DB)
	now := time.Now().UTC()

	bank, err := bankModel.FindByUserID(userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	for _, data := range bank {
		res = append(res, viewmodel.BankVM{
			ID:        data.ID,
			UserID:    userID,
			Name:      data.Name,
			Balance:   data.Balance,
			CreatedAt: now.Format(time.RFC3339),
		})
	}

	return res, err
}

// Store ...
func (uc BankUC) Store(userID string, data request.BankRequest) (res viewmodel.BankVM, err error) {
	ctx := "BankUC.Store"

	bankModel := model.NewBankModel(uc.DB)
	now := time.Now().UTC()

	datas := viewmodel.BankVM{
		Name:    data.Name,
		Balance: data.Balance,
	}

	res, err = bankModel.Store(userID, datas, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, errors.New(helper.StoreFailed)
	}

	return res, err
}

// Update ...
func (uc BankUC) Update(id, userID string, data request.BankRequest) (res viewmodel.BankVM, err error) {
	ctx := "BankUC.Update"

	now := time.Now().UTC()
	bankModel := model.NewBankModel(uc.DB)

	_, err = bankModel.FindByID(id, userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, errors.New(helper.BankNotFound)
	}

	datas := viewmodel.BankVM{
		Name:    data.Name,
		Balance: data.Balance,
	}

	bank, err := bankModel.Update(id, datas, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.BankVM{
		ID:        bank.ID,
		Name:      bank.Name,
		Balance:   bank.Balance,
		UpdatedAt: now.Format(time.RFC3339),
	}

	return res, err
}

// Delete ...
func (uc BankUC) Delete(id, userID string) (res viewmodel.BankVM, err error) {
	ctx := "BankUC.Update"

	now := time.Now().UTC()
	bankModel := model.NewBankModel(uc.DB)

	_, err = bankModel.FindByID(id, userID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, errors.New(helper.BankNotFound)
	}

	bank, err := bankModel.Destroy(id, now)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), ctx, "query", uc.ReqID)
		return res, err
	}

	res = viewmodel.BankVM{
		ID:        bank.ID,
		UpdatedAt: now.Format(time.RFC3339),
		DeletedAt: now.Format(time.RFC3339),
	}

	return res, err
}
