package model

type ReqComment struct {
	ID        uint   `json:"id" binding:"required"`
	Name      string `json:"name"`
	Comment   string `json:"comment" binding:"required"`
	Anonymous bool   `json:"anonymous"`
}
