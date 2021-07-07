package model

type from struct {
	Name string `json:"name"`
	Sex  uint   `json:"sex"`
}

type to struct {
	Name string `json:"name"`
	Sex  uint   `json:"sex"`
}

type ReqLove struct {
	From      from   `json:"from"`
	To        to     `json:"to" binding:"required"`
	Message   string `json:"message" binding:"required"`
	Anonymous bool   `json:"anonymous"`
	ImageUrl  string `json:"imageUrl"`
}
