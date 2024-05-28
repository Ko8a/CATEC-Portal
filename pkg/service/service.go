package service

import (
	"github.com/Ko8a/CATEC-Portal/pkg/repository"
	"github.com/Ko8a/CATEC-Portal/structure"
)

type Authorization interface {
	CreateUser(user structure.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type User interface {
	CreateUserManage(userId int, user structure.User) (int, error)
	GetAllUsers(userId int) ([]structure.UserInfo, error)
	GetUserById(userId int, targetUserId int) (structure.UserFullInfo, error)
	DeleteUserById(userId int, targetUserId int) error
	UpdateUser(userId int, targetUserId int, input structure.UserFullInfo) error
	GetUsersByGroupId(groupId int) ([]structure.UserInfo, error)
}

type Manage interface {
	CreateGroup(userId int, group structure.Group) (int, error)
	GetAllGroups(userId int) ([]structure.Group, error)
	CreateRole(userId int, role structure.Role) (int, error)
	GetAllRoles(userId int) ([]structure.Role, error)
}

type Lesson interface {
	Create(userId int, lesson structure.Lesson) (int, error)
	GetAll(userId int) ([]structure.LessonInfo, error)
	GetTodayLessons(userId int) ([]structure.LessonInfo, error)
	GetWeekLessons(userId int) ([]structure.LessonInfo, error)
	GetLessonById(userId int, lessonId int) (structure.LessonInfo, error)
	DeleteLesson(userId int, lessonId int) error
	UpdateLesson(userId int, lessonId int, input structure.UpdateLessonInput) error
}

type Mark interface {
	CreateMark(userId int, mark structure.Mark) error
	UpdateMark(userId int, mark structure.Mark) error
	GetMark(userId int, lessonId int) (structure.MarkInfo, error)
	GetMarksByUser(userID int) ([]structure.MarkLessonInfo, error)
	GetMarksByLesson(lessonID int) ([]structure.MarkUserInfo, error)
	GetAllMarks() ([]structure.Mark, error)
}

type Service struct {
	Authorization
	Manage
	Lesson
	User
	Mark
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Lesson:        NewLessonService(repos.Lesson),
		Manage:        NewManageService(repos.Manage),
		User:          NewUserService(repos.User),
		Mark:          NewMarkService(repos.Mark),
	}
}
