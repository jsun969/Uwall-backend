package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	Type      string
	FromName  string
	FromSex   uint
	ToName    string
	ToSex     uint
	Message   string
	Anonymous bool
	ImageUrl  string
	Comments  []Comment
	Status    bool
}

type Comment struct {
	gorm.Model
	MessageID uint
	Name      string
	Comment   string
	Anonymous bool
}
