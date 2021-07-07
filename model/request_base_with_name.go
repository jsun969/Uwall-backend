package model

type BaseWithName struct {
	Name      string `json:"name" binding:"required"`
	Message   string `json:"message" binding:"required"`
	ImageUrl  string `json:"imageUrl"`
}
