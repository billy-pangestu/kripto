package model

import (
	"database/sql"
	"time"
	"xettle-backend/usecase/viewmodel"
)

// ITag ...
type ITag interface {
	FindByID(id, userID string) (TagEntity, error)
	FindByTransactionID(TransID string) ([]TagEntity, error)
	FindByUserID(userID string) ([]TagEntity, error)
	Store(Input viewmodel.TagVM, now time.Time, tx *sql.Tx) (TagEntity, error)
	Update(id string, body viewmodel.DetailTagVM, changedAt time.Time) (TagEntity, error)
	Destroy(id string, changedAt time.Time) (TagEntity, error)
	DestroyByTransactionID(UserID, TransID string, changedAt time.Time) (TagEntity, error)
}

// TagEntity ...
type TagEntity struct {
	ID            string         `db:"id"`
	TransactionID string         `db:"transaction_id"`
	UserID        string         `db:"user_id"`
	Hashtag       string         `db:"hash_tags"`
	CreatedAt     string         `db:"created_at"`
	UpdatedAt     sql.NullString `db:"updated_at"`
	DeletedAt     sql.NullString `db:"deleted_at"`
}

// TagModel ...
type TagModel struct {
	DB *sql.DB
}

// NewTagModel ...
func NewTagModel(db *sql.DB) ITag {
	return &TagModel{DB: db}
}

// FindByID ...
func (model TagModel) FindByID(id, userID string) (data TagEntity, err error) {
	query :=
		`SELECT "transaction_id", "user_id", "hash_tags"
		FROM "tags"
		WHERE "id" = $1 AND "user_id" = $2 AND "deleted_at" IS NULL`
	err = model.DB.QueryRow(query, id, userID).Scan(&data.TransactionID, &data.UserID, &data.Hashtag)
	return data, err
}

// FindByUserID ...
func (model TagModel) FindByUserID(userID string) (data []TagEntity, err error) {
	query :=
		`SELECT DISTINCT LOWER("hash_tags")
		FROM "tags"
		WHERE "user_id" = $1 AND "deleted_at" IS NULL
		ORDER BY LOWER("hash_tags") ASC`
	rows, err := model.DB.Query(query, userID)
	dataTemp := TagEntity{}
	for rows.Next() {
		err = rows.Scan(&dataTemp.Hashtag)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}
	return data, err
}

//FindByTransactionID ...
func (model TagModel) FindByTransactionID(TransID string) (data []TagEntity, err error) {
	query := `
		SELECT "id", "hash_tags" FROM tags
		WHERE "transaction_id" = $1 AND "deleted_at" is null
		ORDER BY "hash_tags" asc
	`
	rows, err := model.DB.Query(query, TransID)
	for rows.Next() {
		dataTemp := TagEntity{}

		err = rows.Scan(&dataTemp.ID, &dataTemp.Hashtag)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}
	return data, err
}

// Store ...
func (model TagModel) Store(Input viewmodel.TagVM, now time.Time, tx *sql.Tx) (data TagEntity, err error) {
	query :=
		`INSERT INTO "tags" ("transaction_id", "user_id", "hash_tags", "created_at", "updated_at")
		VALUES($1, $2, $3, $4, $4)
		RETURNING "id", "hash_tags"`
	if tx != nil {
		err = tx.QueryRow(query,
			Input.TransactionID, Input.UserID, Input.Hashtag, now,
		).Scan(&data.ID, &data.Hashtag)
	} else {
		err = model.DB.QueryRow(query,
			Input.TransactionID, Input.UserID, Input.Hashtag, now,
		).Scan(&data.ID, &data.Hashtag)
	}

	return data, err
}

// Update ...
func (model TagModel) Update(id string, body viewmodel.DetailTagVM, changedAt time.Time) (data TagEntity, err error) {
	sql :=
		`UPDATE "tags"
		SET "hash_tags" = $1, "updated_at" = $2
		WHERE "deleted_at" IS NULL AND "id" = $3
		RETURNING "id", "hash_tags"`
	err = model.DB.QueryRow(sql, body.Hashtag, changedAt, id).Scan(&data.ID, &data.Hashtag)

	return data, err
}

// Destroy ...
func (model TagModel) Destroy(id string, changedAt time.Time) (data TagEntity, err error) {
	sql :=
		`UPDATE "tags"
		SET "updated_at" = $1, "deleted_at" = $1
		WHERE "deleted_at" IS NULL AND "id" = $2
		RETURNING "id", "user_id", "deleted_at"`
	err = model.DB.QueryRow(sql, changedAt, id).Scan(&data.ID, &data.UserID, &data.DeletedAt)

	return data, err
}

// DestroyByTransactionID ...
func (model TagModel) DestroyByTransactionID(userID, transID string, changedAt time.Time) (data TagEntity, err error) {
	sql :=
		`UPDATE "tags"
		SET "updated_at" = $1, "deleted_at" = $1
		where "user_id" = $2 AND "transaction_id" = $3
		RETURNING "id", "user_id", "updated_at", "deleted_at"`
	err = model.DB.QueryRow(sql, changedAt, userID, transID).Scan(&data.ID, &data.UserID, &data.UpdatedAt, &data.DeletedAt)

	return data, err
}
