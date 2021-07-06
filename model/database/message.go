package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	FromName  string
	FromSex   uint
	ToName    string
	ToSex     uint
	Message   string
	Anonymous bool
	Status    bool
}
