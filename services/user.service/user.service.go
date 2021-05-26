package user_service

import (
	"admin-server/models"
	ur "admin-server/repositories/user.repository"
)

var repository *ur.Repository

func init() {
	repository = ur.GetRepository()
}

func Create(user *models.User) error {
	return repository.Create(user)
}

func GetUserById(id string) (*models.User, error) {
	return repository.GetUserById(id)
}

func GetUsers() (*models.Users, error) {
	return repository.GetUsers()
}
