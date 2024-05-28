package repository

import (
	"database/sql"
	"fmt"
	"github.com/Ko8a/CATEC-Portal/structure"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strings"
)

type MarkPostgres struct {
	db *sqlx.DB
}

func NewMarkPostgres(db *sqlx.DB) *MarkPostgres {
	return &MarkPostgres{db: db}
}

// CreateMark inserts a new mark into the marks table
func (r *MarkPostgres) CreateMark(userId int, mark structure.Mark) error {
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
		if roleName == role {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return fmt.Errorf("user with id %d does not have permission to update lesson", userId)
	}

	query = fmt.Sprintf("INSERT INTO %s (user_id, lesson_id, mark, is_came) VALUES ($1, $2, $3, $4) ON CONFLICT (user_id, lesson_id) DO UPDATE SET mark = EXCLUDED.mark", studentsTable)
	_, err = r.db.Exec(query, mark.UserId, mark.LessonId, mark.Mark, mark.IsCame)
	return err
}

// UpdateMark updates an existing mark in the marks table
func (r *MarkPostgres) UpdateMark(userId int, mark structure.Mark) error {
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
		if roleName == role {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return fmt.Errorf("user with id %d does not have permission to update lesson", userId)
	}

	// Step 2: Build the update query dynamically based on provided fields
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argID := 1

	if mark.Mark != nil {
		setValues = append(setValues, fmt.Sprintf("mark=$%d", argID))
		args = append(args, *mark.Mark)
		argID++
	}
	if mark.IsCame != nil {
		setValues = append(setValues, fmt.Sprintf("is_came=$%d", argID))
		args = append(args, *mark.IsCame)
		argID++
	}

	if len(setValues) == 0 {
		return fmt.Errorf("no valid fields provided for update")
	}

	// Step 3: Execute the update query
	query = fmt.Sprintf("UPDATE %s SET %s WHERE user_id=$%d AND lesson_id=$%d", studentsTable, strings.Join(setValues, ", "), argID)
	args = append(args, mark.UserId, mark.LessonId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err = r.db.Exec(query, args...)

	return err
}

// GetMark fetches mark and performance for specific lesson of user
func (r *MarkPostgres) GetMark(userId int, lessonId int) (structure.MarkInfo, error) {
	var mark structure.MarkInfo

	// Step 1: Check if the user has the "student" role
	var roleName string
	query := fmt.Sprintf(`
		SELECT r.name FROM %s u
		INNER JOIN %s r ON u.role_id = r.id
		WHERE u.id = $1`, usersTable, rolesTable)
	err := r.db.Get(&roleName, query, userId)
	if err != nil {
		return mark, err
	}

	allowedRoles := map[string]bool{
		"administrator": true,
		"manager":       true,
		"teacher":       true,
		"student":       true,
	}

	if allowedRoles[roleName] {
		return mark, fmt.Errorf("user with id %d does not have student role", userId)
	}

	// Step 2: Check if the lesson and user have the same group_id
	var userGroupId, lessonGroupId int
	query = fmt.Sprintf("SELECT group_id FROM %s WHERE id = $1", usersTable)
	err = r.db.Get(&userGroupId, query, userId)
	if err != nil {
		return mark, err
	}

	query = fmt.Sprintf("SELECT group_id FROM %s WHERE id = $1", lessonsTable)
	err = r.db.Get(&lessonGroupId, query, lessonId)
	if err != nil {
		return mark, err
	}

	if userGroupId != lessonGroupId {
		return mark, fmt.Errorf("specified user do not have specified lesson")
	}

	// Step 3: Get the mark and attendance status for the lesson and user
	query = fmt.Sprintf(`
		SELECT mark, is_came FROM student_performance
		WHERE lesson_id = $1 AND user_id = $2
	`)
	err = r.db.Get(&mark, query, lessonId, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			// Return default values if no rows are found
			mark.Mark = nil
			mark.IsCame = false
			return mark, nil
		}
		return mark, err
	}

	return mark, err
}

// GetMarksByUser fetches all marks for a specific user
func (r *MarkPostgres) GetMarksByUser(userID int) ([]structure.MarkLessonInfo, error) {
	var marks []structure.MarkLessonInfo
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1", studentsTable)
	err := r.db.Select(&marks, query, userID)
	return marks, err
}

// GetMarksByLesson fetches all marks for a specific lesson
func (r *MarkPostgres) GetMarksByLesson(lessonID int) ([]structure.MarkUserInfo, error) {
	var marks []structure.MarkUserInfo
	query := fmt.Sprintf("SELECT * FROM %s WHERE lesson_id = $1", studentsTable)
	err := r.db.Select(&marks, query, lessonID)
	return marks, err
}

// GetAllMarks fetches all marks
func (r *MarkPostgres) GetAllMarks() ([]structure.Mark, error) {
	var marks []structure.Mark
	query := fmt.Sprintf("SELECT * FROM marks", studentsTable)
	err := r.db.Select(&marks, query)
	return marks, err
}
