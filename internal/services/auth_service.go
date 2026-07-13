package services

import (
	"errors"

	"github.com/raghunandan/gin-crud/internal/dto"
	"github.com/raghunandan/gin-crud/internal/models"
	"github.com/raghunandan/gin-crud/internal/repository"
	"github.com/raghunandan/gin-crud/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func Signup(body dto.SignupRequest) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(body.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	user := models.User{
		Name: body.Name,
		Email: body.Email,
		Password: string(hashedPassword),
	}

	err = repository.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Login(body dto.LoginRequest) (string, error) {
	user, err := repository.GetUserByEmail(body.Email)

	if err != nil {
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(body.Password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetProfile(id uint) (*models.User, error) {
	return repository.GetUserByID(id)
}