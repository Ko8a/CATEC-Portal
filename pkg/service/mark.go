package service

import (
	"github.com/Ko8a/CATEC-Portal/pkg/repository"
	"github.com/Ko8a/CATEC-Portal/structure"
)

type MarkService struct {
	repo repository.Mark
}

func NewMarkService(repo repository.Mark) *MarkService {
	return &MarkService{repo: repo}
}

func (s *MarkService) CreateMark(userId int, mark structure.Mark) error {
	return s.repo.CreateMark(userId, mark)
}

func (s *MarkService) UpdateMark(userId int, mark structure.Mark) error {
	return s.repo.UpdateMark(userId, mark)
}

func (s *MarkService) GetMark(userId int, lessonId int) (structure.MarkInfo, error) {
	return s.repo.GetMark(userId, lessonId)
}

func (s *MarkService) GetMarksByUser(userID int) ([]structure.MarkLessonInfo, error) {
	return s.repo.GetMarksByUser(userID)
}

func (s *MarkService) GetMarksByLesson(lessonID int) ([]structure.MarkUserInfo, error) {
	return s.repo.GetMarksByLesson(lessonID)
}

func (s *MarkService) GetAllMarks() ([]structure.Mark, error) {
	return s.repo.GetAllMarks()
}
