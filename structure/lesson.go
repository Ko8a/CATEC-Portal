package structure

import (
	"errors"
	"time"
)

type Lesson struct {
	Id         uint64  `json:"id" db:"id"`
	LessonName string  `json:"lesson_name" db:"lesson_name" binding:"required"`
	Date       string  `json:"date" db:"date" binding:"required"`
	TeacherId  *uint64 `json:"teacher_id" db:"teacher_id"`
	GroupId    *uint64 `json:"group_id" db:"group_id" `
	Title      string  `json:"title" db:"title" binding:"required"`
	IsOnline   bool    `json:"is_online" db:"is_online"`
	StartTime  *int64  `json:"start_time" db:"start_time"`
	EndTime    *int64  `json:"end_time" db:"end_time"`
	TypeId     *uint64 `json:"type_id" db:"type_id"`
	Classroom  string  `json:"classroom" db:"classroom"`
}

type LessonInfo struct {
	Id         uint64    `json:"id" db:"id"`
	LessonName string    `json:"lesson_name" db:"lesson_name" binding:"required"`
	Date       string    `json:"date" db:"date" binding:"required"`
	TeacherId  *uint64   `json:"teacher_id" db:"teacher_id"`
	GroupId    *uint64   `json:"group_id" db:"group_id" `
	Title      string    `json:"title" db:"title" binding:"required"`
	IsOnline   bool      `json:"is_online" db:"is_online"`
	StartTime  time.Time `json:"start_time" db:"start_time"`
	EndTime    time.Time `json:"end_time" db:"end_time"`
	TypeId     *uint64   `json:"type_id" db:"type_id"`
	Classroom  string    `json:"classroom" db:"classroom"`
}

type UpdateLessonInput struct {
	LessonName *string    `json:"lesson_name" db:"lesson_name"`
	Date       *string    `json:"date" db:"date"`
	TeacherId  *uint64    `json:"teacher_id" db:"teacher_id"`
	GroupId    *uint64    `json:"group_id" db:"group_id" `
	Title      *string    `json:"title" db:"title"`
	IsOnline   *bool      `json:"is_online" db:"is_online"`
	StartTime  *time.Time `json:"start_time" db:"start_time"`
	EndTime    *time.Time `json:"end_time" db:"end_time"`
	TypeId     *uint64    `json:"type_id" db:"type_id"`
	Classroom  *string    `json:"classroom" db:"classroom"`
}

func (i UpdateLessonInput) Validate() error {
	if i.LessonName == nil && i.Date == nil && i.TeacherId == nil && i.GroupId == nil && i.Title == nil && i.IsOnline == nil && i.StartTime == nil && i.EndTime == nil && i.TypeId == nil && i.Classroom == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type Schedule struct {
	Id        uint64 `json:"-"`
	GroupId   uint64 `json:"group_id"`
	Monday    uint64 `json:"monday"`
	Tuesday   uint64 `json:"tuesday"`
	Wednesday uint64 `json:"wednesday"`
	Thursday  uint64 `json:"thursday"`
	Friday    uint64 `json:"friday"`
	Saturday  uint64 `json:"saturday"`
	Sunday    uint64 `json:"sunday"`
}

type DaySchedule struct {
	Id      uint64 `json:"-"`
	Lesson1 uint64 `json:"1st_lesson"`
	Lesson2 uint64 `json:"2nd_lesson"`
	Lesson3 uint64 `json:"3rd_lesson"`
	Lesson4 uint64 `json:"4th_lesson"`
	Lesson5 uint64 `json:"5th_lesson"`
	Lesson6 uint64 `json:"6th_lesson"`
}

type LessonType struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}
