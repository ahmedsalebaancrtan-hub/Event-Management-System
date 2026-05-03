package repository

import (
	"github.com/ahmedsaleban/eventManagementsystem/dtos"
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

func (r *EventRepo) UpdateEvent(event models.Event) error {
	return r.DB.Save(&event).Error
}

func (r *EventRepo) FilterEvents(filter dtos.EventFilterDTO) ([]models.Event, error) {

	var events []models.Event
	query := r.DB.Model(&models.Event{})

	if filter.Location != "" {
		query = query.Where("LOWER(location) = LOWER(?)", filter.Location)
	}

	if filter.Type != "" {
		query = query.Where("LOWER(type) = LOWER(?)", filter.Type)
	}

	if filter.Search != "" {
		query = query.Where("LOWER(title) LIKE LOWER(?)", "%"+filter.Search+"%")
	}

	if filter.StartDate != "" {
		query = query.Where("start_time >= ?", filter.StartDate)
	}

	if filter.EndDate != "" {
		query = query.Where("start_time <= ?", filter.EndDate)
	}

	err := query.Find(&events).Error
	return events, err
}
