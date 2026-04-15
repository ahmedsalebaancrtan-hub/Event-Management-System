package handlers

import (
	"net/http"

	"github.com/ahmedsaleban/eventManagementsystem/dtos"
	"github.com/ahmedsaleban/eventManagementsystem/infra"
	"github.com/ahmedsaleban/eventManagementsystem/repository"
	"github.com/ahmedsaleban/eventManagementsystem/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Usersvc *service.UserService
}

func RegisterUserHandler() *UserHandler {
	userRepo := repository.RegisterRepo(infra.DB)
	usersvc := service.Registersvc(userRepo)

	return &UserHandler{
		Usersvc: usersvc,
	}

}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var body dtos.CreateUserdto
	err := c.ShouldBindBodyWithJSON(&body)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messege":    "failed to binding body",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}
	statuscode, err := h.Usersvc.CreateUser(&body)
	if err != nil {
		c.JSON(statuscode, gin.H{
			"messege":   err.Error(),
			"is_sucess": false,
		})
		return
	}

	c.JSON(statuscode, gin.H{
		"is_success": true,
		"messege":    "User Created sucessfully",
	})

}
