package model

type Base struct {
	Name      string `json:"name"`
	Message   string `json:"message" binding:"required"`
	Anonymous bool   `json:"anonymous"`
	ImageUrl  string `json:"imageUrl"`
}
