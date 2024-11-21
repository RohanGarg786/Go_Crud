package services

import (
	"Crud_fiber_Go/models"
	"errors"
)

func ValidateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return nil
}
