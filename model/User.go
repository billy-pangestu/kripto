package model

import (
	"database/sql"
	"time"
	"xettle-backend/usecase/viewmodel"
)

// userModel ...
type userModel struct {
	DB *sql.DB
}

// IUser ...
type IUser interface {
	FindByID(id string) (UserEntity, error)
	FindByPhoneInactiveNumber(number string) (UserEntity, error)
	FindByIDActiveNumber(id string) (UserEntity, error)
	ChangePin(id string, newpin string, now time.Time) (string, error)
	GrepPin(id string) (string, error)
	Store(data viewmodel.UserStoreVM) (UserEntity, error)
	ValidatedOTP(id string, now time.Time) (string, error)
	FindByPhoneNumber(PhoneNumber string) (UserEntity, error)
	SetPin(id string, pin string, now time.Time) (string, error)
}

// UserEntity ....
type UserEntity struct {
	ID               string         `db:"id"`
	Name             sql.NullString `db:"name"`
	PhoneNumber      sql.NullString `db:"phone_number"`
	PhoneValidatedAt sql.NullString `db:"phone_validated_at"`
	Pin              sql.NullString `db:"pin"`
	CreatedAt        sql.NullString `db:"createdAt"`
	UpdatedAt        sql.NullString `db:"updatedAt"`
	DeletedAt        sql.NullString `db:"deletedAt"`
}

// NewUserModel ...
func NewUserModel(db *sql.DB) IUser {
	return &userModel{DB: db}
}

// FindByID ...
func (model userModel) FindByID(id string) (data UserEntity, err error) {
	query :=
		`SELECT "id", "name", "phone_number", "pin", "phone_validated_at",
		"created_at", "updated_at", "deleted_at" from users
		WHERE "id"=$1 AND "deleted_at" is null AND "phone_validated_at" is not null`
	err = model.DB.QueryRow(query, id).Scan(
		&data.ID, &data.Name, &data.PhoneNumber, &data.Pin,
		&data.PhoneValidatedAt, &data.CreatedAt, &data.UpdatedAt,
		&data.DeletedAt,
	)

	return data, err
}

//FindByPhoneInactiveNumber ...
func (model userModel) FindByPhoneInactiveNumber(number string) (data UserEntity, err error) {
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
func (model userModel) FindByIDActiveNumber(id string) (data UserEntity, err error) {
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
func (model userModel) ChangePin(id string, newpin string, now time.Time) (data string, err error) {
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
func (model userModel) GrepPin(id string) (data string, err error) {
	query := `
		SELECT "pin"
		From "users"
		WHERE "id" = $1 and "deleted_at" is NULL
	`
	err = model.DB.QueryRow(query, id).Scan(&data)

	return data, err
}

// Store ...
func (model userModel) Store(data viewmodel.UserStoreVM) (res UserEntity, err error) {
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
func (model userModel) ValidatedOTP(id string, now time.Time) (res string, err error) {
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
func (model userModel) FindByPhoneNumber(PhoneNumber string) (data UserEntity, err error) {
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
func (model userModel) SetPin(id string, pin string, now time.Time) (res string, err error) {
	query := `
		UPDATE "users" SET
		"pin" = $2, "updated_at" = $3
		WHERE "id" = $1 and "deleted_at" is null
		RETURNING "id"
	`
	err = model.DB.QueryRow(query, id, pin, now).Scan(&res)

	return res, err
}
