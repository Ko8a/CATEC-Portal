package service

import (
	"github.com/Ko8a/CATEC-Portal/pkg/repository"
	"github.com/Ko8a/CATEC-Portal/structure"
)

type LessonService struct {
	repo repository.Lesson
}

func NewLessonService(repo repository.Lesson) *LessonService {
	return &LessonService{repo: repo}
}

func (s *LessonService) Create(userId int, lesson structure.Lesson) (int, error) {
	return s.repo.Create(userId, lesson)
}

func (s *LessonService) GetAll(userId int) ([]structure.LessonInfo, error) {
	return s.repo.GetAll(userId)
}

func (s *LessonService) GetTodayLessons(userId int) ([]structure.LessonInfo, error) {
	return s.repo.GetTodayLessons(userId)
}

func (s *LessonService) GetWeekLessons(userId int) ([]structure.LessonInfo, error) {
	return s.repo.GetWeekLessons(userId)
}

func (s *LessonService) GetLessonById(userId int, lessonId int) (structure.LessonInfo, error) {
	return s.repo.GetLessonById(userId, lessonId)
}

func (s *LessonService) DeleteLesson(userId int, lessonId int) error {
	return s.repo.DeleteLesson(userId, lessonId)
}

func (s *LessonService) UpdateLesson(userId int, lessonId int, input structure.UpdateLessonInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.UpdateLesson(userId, lessonId, input)
}
