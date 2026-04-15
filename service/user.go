package service

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/ahmedsaleban/eventManagementsystem/dtos"
	"github.com/ahmedsaleban/eventManagementsystem/models"
	"github.com/ahmedsaleban/eventManagementsystem/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func Registersvc(Repo *repository.UserRepo) *UserService {
	return &UserService{
		Repo: Repo,
	}
}

func (svc *UserService) CreateUser(data *dtos.CreateUserdto) (int, error) {
	email := strings.ToLower(data.Email)
	_, err := svc.Repo.GetUserByEmail(email)
	if err == nil {
		slog.Error("User with that email already exist")
		return http.StatusConflict, errors.New("User with that email already exist")
	}

	err = svc.Repo.CreateUser(models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		slog.Error("failed created User")
		return http.StatusInternalServerError, errors.New("failed to create new User")
	}
	return http.StatusCreated, nil
}
