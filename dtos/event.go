package dtos

type CreateEventDTO struct {
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Location    string `json:"location" binding:"required"`
	StartTime   string `json:"start_time" binding:"required"`
	EndTime     string `json:"end_time" binding:"required"`
	Capacity    int    `json:"capacity" binding:"required"`
	Description string `json:"description"`
}
