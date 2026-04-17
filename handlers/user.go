package handlers

import (
	"net/http"
	"strconv"

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

func (h *UserHandler) LoginUser(c *gin.Context) {
	var req dtos.CreateLogindto

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
		return
	}

	response, statusCode, err := h.Usersvc.LoginUser(&req)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Success response
	c.JSON(statusCode, gin.H{
		"message": "Login successful",
		"data":    response,
	})
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	status, data, err := h.Usersvc.GetAllUsers()

	if err != nil {
		c.JSON(status, gin.H{
			"is_success": false,
			"messege":    err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"is_sucess": true,
		"messege":   "Users fecthed sucessfully!",
		"data":      data,
	})

}

func (h *UserHandler) GetUserById(c *gin.Context) {
	IdStr := c.Param("userId")

	id, err := strconv.Atoi(IdStr)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"messege":    "failed to get  user param",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	status, user, err := h.Usersvc.GetUserById(uint(id))

	if err != nil {
		c.JSON(status, gin.H{
			"is_success": false,
			"messege":    err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"is_sucess": true,
		"messege":   "users fecthed sucessfully!",
		"data":      user,
	})

}

func (h *UserHandler) WhoAmI(c *gin.Context) {
	email, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	user, status, err := h.Usersvc.WhoAmI(email.(string))
	if err != nil {
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(status, gin.H{"user": user})
}
