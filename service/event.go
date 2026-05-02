package service

import (
	"errors"
	"net/http"
	"strings"
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

	if data.ImgUrl != "" && !strings.HasPrefix(data.ImgUrl, "http") {
		return http.StatusBadRequest, errors.New("invalid image url")
	}

	event := models.Event{
		Title:       data.Title,
		Type:        data.Type,
		Location:    data.Location,
		StartTime:   start,
		EndTime:     end,
		Capacity:    data.Capacity,
		ImgUrl:      data.ImgUrl,
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

func (svc *EventSvc) UpdateEvent(id uint, data *dtos.UpdateEventDTO) (int, error) {

	event, err := svc.Repo.FindEventById(id)
	if err != nil {
		return http.StatusNotFound, errors.New("event not found")
	}

	layout := "2006-01-02 15:04"

	// Update fields only if provided

	if data.Title != nil {
		event.Title = *data.Title
	}

	if data.Type != nil {
		event.Type = *data.Type
	}

	if data.Location != nil {
		event.Location = *data.Location
	}

	if data.Description != nil {
		event.Description = *data.Description
	}

	if data.ImgURL != nil {
		if !strings.HasPrefix(*data.ImgURL, "http") {
			return http.StatusBadRequest, errors.New("invalid image url")
		}
		event.ImgUrl = *data.ImgURL
	}

	// Handle time updates carefully
	var start = event.StartTime
	var end = event.EndTime

	if data.StartTime != nil {
		parsedStart, err := time.Parse(layout, *data.StartTime)
		if err != nil {
			return http.StatusBadRequest, errors.New("invalid start_time format")
		}
		start = parsedStart
	}

	if data.EndTime != nil {
		parsedEnd, err := time.Parse(layout, *data.EndTime)
		if err != nil {
			return http.StatusBadRequest, errors.New("invalid end_time format")
		}
		end = parsedEnd
	}

	// Validate after updates
	if end.Before(start) {
		return http.StatusBadRequest, errors.New("end_time cannot be before start_time")
	}

	event.StartTime = start
	event.EndTime = end

	if data.Capacity != nil {
		if *data.Capacity <= 0 {
			return http.StatusBadRequest, errors.New("capacity must be greater than 0")
		}
		event.Capacity = *data.Capacity
	}

	if err := svc.Repo.UpdateEvent(event); err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
