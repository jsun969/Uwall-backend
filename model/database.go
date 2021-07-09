package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Type      string         `json:"type"`
	FromName  string         `json:"fromName,omitempty"`
	FromSex   uint           `json:"fromSex,omitempty"`
	ToName    string         `json:"toName,omitempty"`
	ToSex     uint           `json:"toSex,omitempty"`
	Message   string         `json:"message"`
	Anonymous bool           `json:"anonymous,omitempty"`
	ImageUrl  string         `json:"imageUrl,omitempty"`
	Comments  []Comment      `json:"comments,omitempty"`
	Status    bool           `json:"-"`
}

type Comment struct {
	gorm.Model `json:"-"`
	MessageID  uint   `json:"-"`
	Name       string `json:"name,omitempty"`
	Comment    string `json:"comment"`
	Anonymous  bool   `json:"anonymous,omitempty"`
	Status     bool   `json:"-"`
}
