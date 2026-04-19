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

func NewUserService(Repo *repository.UserRepo) *UserService {
	return &UserService{
		Repo: Repo,
	}
}

// CREATE USER
func (svc *UserService) CreateUser(data *dtos.CreateUserdto) (int, error) {
	email := strings.ToLower(data.Email)
	_, err := svc.Repo.GetUserByEmail(email)
	if err == nil {
		slog.Error("User with that email already exists")
		return http.StatusConflict, errors.New("user with that email already exists")
	}

	hashbyte, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, errors.New(constants.DefaultErrorMsg)
	}

	err = svc.Repo.CreateUser(models.User{
		Name:     data.Name,
		Email:    email,
		Password: string(hashbyte),
		Role:     data.Role,
	})
	if err != nil {
		return http.StatusInternalServerError, errors.New("failed to create new user")
	}
	return http.StatusCreated, nil
}

// GET ALL USERS
func (svc *UserService) GetAllUsers() (int, []models.User, error) {
	data, err := svc.Repo.GetAllusers()
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, data, nil
}

// GET USER BY ID
func (svc *UserService) GetUserById(id uint) (int, models.User, error) {
	data, err := svc.Repo.GetUserbyId(id)
	if err != nil {
		return http.StatusNotFound, models.User{}, errors.New("user not found")
	}
	return http.StatusOK, data, nil
}

// WHO AM I
func (svc *UserService) WhoAmI(email string) (*models.User, int, error) {
	email = strings.ToLower(email)
	user, err := svc.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, http.StatusNotFound, errors.New("user not found")
	}
	return &user, http.StatusOK, nil
}

// LOGIN
func (svc *UserService) LoginUser(data *dtos.CreateLogindto) (dtos.LoginUserResponse, int, error) {
	email := strings.ToLower(data.Email)
	user, err := svc.Repo.GetUserByEmail(email)
	if err != nil {
		return dtos.LoginUserResponse{}, http.StatusUnauthorized, errors.New("invalid credentials")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return dtos.LoginUserResponse{}, http.StatusUnauthorized, errors.New("invalid credentials")
	}

	accessToken, err := helpers.GenerateJwt(user.Role, user.Email, time.Now().Add(30*time.Minute).Unix(), false)
	refreshToken, err2 := helpers.GenerateJwt(user.Role, user.Email, time.Now().Add(72*time.Hour).Unix(), true)

	if err != nil || err2 != nil {
		return dtos.LoginUserResponse{}, http.StatusInternalServerError, errors.New(constants.DefaultErrorMsg)
	}

	return dtos.LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         user,
	}, http.StatusOK, nil
}

// REFRESH TOKEN
func (svc *UserService) RefreshToken(email string) (*dtos.LoginUserResponse, int, error) {
	user, err := svc.Repo.GetUserByEmail(email)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New("unauthorized")
	}

	accessToken, _ := helpers.GenerateJwt(user.Role, user.Email, time.Now().Add(15*time.Minute).Unix(), false)
	refreshToken, _ := helpers.GenerateJwt(user.Role, user.Email, time.Now().Add(72*time.Hour).Unix(), true)

	return &dtos.LoginUserResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, http.StatusOK, nil
}

// FORGOT PASSWORD (Request Token)
func (svc *UserService) ForgotPassword(data *dtos.ForgotPasswordDTO) (int, error) {
	email := strings.ToLower(data.Email)
	_, err := svc.Repo.GetUserByEmail(email)
	if err != nil {
		return http.StatusOK, nil // Silent return for security
	}

	token := helpers.GenerateSecureToken()
	reset := models.PasswordResetToken{
		Email:     email,
		Token:     token,
		ExpiresAt: time.Now().Add(15 * time.Minute),
	}

	if err := svc.Repo.SaveResetToken(reset); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// RESET PASSWORD (using Token)
func (svc *UserService) ResetPasswords(data *dtos.ResetPasswordDTO) (int, error) {
	record, err := svc.Repo.GetResetToken(data.Token)
	if err != nil {
		return http.StatusBadRequest, errors.New("invalid or expired token")
	}

	if time.Now().After(record.ExpiresAt) {
		return http.StatusBadRequest, errors.New("token expired")
	}

	user, err := svc.Repo.GetUserByEmail(record.Email)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data.NewPassword), bcrypt.DefaultCost)
	if err := svc.Repo.UpdatePasswordById(user.ID, string(hash)); err != nil {
		return http.StatusInternalServerError, err
	}

	_ = svc.Repo.DeleteResetToken(data.Token)
	return http.StatusOK, nil
}
