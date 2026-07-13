package repository

import (
	"github.com/raghunandan/gin-crud/internal/config"
	"github.com/raghunandan/gin-crud/internal/models"
)

func GetAllUsers(page, limit int) ([]models.User, error) {
	var users []models.User

	offset := (page - 1) * limit

	result := config.DB.
		Offset(offset).
		Limit(limit).
		Find(&users)

	if result.Error != nil {
		return users, nil
	}

	return users, nil
}