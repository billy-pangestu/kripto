package model

import (
	"database/sql"
	"time"
	"xettle-backend/usecase/viewmodel"
)

// secretModel ...
type secretModel struct {
	DB *sql.DB
}

// ISecret ...
type ISecret interface {
	FindByID(id string) (SecretEntity, error)
	FindByPhoneInactiveNumber(number string) (UserEntity, error)
	FindByIDActiveNumber(id string) (UserEntity, error)
	ChangePin(id string, newpin string, now time.Time) (string, error)
	GrepPin(id string) (string, error)
	Store(data viewmodel.UserStoreVM) (UserEntity, error)
	ValidatedOTP(id string, now time.Time) (string, error)
	FindByPhoneNumber(PhoneNumber string) (UserEntity, error)
	SetPin(id string, pin string, now time.Time) (string, error)

	InsertNotes(data viewmodel.SecretInsertVM) (string, error)
	FindByUserID(UserID string) ([]SecretEntity, error)
	DeleteByID(ID string, t time.Time) (string, error)
	UpdateNote(ID string, data viewmodel.SecretUpdateVM) (string, error)
}

// SecretEntity ....
type SecretEntity struct {
	ID        string         `db:"id"`
	UserID    sql.NullString `db:"user_id"`
	Subject   sql.NullString `db:"subject"`
	Notes     sql.NullString `db:"notes"`
	Password  sql.NullString `db:"password"`
	CreatedAt sql.NullString `db:"createdAt"`
	UpdatedAt sql.NullString `db:"updatedAt"`
	DeletedAt sql.NullString `db:"deletedAt"`
}

// NewSecretModel ...
func NewSecretModel(db *sql.DB) ISecret {
	return &secretModel{DB: db}
}

// FindByID ...
func (model secretModel) FindByID(id string) (data SecretEntity, err error) {
	query :=
		`SELECT "id", "user_id", "subject", "notes", "password",
		"created_at", "updated_at", "deleted_at" from notes
		WHERE "id"=$1 AND "deleted_at" is null`
	err = model.DB.QueryRow(query, id).Scan(
		&data.ID, &data.UserID, &data.Subject, &data.Notes,
		&data.Password, &data.CreatedAt, &data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

//FindByPhoneInactiveNumber ...
func (model secretModel) FindByPhoneInactiveNumber(number string) (data UserEntity, err error) {
	query :=
		`SELECT "id", "name", "phone_number", "phone_validated_at",
		"created_at", "updated_at", "deleted_at" from users
		WHERE "phone_number"=$1 AND "deleted_at" is null AND "phone_validated_at" is null`
	err = model.DB.QueryRow(query, number).Scan(
		&data.ID, &data.Name, &data.PhoneNumber,
		&data.PhoneValidatedAt, &data.CreatedAt, &data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

//FindByIDActiveNumber ...
func (model secretModel) FindByIDActiveNumber(id string) (data UserEntity, err error) {
	query :=
		`SELECT "id", "name", "phone_number", "phone_validated_at",
	"created_at", "updated_at", "deleted_at" from users
	WHERE "id"=$1 AND "deleted_at" is null AND "phone_validated_at" is not null`
	err = model.DB.QueryRow(query, id).Scan(
		&data.ID, &data.Name, &data.PhoneNumber,
		&data.PhoneValidatedAt, &data.CreatedAt, &data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

//ChangePin ...
func (model secretModel) ChangePin(id string, newpin string, now time.Time) (data string, err error) {
	query := `
		UPDATE "users" set
		"pin"=$2, "updated_at"=$3
		WHERE "id"=$1 and "deleted_at" is NULL
		RETURNING "id"
	`
	err = model.DB.QueryRow(query, id, newpin, now).Scan(&data)

	return data, err
}

//GrepPin ...
func (model secretModel) GrepPin(id string) (data string, err error) {
	query := `
		SELECT "pin"
		From "users"
		WHERE "id" = $1 and "deleted_at" is NULL
	`
	err = model.DB.QueryRow(query, id).Scan(&data)

	return data, err
}

// Store ...
func (model secretModel) Store(data viewmodel.UserStoreVM) (res UserEntity, err error) {
	query := `
		INSERT INTO "users" ("phone_number", "name", "created_at", "updated_at")
		VALUES ($1, $2, $3, $3)
		RETURNING "id", "phone_number", "name"
	`
	err = model.DB.QueryRow(
		query, data.PhoneNumber, data.Name, data.CreatedAt,
	).Scan(&res.ID, &res.PhoneNumber, &res.Name)

	return res, err
}

// ValidatedOTP ...
func (model secretModel) ValidatedOTP(id string, now time.Time) (res string, err error) {
	query := `
		UPDATE "users" SET
		"phone_validated_at" = $2, "updated_at" = $2
		WHERE "id" = $1 and "deleted_at" is null
		RETURNING "id"
	`
	err = model.DB.QueryRow(query, id, now).Scan(&res)

	return res, err
}

//FindByPhoneNumber ...
func (model secretModel) FindByPhoneNumber(PhoneNumber string) (data UserEntity, err error) {
	query := `
		SELECT "id", "name", "phone_number", "pin", 
		"phone_validated_at", "created_at", "updated_at", 
		"deleted_at" from users
		WHERE "phone_number"=$1 AND "deleted_at" is null
	`
	err = model.DB.QueryRow(query, PhoneNumber).Scan(
		&data.ID, &data.Name, &data.PhoneNumber, &data.Pin,
		&data.PhoneValidatedAt, &data.CreatedAt, &data.UpdatedAt,
		&data.DeletedAt)

	return data, err
}

//SetPin ...
func (model secretModel) SetPin(id string, pin string, now time.Time) (res string, err error) {
	query := `
		UPDATE "users" SET
		"pin" = $2, "updated_at" = $3
		WHERE "id" = $1 and "deleted_at" is null
		RETURNING "pin"
	`
	err = model.DB.QueryRow(query, id, pin, now).Scan(&res)

	return res, err
}

//InsertNotes ...
func (model secretModel) InsertNotes(data viewmodel.SecretInsertVM) (res string, err error) {
	query := `
		INSERT INTO "notes" ("user_id", "subject", "notes", "password", "created_at", "updated_at")
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING "id"
	`
	err = model.DB.QueryRow(query, data.UserID, data.Subject, data.Notes, data.Password, data.CreatedAt, data.UpdatedAt).Scan(&res)

	return res, err
}

//FindByUserID ...
func (model secretModel) FindByUserID(UserID string) (data []SecretEntity, err error) {
	query := `
		SELECT "id", "subject", "notes", "password" FROM notes
		WHERE "user_id" = $1 AND "deleted_at" is null
		ORDER BY "created_at" asc
	`
	rows, err := model.DB.Query(query, UserID)
	for rows.Next() {
		dataTemp := SecretEntity{}

		err = rows.Scan(&dataTemp.ID, &dataTemp.Subject, &dataTemp.Notes, &dataTemp.Password)
		if err != nil {
			return data, err
		}

		data = append(data, dataTemp)
	}
	return data, err
}

//DeleteByID ...
func (model secretModel) DeleteByID(ID string, t time.Time) (res string, err error) {
	query := `
		UPDATE "notes" SET
		"updated_at" = $2, "deleted_at" = $2
		WHERE "id" = $1 AND "deleted_at" IS NULL
		RETURNING "id"
	`
	err = model.DB.QueryRow(query, ID, t).Scan(&res)

	return res, err
}

// UpdateNote ...
func (model secretModel) UpdateNote(ID string, data viewmodel.SecretUpdateVM) (res string, err error) {
	query := `
		UPDATE "notes" SET
		"subject" = $2, "notes" = $3, "updated_at" = $4
		WHERE "id" = $1 AND "deleted_at" IS NULL
		RETURNING "id"
	`

	err = model.DB.QueryRow(query, ID, data.Subject, data.Notes, data.UpdatedAt).Scan(&res)

	return res, err
}
