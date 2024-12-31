package service

import (
	"splitwise/domain/dto"
	"splitwise/repository"
)

type User struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *User {
	return &User{userRepository: userRepository}
}

func (u *User) GetUserByEmail(email string) *dto.UserInfo {
	return u.userRepository.GetUserByEmail(email)
}
