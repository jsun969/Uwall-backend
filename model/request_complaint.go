package model

type Complaint struct {
	Name      string `json:"name"`
	Message   string `json:"message" binding:"required"`
	Anonymous bool   `json:"anonymous"`
}
