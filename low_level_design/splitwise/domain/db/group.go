package db

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Id    uint    `gorm:"primaryKey;autoIncrement"`
	Name  string  `gorm:"unique;not null"`
	Users []*User `gorm:"many2many:group_users;"`
}
