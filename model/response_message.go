package model

type APIMessage struct {
	ID        uint   `json:"id"`
	Type      string `json:"type"`
	FromName  string `json:"fromName"`
	FromSex   uint   `json:"fromSex"`
	ToName    string `json:"toName"`
	ToSex     uint   `json:"toSex"`
	Message   string `json:"message"`
	Anonymous bool   `json:"anonymous"`
	ImageUrl  string `json:"imageUrl"`
}
