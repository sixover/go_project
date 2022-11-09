package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	MyselfId uint
	yourId   uint
	Type     int
	Desc     string
}

func (table *Contact) TableName() string {
	return "contact"
}
