package service

import (
	"github.com/Ko8a/CATEC-Portal/pkg/repository"
	"github.com/Ko8a/CATEC-Portal/structure"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUserManage(userId int, user structure.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUserManage(userId, user)
}

func (s *UserService) GetAllUsers(userId int) ([]structure.UserInfo, error) {
	return s.repo.GetAllUsers(userId)
}

func (s *UserService) GetUserById(userId int, targetUserId int) (structure.UserFullInfo, error) {
	return s.repo.GetUserById(userId, targetUserId)
}

func (s *UserService) DeleteUserById(userId int, targetUserId int) error {
	return s.repo.DeleteUserById(userId, targetUserId)
}

func (s *UserService) UpdateUser(userId int, targetUserId int, input structure.UserFullInfo) error {
	return s.repo.UpdateUser(userId, targetUserId, input)
}

func (s *UserService) GetUsersByGroupId(groupId int) ([]structure.UserInfo, error) {
	return s.repo.GetUsersByGroupId(groupId)
}
