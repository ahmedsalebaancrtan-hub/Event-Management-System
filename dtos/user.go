package dtos

import "github.com/ahmedsaleban/eventManagementsystem/models"

type CreateUserdto struct {
	Name     string      `json:"name" binding:"required"`
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required,min=8,max=128"`
	Role     models.Role `json:"role" binding:"required,oneof=ADMIN ORGANIZER STAFF"`
}

type CreateLogindto struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=128"`
}

type LoginUserResponse struct {
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
	User         models.User `json:"user"`
}

// Standardized to ResetPasswordDTO
type ResetPasswordDTO struct {
	UserID      uint   `json:"user_id" binding:"required"`
	NewPassword string `json:"newpassword" binding:"required"`
	Token       string `json:"token"`
}

type ForgotPasswordDTO struct {
	Email string `json:"email" binding:"required,email"`
}
