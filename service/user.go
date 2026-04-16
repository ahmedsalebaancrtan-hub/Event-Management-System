package service

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/ahmedsaleban/eventManagementsystem/constants"
	"github.com/ahmedsaleban/eventManagementsystem/dtos"
	"github.com/ahmedsaleban/eventManagementsystem/helpers"
	"github.com/ahmedsaleban/eventManagementsystem/models"
	"github.com/ahmedsaleban/eventManagementsystem/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepo
}

//REGISTER FUNCTION SERVICE

func Registersvc(Repo *repository.UserRepo) *UserService {
	return &UserService{
		Repo: Repo,
	}
}

//CREATE USER SERVICE

func (svc *UserService) CreateUser(data *dtos.CreateUserdto) (int, error) {
	email := strings.ToLower(data.Email)
	_, err := svc.Repo.GetUserByEmail(email)
	if err == nil {
		slog.Error("User with that email already exist")
		return http.StatusConflict, errors.New("User with that email already exist")
	}
	slog.Info("hashing password")

	hashbyte, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password")
		return http.StatusInternalServerError, errors.New(constants.DefaultErrorMsg)
	}
	data.Password = string(hashbyte)

	slog.Info("created user")

	err = svc.Repo.CreateUser(models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	})
	if err != nil {
		slog.Error("failed created User")
		return http.StatusInternalServerError, errors.New("failed to create new User")
	}
	return http.StatusCreated, nil
}

// GET ALL USERS  API
func (svc *UserService) GetAllUsers() (int, []models.User, error) {

	data, err := svc.Repo.GetAllusers()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, data, nil
}

// GET USER BY ID API

func (svc *UserService) GetUserById(id uint) (int, models.User, error) {

	data, err := svc.Repo.GetUserbyId(id)

	if err != nil {
		return http.StatusInternalServerError, models.User{}, err
	}

	return http.StatusOK, data, nil
}

// LOGIN USER API

func (svc *UserService) LoginUser(data *dtos.CreateLogindto) (response dtos.LoginUserResponse, statusCode int, err error) {
	slog.Info("Login user")

	email := strings.ToLower(data.Email)

	// Find user
	user, err := svc.Repo.GetUserByEmail(email)
	if err != nil {
		slog.Error("user not found")
		statusCode = http.StatusUnauthorized
		err = errors.New("invalid credentials")
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		slog.Error("invalid password")
		statusCode = http.StatusUnauthorized
		err = errors.New("invalid credentials")
		return
	}
	AccessToken, err := helpers.GenerateJwt(user.Role, user.Email, time.Now().Add(30*time.Minute).Unix(), false)

	if err != nil {
		slog.Error("Failed to Generate access token")
		statusCode = http.StatusInternalServerError
		err = errors.New(constants.DefaultErrorMsg)

		return
	}
	RefreshToken, err := helpers.GenerateJwt(user.Role, user.Email, time.Now().Add(72*time.Hour).Unix(), true)

	if err != nil {
		slog.Error("Failed to Generate refresh token token")
		statusCode = http.StatusInternalServerError
		err = errors.New(constants.DefaultErrorMsg)

		return
	}

	response = dtos.LoginUserResponse{
		AccessToken:  AccessToken,
		RefreshToken: RefreshToken,
		User:         user,
	}

	statusCode = http.StatusOK
	return
}
