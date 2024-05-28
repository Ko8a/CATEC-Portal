package repository

import (
	"database/sql"
	"fmt"
	"github.com/Ko8a/CATEC-Portal/structure"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type UserPostgres struct {
	db *sqlx.DB
}

func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) CreateUserManage(userId int, user structure.User) (int, error) {
	var id int
	var groupId sql.NullInt32
	if user.GroupId != nil {
		groupId = sql.NullInt32{Int32: int32(*user.GroupId), Valid: true}
	} else {
		groupId = sql.NullInt32{Valid: false}
	}

	query := fmt.Sprintf("INSERT INTO %s (name, surname, email, age, password_hash, phone, group_id, role_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id", usersTable, rolesTable)
	row := r.db.QueryRow(query, user.Name, user.Surname, user.Email, user.Age, user.Password, user.Phone, groupId, user.RoleId)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserPostgres) GetAllUsers(userId int) ([]structure.UserInfo, error) {
	var users []structure.UserInfo
	query := fmt.Sprintf("SELECT u.id, u.email, u.name, u.surname, COALESCE(g.name, '') AS group FROM %s u LEFT JOIN %s g ON u.group_id = g.id", usersTable, groupsTable)
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserPostgres) GetUserById(userId int, targetUserId int) (structure.UserFullInfo, error) {
	// Step 1: Check if the user has the required role
	var user structure.UserFullInfo
	var roleName string
	query := fmt.Sprintf(`
		SELECT r.name FROM %s u
		INNER JOIN %s r ON u.role_id = r.id
		WHERE u.id = $1`, usersTable, rolesTable)
	err := r.db.Get(&roleName, query, userId)
	if err != nil {
		return user, err
	}

	isAllowed := true
	if roleName == "guest" && targetUserId != userId {
		isAllowed = false
	}

	if !isAllowed {
		return user, fmt.Errorf("user with id %d does not have permission to access this user", userId)
	}

	query = fmt.Sprintf("SELECT id, name, surname, age, email, phone, group_id, time_update, role_id FROM %s WHERE id = $1", usersTable)
	err = r.db.Get(&user, query, targetUserId)

	return user, err
}

func (r *UserPostgres) DeleteUserById(userId int, targetUserId int) error {
	// Step 1: Get the user's role
	var roleName string
	query := fmt.Sprintf(`SELECT r.name FROM %s u 
		INNER JOIN %s r ON u.role_id = r.id 
		WHERE u.id = $1`, usersTable, rolesTable)
	err := r.db.Get(&roleName, query, userId)
	if err != nil {
		return err
	}

	// Step 2: Check if the role is one of the allowed roles
	allowedRoles := []string{"administrator", "manager", "teacher"}
	isAllowed := false
	for _, role := range allowedRoles {
		if roleName == role || targetUserId == userId {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return fmt.Errorf("user with id %d does not have permission to update lesson", userId)
	}

	// Step 3: Delete the lesson
	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)
	_, err = r.db.Exec(query, targetUserId)

	return err
}

func (r *UserPostgres) UpdateUser(userId int, targetUserId int, input structure.UserFullInfo) error {
	// Step 1: Check if the user has the required role
	var roleName string
	query := fmt.Sprintf(`
		SELECT r.name FROM %s u
		INNER JOIN %s r ON u.role_id = r.id
		WHERE u.id = $1`, usersTable, rolesTable)
	err := r.db.Get(&roleName, query, userId)
	if err != nil {
		return err
	}

	allowedRoles := []string{"administrator", "manager", "teacher"}
	isAllowed := false
	for _, role := range allowedRoles {
		if roleName == role || targetUserId == userId {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return fmt.Errorf("user with id %d does not have permission to update lesson", userId)
	}

	if input.TimeUpdate != nil {
		return fmt.Errorf("user with id %d does not have permission to update \"TimeUpdate\" value", userId)
	}

	// Step 2: Build the update query dynamically based on provided fields
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argID))
		args = append(args, *input.Name)
		argID++
	}
	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argID))
		args = append(args, *input.Surname)
		argID++
	}
	if input.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argID))
		args = append(args, *input.Age)
		argID++
	}
	if input.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argID))
		args = append(args, *input.Email)
		argID++
	}
	if input.Phone != nil {
		setValues = append(setValues, fmt.Sprintf("phone=$%d", argID))
		args = append(args, *input.Phone)
		argID++
	}
	if input.GroupId != nil {
		setValues = append(setValues, fmt.Sprintf("group_id=$%d", argID))
		args = append(args, *input.GroupId)
		argID++
	}
	if input.TimeUpdate != nil {
		setValues = append(setValues, fmt.Sprintf("time_update=$%d", argID))
		args = append(args, *input.TimeUpdate)
		argID++
	}
	if input.RoleId != nil {
		setValues = append(setValues, fmt.Sprintf("role_id=$%s", argID))
		args = append(args, *input.RoleId)
		argID++
	}

	if len(setValues) == 0 {
		return fmt.Errorf("no valid fields provided for update")
	}

	// Step 3: Execute the update query
	query = fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", usersTable, strings.Join(setValues, ", "), argID)
	args = append(args, targetUserId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err = r.db.Exec(query, args...)

	return err
}

func (r *UserPostgres) GetUsersByGroupId(groupId int) ([]structure.UserInfo, error) {
	var users []structure.UserInfo
	query := fmt.Sprintf("SELECT u.id, u.email, u.name, u.surname, COALESCE(g.name, '') AS group FROM %s u LEFT JOIN %s g ON u.group_id = g.id WHERE u.group_id = $1", usersTable, groupsTable)
	err := r.db.Select(&users, query, groupId)
	if err != nil {
		return nil, err
	}

	return users, nil
}
