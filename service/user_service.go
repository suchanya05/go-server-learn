package service

import (
	"go-server/models"
	"go-server/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

// CreateUser - สร้างผู้ใช้
func (s *UserService) CreateUser(user models.User) (models.User, error) {
	return s.UserRepo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.GetUsers(nil, nil)
}

func (s *UserService) GetUsers(id *int) ([]models.User, error) {
	return s.UserRepo.GetUsers(id, nil)
}

func (s *UserService) GetUserName(username *string) ([]models.User, error) {
	return s.UserRepo.GetUsers(nil, username)
}

// UpdateUser - แก้ไขผู้ใช้
func (s *UserService) UpdateUser(user models.User) error {
	return s.UserRepo.UpdateUser(user)
}

// DeleteUser - ลบผู้ใช้
func (s *UserService) DeleteUser(userID int) error {
	return s.UserRepo.DeleteUser(userID)
}
