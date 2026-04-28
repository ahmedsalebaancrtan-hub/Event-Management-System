package repository

import (
	"github.com/ahmedsaleban/eventManagementsystem/models"
	"gorm.io/gorm"
)

type EventRepo struct {
	DB *gorm.DB
}

func RegisterEventRepo(db *gorm.DB) *EventRepo {
	return &EventRepo{
		DB: db,
	}

}

func (r *EventRepo) CreateEvent(event models.Event) error {
	return r.DB.Create(&event).Error

}

func (r *EventRepo) GetallEvents() ([]models.Event, error) {
	var events []models.Event

	err := r.DB.Find(&events).Error
	if err != nil {
		return nil, err
	}

	return events, nil

}

func (r *EventRepo) FindEventById(id uint) (models.Event, error) {
	var event models.Event

	err := r.DB.Where("id = ?", id).First(&event).Error

	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}
