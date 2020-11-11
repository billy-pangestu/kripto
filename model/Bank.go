package model

import (
	"database/sql"
	"time"
	"xettle-backend/usecase/viewmodel"
)

// IBank ...
type IBank interface {
	FindByID(id, userID string) (BankEntity, error)
	FindByUserID(userID string) ([]BankEntity, error)
	Store(userID string, data viewmodel.BankVM, now time.Time) (viewmodel.BankVM, error)
	Update(id string, body viewmodel.BankVM, changedAt time.Time) (BankEntity, error)
	Destroy(id string, changedAt time.Time) (BankEntity, error)
}

// BankEntity ...
type BankEntity struct {
	ID        string         `db:"id"`
	UserID    string         `db:"user_id"`
	Name      string         `db:"name"`
	Balance   float64        `db:"balance"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt sql.NullString `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}

// BankModel ...
type BankModel struct {
	DB *sql.DB
}

// NewBankModel ...
func NewBankModel(db *sql.DB) IBank {
	return &BankModel{DB: db}
}

// FindByID ...
func (model BankModel) FindByID(id, userID string) (data BankEntity, err error) {
	query :=
		`SELECT "user_id", "name", "balance"
		FROM "banks"
		WHERE "id" = $1 AND "user_id" = $2 AND "deleted_at" IS NULL`
	err = model.DB.QueryRow(
		query, id, userID,
	).Scan(&data.UserID, &data.Name, &data.Balance)

	return data, err
}

// FindByUserID ...
func (model BankModel) FindByUserID(userID string) (data []BankEntity, err error) {
	query :=
		`SELECT "id", "name", "balance"
		FROM "banks"
		WHERE "user_id" = $1 AND "deleted_at" IS NULL`
	rows, err := model.DB.Query(query, userID)
	dataTemp := BankEntity{}
	for rows.Next() {
		err = rows.Scan(&dataTemp.ID, &dataTemp.Name, &dataTemp.Balance)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}

	return data, err
}

// Store ...
func (model BankModel) Store(userID string, body viewmodel.BankVM, now time.Time) (data viewmodel.BankVM, err error) {
	query :=
		`INSERT INTO "banks" ("user_id", "name", "balance", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $4)
		RETURNING "id", "name", "balance"`
	err = model.DB.QueryRow(
		query, userID, body.Name, body.Balance, now,
	).Scan(&data.ID, &data.Name, &data.Balance)

	return data, err
}

// Update ...
func (model BankModel) Update(id string, body viewmodel.BankVM, changedAt time.Time) (data BankEntity, err error) {
	query :=
		`UPDATE "banks"
		SET "name" = $1, "balance" = $2, "updated_at" = $3
		WHERE "deleted_at" IS NULL AND "id" = $4
		RETURNING "id", "name", "balance"`
	err = model.DB.QueryRow(
		query, body.Name, body.Balance, changedAt, id,
	).Scan(&data.ID, &data.Name, &data.Balance)

	return data, err
}

// Destroy ...
func (model BankModel) Destroy(id string, changedAt time.Time) (data BankEntity, err error) {
	query :=
		`UPDATE "banks"
		SET "updated_at" = $1, "deleted_at" = $1
		WHERE "deleted_at" IS NULL AND "id" = $2
		RETURNING "id", "name", "balance"`
	err = model.DB.QueryRow(
		query, changedAt, id,
	).Scan(&data.ID, &data.Name, &data.Balance)

	return data, err
}
