package repository

import (
	"errors"
	"fmt"
	"github.com/Ko8a/CATEC-Portal/structure"
	"github.com/jmoiron/sqlx"
)

type ManagePostgres struct {
	db *sqlx.DB
}

func NewManagePostgres(db *sqlx.DB) *ManagePostgres {
	return &ManagePostgres{db: db}
}

func (r *ManagePostgres) CreateGroup(userId int, group structure.Group) (int, error) {
	// Check if the user has administrator role
	isAdmin, err := r.isUserAdmin(userId)
	if err != nil {
		return 0, err
	}
	if !isAdmin {
		return 0, errors.New("user does not have administrator role")
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", groupsTable)
	row := r.db.QueryRow(query, group.Name)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ManagePostgres) GetAllGroups(userId int) ([]structure.Group, error) {
	var groups []structure.Group
	query := fmt.Sprintf("SELECT id, name FROM %s", groupsTable)
	err := r.db.Select(&groups, query)
	if err != nil {
		return nil, err
	}

	return groups, nil
}

func (r *ManagePostgres) CreateRole(userId int, role structure.Role) (int, error) {
	// Check if the user has administrator role
	isAdmin, err := r.isUserAdmin(userId)
	if err != nil {
		return 0, err
	}
	if !isAdmin {
		return 0, errors.New("user does not have administrator role")
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", rolesTable)
	row := r.db.QueryRow(query, role.Name)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *ManagePostgres) GetAllRoles(userId int) ([]structure.Role, error) {
	var roles []structure.Role
	query := fmt.Sprintf("SELECT id, name FROM %s", rolesTable)
	err := r.db.Select(&roles, query)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *ManagePostgres) isUserAdmin(userId int) (bool, error) {
	var roleId int
	query := fmt.Sprintf("SELECT role_id FROM %s WHERE id = $1", usersTable)
	err := r.db.Get(&roleId, query, userId)
	if err != nil {
		return false, err
	}

	adminRoleId, err := r.getAdminRoleId()
	if err != nil {
		return false, err
	}

	// Check if the roleId matches the administrator role ID (assuming adminRoleId is the ID of the administrator role)
	return roleId == adminRoleId, nil
}

func (r *ManagePostgres) getAdminRoleId() (int, error) {
	var adminRoleId int
	query := fmt.Sprintf("SELECT id FROM %s WHERE name = 'administrator'", rolesTable)

	err := r.db.Get(&adminRoleId, query)
	if err != nil {
		return 0, err
	}

	return adminRoleId, nil
}
