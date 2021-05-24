package user_service

import (
	"admin-server/models"
	user_repository "admin-server/repositories/user.repository"
)

func Create(user models.User) error {

	if err := user_repository.Create(user); err != nil {
		return err
	}

	return nil
}
