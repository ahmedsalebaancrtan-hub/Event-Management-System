package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/ahmedsaleban/eventManagementsystem/dtos"
	"github.com/ahmedsaleban/eventManagementsystem/models"
	"github.com/ahmedsaleban/eventManagementsystem/repository"
)

type EventSvc struct {
	Repo *repository.EventRepo
}

func RegistersvcRepo(Repo *repository.EventRepo) *EventSvc {
	return &EventSvc{
		Repo: Repo,
	}
}

func (svc *EventSvc) CreateEvent(data *dtos.CreateEventDTO) (int, error) {

	layout := "2006-01-02 15:04"

	start, err := time.Parse(layout, data.StartTime)
	if err != nil {
		return http.StatusBadRequest, errors.New("invalid start_time format. use YYYY-MM-DD HH:MM")
	}

	end, err := time.Parse(layout, data.EndTime)
	if err != nil {
		return http.StatusBadRequest, errors.New("invalid end_time format. use YYYY-MM-DD HH:MM")
	}

	if end.Before(start) {
		return http.StatusBadRequest, errors.New("end time cannot be Before start time ")
	}

	event := models.Event{
		Title:       data.Title,
		Type:        data.Type,
		Location:    data.Location,
		StartTime:   start,
		EndTime:     end,
		Capacity:    data.Capacity,
		Description: data.Description,
	}

	if err := svc.Repo.CreateEvent(event); err != nil {
		return http.StatusInternalServerError, err

	}
	return http.StatusCreated, nil
}

func (svc *EventSvc) Getall() (int, []models.Event, error) {

	events, err := svc.Repo.GetallEvents()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, events, nil
}

func (svc *EventSvc) GetEventById(id uint) (int, models.Event, error) {

	data, err := svc.Repo.FindEventById(id)

	if err != nil {
		return http.StatusInternalServerError, models.Event{}, err
	}

	return http.StatusOK, data, nil
}
