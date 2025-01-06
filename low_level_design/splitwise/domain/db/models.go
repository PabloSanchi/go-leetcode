package models

import "gorm.io/gorm"

// User represents a user in the system.
type User struct {
	gorm.Model
	Name     string   `gorm:"not null"`
	Email    string   `gorm:"unique;not null"`
	Password string   `gorm:"not null"`
	Groups   []*Group `gorm:"many2many:user_groups;"`
}

// Group represents a group of users and their shared expenses.
type Group struct {
	gorm.Model
	Name        string `gorm:"unique;not null"`
	Description string
	Users       []*User    `gorm:"many2many:user_groups;"`
	Expenses    []*Expense `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
}

// Expense represents an expense made by a user in a group.
// It contains the total amount and the shares of each user that has to pay for the expense.
type Expense struct {
	gorm.Model
	Description string          `gorm:"not null"`
	Amount      float64         `gorm:"not null"`
	GroupID     uint            `gorm:"not null;index"`
	Group       Group           `gorm:"foreignKey:GroupID;constraint:OnDelete:CASCADE;"`
	PayerID     uint            `gorm:"not null;index"`
	Payer       User            `gorm:"foreignKey:PayerID;constraint:OnDelete:CASCADE;"`
	Shares      []*ExpenseShare `gorm:"foreignKey:ExpenseID;constraint:OnDelete:CASCADE;"`
}

// ExpenseShare represents how much a specific user owes for a specific expense.
type ExpenseShare struct {
	UserID    uint    `gorm:"primaryKey;not null;index"`
	User      User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	ExpenseID uint    `gorm:"primaryKey;not null;index"`
	Expense   Expense `gorm:"foreignKey:ExpenseID;constraint:OnDelete:CASCADE;"`
	Amount    float64 `gorm:"not null"`
	IsSettled bool    `gorm:"not null"`
}
