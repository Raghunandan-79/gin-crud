package services

import (
	"github.com/raghunandan/gin-crud/internal/models"
	"github.com/raghunandan/gin-crud/internal/repository"
)

func GetAllUsers(page, limit int) ([]models.User, error) {
	return repository.GetAllUsers(page, limit)
}