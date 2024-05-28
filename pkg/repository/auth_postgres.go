package repository

import (
	"database/sql"
	"fmt"
	"github.com/Ko8a/CATEC-Portal/structure"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user structure.User) (int, error) {
	var id int

	var groupId sql.NullInt32
	if user.GroupId != nil {
		groupId = sql.NullInt32{Int32: int32(*user.GroupId), Valid: true}
	} else {
		groupId = sql.NullInt32{Valid: false}
	}

	query := fmt.Sprintf("INSERT INTO %s (name, surname, email, age, password_hash, phone, group_id, role_id) VALUES ($1, $2, $3, $4, $5, $6, $7, (SELECT id FROM %s WHERE name = 'guest')) RETURNING id", usersTable, rolesTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Age, user.Password, user.Phone, groupId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (structure.User, error) {
	var user structure.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)

	return user, err
}
