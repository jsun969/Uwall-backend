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
	Status    bool
}
