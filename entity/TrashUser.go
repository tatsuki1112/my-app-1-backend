package entity

import "gorm.io/gorm"

type TrashUser struct {
	gorm.Model
	Name string
	Email *string
}