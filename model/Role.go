package model

import (
	"database/sql"
)

var (
	// RoleCodeSuperadmin ...
	RoleCodeSuperadmin = "superadmin"
	// RoleCodeAdmin ...
	RoleCodeAdmin = "admin"
)

// roleModel ...
type roleModel struct {
	DB *sql.DB
}

// IRole ...
type IRole interface {
	FindByID(id string) (RoleEntity, error)
	FindByCode(code string) (RoleEntity, error)
}

// RoleEntity ....
type RoleEntity struct {
	ID        string         `db:"id"`
	Code      sql.NullString `db:"code"`
	Name      sql.NullString `db:"name"`
	Status    sql.NullBool   `db:"status"`
	CreatedAt string         `db:"created_at"`
	UpdatedAt string         `db:"updated_at"`
	DeletedAt sql.NullString `db:"deleted_at"`
}

// NewRoleModel ...
func NewRoleModel(db *sql.DB) IRole {
	return &roleModel{DB: db}
}

// FindByID ...
func (model roleModel) FindByID(id string) (data RoleEntity, err error) {
	query :=
		`SELECT "id", "name", "created_at", "updated_at", "deleted_at"
		FROM "user_roles" WHERE "deleted_at" IS NULL AND "id" = $1
		ORDER BY "created_at" DESC LIMIT 1`
	err = model.DB.QueryRow(query, id).Scan(
		&data.ID, &data.Name, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
	)

	return data, err
}

// FindByCode ...
func (model roleModel) FindByCode(code string) (data RoleEntity, err error) {
	query :=
		`SELECT "id", "code", "name", "status", "created_at", "updated_at", "deleted_at"
		FROM "roles" WHERE "deleted_at" IS NULL AND "code" = $1
		ORDER BY "created_at" DESC LIMIT 1`
	err = model.DB.QueryRow(query, code).Scan(
		&data.ID, &data.Code, &data.Name, &data.Status, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt,
	)

	return data, err
}
