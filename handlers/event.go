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

type EventHandler struct {
	EventSvc *service.EventSvc
}

func RegisterEventHandler() *EventHandler {

	eventRepo := repository.RegisterEventRepo(infra.DB)
	EventSvc := service.RegistersvcRepo(eventRepo)

	return &EventHandler{
		EventSvc: EventSvc,
	}
}

func (h *EventHandler) CreateEvent(c *gin.Context) {
	var body dtos.CreateEventDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"messege":    err.Error(),
			"is_success": false,
		})
		return
	}

	status, err := h.EventSvc.CreateEvent(&body)

	if err != nil {
		c.JSON(status, gin.H{
			"messege":    err.Error(),
			"is_success": false,
		})
		return

	}

	c.JSON(status, gin.H{"message": "event created successfully"})
}

func (h *EventHandler) Getall(c *gin.Context) {
	status, event, err := h.EventSvc.Getall()

	if err != nil {
		c.JSON(status, gin.H{
			"is_success": false,
			"messege":    err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"is_sucess": true,
		"messege":   "events fecthed sucessfully!",
		"data":      event,
	})

}

func (h *EventHandler) FindEventByid(c *gin.Context) {
	IdStr := c.Param("event_id")

	id, err := strconv.Atoi(IdStr)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"messege":    "failed to get  Classid param",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	status, event, err := h.EventSvc.GetEventById(uint(id))

	if err != nil {
		c.JSON(status, gin.H{
			"is_success": false,
			"messege":    err.Error(),
		})
		return
	}

	c.JSON(status, gin.H{
		"is_sucess": true,
		"messege":   "event fecthed sucessfully!",
		"data":      event,
	})

}
