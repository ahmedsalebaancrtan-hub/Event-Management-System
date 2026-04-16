package dtos

import "github.com/ahmedsaleban/eventManagementsystem/models"

type CreateUserdto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=128"`
}

type CreateLogindto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=128"`
}

type LoginUserResponse struct {
	AccessToken  string      `json:"Access_token"`
	RefreshToken string      `json:"Refresh_token"`
	User         models.User `json:"user"`
}
