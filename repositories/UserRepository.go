package repositories

import (
	"errors"
	"praktikum/configs"
	"praktikum/models"
)

func CheckUserExistsRepositories(Request *models.Users) (*models.Users, error) {

	var Users models.Users

	results := configs.DB.Find(&Users, "email = ? AND password = ?", Request.Email, Request.Password)

	if results.RowsAffected == 0 {
		return nil, errors.New("Credentials Error")
	}

	if results.Error != nil {
		return nil, results.Error
	}

	return &Users, nil
}
