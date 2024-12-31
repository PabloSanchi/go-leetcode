package repository

import "splitwise/domain/dto"

type UserRepository interface {
	CreateUser(user *dto.User) (uint, error)
	GetUserById(id uint) *dto.UserInfo
	GetUserByEmail(email string) *dto.UserInfo
	VerifyUser(user *dto.User) bool
}
