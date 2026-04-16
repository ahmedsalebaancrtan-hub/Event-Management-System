package repository

import (
	"github.com/ahmedsaleban/eventManagementsystem/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func RegisterRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (r *UserRepo) CreateUser(data models.User) error {
	return r.DB.Create(&data).Error
}

func (r *UserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r *UserRepo) GetAllusers() ([]models.User, error) {
	var User []models.User

	err := r.DB.Find(&User).Error

	if err != nil {
		return nil, err
	}
	return User, nil

}

func (r *UserRepo) GetUserbyId(id uint) (models.User, error) {
	var user models.User

	err := r.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
