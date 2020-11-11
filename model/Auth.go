package model

import (
	"database/sql"
	"strings"
	"time"
	"xettle-backend/usecase/viewmodel"
)

// authModel ...
type authModel struct {
	DB *sql.DB
}

// IAuth ...
type IAuth interface {
	FindAll(offset, limit int) ([]AdminEntity, int, error)
	FindByID(id string) (AdminEntity, error)
	FindByCode(code string) (AdminEntity, error)
	FindByEmail(email string) (AuthEntity, error)
	FindByPhoneNum(phoneNum string) (data AuthEntity, err error)

	Update(id string, body viewmodel.AdminVM, changedAt time.Time) (string, error)
	Destroy(id string, changedAt time.Time) (string, error)
}

// AuthEntity ....
type AuthEntity struct {
	ID                string         `db:"id"`
	CodePhoneID       sql.NullString `db:"code_phone_id"`
	PhoneNum          sql.NullString `db:"phone_number"`
	Email             sql.NullString `db:"email"`
	Password          sql.NullString `db:"password"`
	FullName          sql.NullString `db:"full_name"`
	BirthDate         sql.NullString `db:"birth_date"`
	GenderID          sql.NullString `db:"gender_id"`
	Member            sql.NullString `db:"member"`
	Coins             sql.NullInt64  `db:"coins"`
	KtpAuth           sql.NullString `db:"ktp_auth"`
	KtpValidatedAt    sql.NullString `db:"ktp_validated_at"`
	ProfileArtistAuth sql.NullString `db:"profile_artist_auth"`
	ArtistValidatedAt sql.NullString `db:"artist_validated_at"`
	CreatedAt         string         `db:"created_at"`
	UpdatedAt         string         `db:"updated_at"`
	DeletedAt         sql.NullString `db:"deleted_at"`
	BannedAt          sql.NullString `db:"banned_at"`
}

// NewAuthModel ...
func NewAuthModel(db *sql.DB) IAuth {
	return &authModel{DB: db}
}

// FindAll ...
func (model authModel) FindAll(offset, limit int) (data []AdminEntity, count int, err error) {
	query := `SELECT a."id", a."code", a."name", a."email", a."password", a."role_id", a."status",
	a."created_at", a."updated_at", a."deleted_at", r."code", r."name"
	FROM "admins" a
	LEFT JOIN "roles" r ON r."id" = a."role_id"
	WHERE a."deleted_at" IS NULL ORDER BY a."created_at" DESC OFFSET $1 LIMIT $2`
	rows, err := model.DB.Query(query, offset, limit)
	if err != nil {
		return data, count, err
	}

	defer rows.Close()
	for rows.Next() {
		d := AdminEntity{}
		err = rows.Scan(
			&d.ID, &d.Code, &d.Name, &d.Email, &d.Password, &d.RoleID, &d.Status, &d.CreatedAt,
			&d.UpdatedAt, &d.DeletedAt, &d.Role.Code, &d.Role.Name,
		)
		if err != nil {
			return data, count, err
		}
		data = append(data, d)
	}

	err = rows.Err()
	if err != nil {
		return data, count, err
	}

	query = `SELECT COUNT("id") FROM "admins" WHERE "deleted_at" IS NULL`
	err = model.DB.QueryRow(query).Scan(&count)

	return data, count, err
}

// FindByID ...
func (model authModel) FindByID(id string) (data AdminEntity, err error) {
	query :=
		`SELECT a."id", a."code", a."name", a."email", a."password", a."role_id", a."status",
		a."created_at", a."updated_at", a."deleted_at", r."code", r."name"
		FROM "admins" a
		LEFT JOIN "roles" r ON r."id" = a."role_id"
		WHERE a."deleted_at" IS NULL AND a."id" = $1
		ORDER BY a."created_at" DESC LIMIT 1`
	err = model.DB.QueryRow(query, id).Scan(
		&data.ID, &data.Code, &data.Name, &data.Email, &data.Password, &data.RoleID, &data.Status,
		&data.CreatedAt, &data.UpdatedAt, &data.DeletedAt, &data.Role.Code, &data.Role.Name,
	)

	return data, err
}

// FindByCode ...
func (model authModel) FindByCode(code string) (data AdminEntity, err error) {
	query :=
		`SELECT a."id", a."code", a."name", a."email", a."password", a."role_id", a."status",
		a."created_at", a."updated_at", a."deleted_at", r."code", r."name"
		FROM "admins" a
		LEFT JOIN "roles" r ON r."id" = a."role_id"
		WHERE a."deleted_at" IS NULL AND a."code" = $1
		ORDER BY a."created_at" DESC LIMIT 1`
	err = model.DB.QueryRow(query, code).Scan(
		&data.ID, &data.Code, &data.Name, &data.Email, &data.Password, &data.RoleID, &data.Status,
		&data.CreatedAt, &data.UpdatedAt, &data.DeletedAt, &data.Role.Code, &data.Role.Name,
	)

	return data, err
}

// FindByEmail ...
func (model authModel) FindByEmail(email string) (data AuthEntity, err error) {
	query :=
		`SELECT id from "users" where "email"=$1 
		and banned_at is NULL and deleted_at is NULL`
	err = model.DB.QueryRow(query, strings.ToLower(email)).Scan(
		&data.ID,
	)

	return data, err
}

// FindByPhoneNum ...
func (model authModel) FindByPhoneNum(phoneNum string) (data AuthEntity, err error) {
	query :=
		`SELECT id from "users" where "phone_number"=$1 
		and deleted_at is NULL`
	err = model.DB.QueryRow(query, phoneNum).Scan(
		&data.ID,
	)

	return data, err
}

// Update ...
func (model authModel) Update(id string, body viewmodel.AdminVM, changedAt time.Time) (res string, err error) {
	sql :=
		`UPDATE "admins"
		SET "code" = $1, "name" = $2, "email" = $3, "password" = $4, "role_id" = $5, "status" = $6,
		"updated_at" = $7 WHERE "deleted_at" IS NULL AND "id" = $8 RETURNING "id"`
	err = model.DB.QueryRow(sql,
		body.Code, body.Name, body.Email, body.Password, body.RoleID, body.Status, changedAt, id,
	).Scan(&res)

	return res, err
}

// Destroy ...
func (model authModel) Destroy(id string, changedAt time.Time) (res string, err error) {
	sql :=
		`UPDATE "admins" SET "updated_at" = $1, "deleted_at" = $1
		WHERE "deleted_at" IS NULL AND "id" = $2 RETURNING "id"`
	err = model.DB.QueryRow(sql, changedAt, id).Scan(&res)

	return res, err
}
