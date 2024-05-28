package service

import (
	"github.com/Ko8a/CATEC-Portal/pkg/repository"
	"github.com/Ko8a/CATEC-Portal/structure"
)

type ManageService struct {
	repo repository.Manage
}

func NewManageService(repo repository.Manage) *ManageService {
	return &ManageService{repo: repo}
}

func (s *ManageService) CreateGroup(userId int, group structure.Group) (int, error) {
	return s.repo.CreateGroup(userId, group)
}

func (s *ManageService) GetAllGroups(userId int) ([]structure.Group, error) {
	return s.repo.GetAllGroups(userId)
}

func (s *ManageService) CreateRole(userId int, role structure.Role) (int, error) {
	return s.repo.CreateRole(userId, role)
}

func (s *ManageService) GetAllRoles(userId int) ([]structure.Role, error) {
	return s.repo.GetAllRoles(userId)
}
