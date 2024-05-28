package repository

import (
	"database/sql"
	"fmt"
	"github.com/Ko8a/CATEC-Portal/structure"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type LessonPostgres struct {
	db *sqlx.DB
}

func NewLessonPostgres(db *sqlx.DB) *LessonPostgres {
	return &LessonPostgres{db: db}
}

func (r *LessonPostgres) Create(userId int, lesson structure.Lesson) (int, error) {
	var id int
	layout := "2006-01-02"
	date, err := time.Parse(layout, lesson.Date)
	if err != nil {
		return 0, err
	}

	startTime := pq.NullTime{Valid: false}
	if lesson.StartTime != nil {
		startTime = pq.NullTime{Time: time.Unix(*lesson.StartTime/1000, 0), Valid: true}
	}
	endTime := pq.NullTime{Valid: false}
	if lesson.EndTime != nil {
		endTime = pq.NullTime{Time: time.Unix(*lesson.EndTime/1000, 0), Valid: true}
	}

	var teacherId, groupId, typeId sql.NullInt32
	if lesson.TeacherId != nil {
		teacherId = sql.NullInt32{Int32: int32(*lesson.TeacherId), Valid: true}
	} else {
		teacherId = sql.NullInt32{Valid: false}
	}

	if lesson.GroupId != nil {
		groupId = sql.NullInt32{Int32: int32(*lesson.GroupId), Valid: true}
	} else {
		groupId = sql.NullInt32{Valid: false}
	}

	if lesson.TypeId != nil {
		typeId = sql.NullInt32{Int32: int32(*lesson.TypeId), Valid: true}
	} else {
		typeId = sql.NullInt32{Valid: false}
	}

	query := fmt.Sprintf("INSERT INTO %s (lesson_name, date, teacher_id, group_id, title, is_online, start_time, end_time, type_id, classroom) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id", lessonsTable)
	row := r.db.QueryRow(query, lesson.LessonName, date, teacherId, groupId, lesson.Title, lesson.IsOnline, startTime, endTime, typeId, lesson.Classroom)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *LessonPostgres) GetAll(userId int) ([]structure.LessonInfo, error) {
	var lessons []structure.LessonInfo

	query := fmt.Sprintf(`
		SELECT l.id, l.lesson_name, l.date, l.teacher_id, l.group_id, l.title, l.is_online, l.start_time, l.end_time, l.type_id, l.classroom
		FROM %s l
		INNER JOIN %s u ON u.group_id = l.group_id
		WHERE u.id = $1`, lessonsTable, usersTable)

	err := r.db.Select(&lessons, query, userId)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func (r *LessonPostgres) GetTodayLessons(userId int) ([]structure.LessonInfo, error) {
	var lessons []structure.LessonInfo

	query := fmt.Sprintf(`
		SELECT l.id, l.lesson_name, l.date, l.teacher_id, l.group_id, l.title, l.is_online, l.start_time, l.end_time, l.type_id, l.classroom
		FROM %s l
		INNER JOIN %s u ON u.group_id = l.group_id
		WHERE u.id = $1 AND l.date = CURRENT_DATE`, lessonsTable, usersTable)

	err := r.db.Select(&lessons, query, userId)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func (r *LessonPostgres) GetWeekLessons(userId int) ([]structure.LessonInfo, error) {
	var lessons []structure.LessonInfo

	query := fmt.Sprintf(`
		SELECT l.id, l.lesson_name, l.date, l.teacher_id, l.group_id, l.title, l.is_online, l.start_time, l.end_time, l.type_id, l.classroom
		FROM %s l
		INNER JOIN %s u ON u.group_id = l.group_id
		WHERE u.id = $1 
		AND l.date >= date_trunc('week', CURRENT_DATE) 
		AND l.date < date_trunc('week', CURRENT_DATE) + INTERVAL '1 week'`, lessonsTable, usersTable)

	err := r.db.Select(&lessons, query, userId)
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func (r *LessonPostgres) GetLessonById(userId int, lessonId int) (structure.LessonInfo, error) {
	var lesson structure.LessonInfo

	query := fmt.Sprintf(`SELECT l.id, l.lesson_name, l.date, l.teacher_id, l.group_id, l.title, l.is_online, l.start_time, l.end_time, l.type_id, l.classroom FROM %s l 
         INNER JOIN %s u on l.group_id = u.group_id WHERE u.id = $1 AND l.id = $2`, lessonsTable, usersTable)
	err := r.db.Get(&lesson, query, userId, lessonId)

	return lesson, err
}

func (r *LessonPostgres) DeleteLesson(userId int, lessonId int) error {
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
	allowedRoles := map[string]bool{
		"administrator": true,
		"manager":       true,
		"teacher":       true,
	}

	if !allowedRoles[roleName] {
		return fmt.Errorf("user with role %s is not allowed to delete lessons", roleName)
	}

	// Step 3: Delete the lesson
	query = fmt.Sprintf("DELETE FROM %s WHERE id = $1", lessonsTable)
	_, err = r.db.Exec(query, lessonId)

	return err
}

func (r *LessonPostgres) UpdateLesson(userId int, lessonId int, input structure.UpdateLessonInput) error {
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

	if input.LessonName != nil {
		setValues = append(setValues, fmt.Sprintf("lesson_name=$%d", argID))
		args = append(args, *input.LessonName)
		argID++
	}
	if input.Date != nil {
		setValues = append(setValues, fmt.Sprintf("date=$%d", argID))
		args = append(args, *input.Date)
		argID++
	}
	if input.TeacherId != nil {
		setValues = append(setValues, fmt.Sprintf("teacher_id=$%d", argID))
		args = append(args, *input.TeacherId)
		argID++
	}
	if input.GroupId != nil {
		setValues = append(setValues, fmt.Sprintf("group_id=$%d", argID))
		args = append(args, *input.GroupId)
		argID++
	}
	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argID))
		args = append(args, *input.Title)
		argID++
	}
	if input.IsOnline != nil {
		setValues = append(setValues, fmt.Sprintf("is_online=$%d", argID))
		args = append(args, *input.IsOnline)
		argID++
	}
	if input.StartTime != nil {
		setValues = append(setValues, fmt.Sprintf("start_time=$%d", argID))
		args = append(args, *input.StartTime)
		argID++
	}
	if input.EndTime != nil {
		setValues = append(setValues, fmt.Sprintf("end_time=$%d", argID))
		args = append(args, *input.EndTime)
		argID++
	}
	if input.TypeId != nil {
		setValues = append(setValues, fmt.Sprintf("type_id=$%d", argID))
		args = append(args, *input.TypeId)
		argID++
	}
	if input.Classroom != nil {
		setValues = append(setValues, fmt.Sprintf("classroom=$%d", argID))
		args = append(args, *input.Classroom)
		argID++
	}

	if len(setValues) == 0 {
		return fmt.Errorf("no valid fields provided for update")
	}

	// Step 3: Execute the update query
	query = fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", lessonsTable, strings.Join(setValues, ", "), argID)
	args = append(args, lessonId)

	logrus.Debugf("updateQuery: %s", query)
	logrus.Debugf("args: %s", args)

	_, err = r.db.Exec(query, args...)

	return err
}
