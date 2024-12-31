package repository

import (
	"gorm.io/gorm"
	"splitwise/domain/db"
	"splitwise/domain/dto"
	"splitwise/util"
)

type UserRepositoryImpl struct {
	db   *gorm.DB
	util *util.Util
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db:   db,
		util: util.NewUtil(),
	}
}

func (ur *UserRepositoryImpl) CreateUser(user *dto.User) (uint, error) {
	hashedPassword, err := ur.util.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	userEntity := &db.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: hashedPassword,
	}

	if err := ur.db.Create(userEntity).Error; err != nil {
		return 0, err
	}
	return userEntity.ID, nil
}

func (ur *UserRepositoryImpl) GetUserById(id uint) *dto.UserInfo {
	userEntity := &db.User{}
	ur.db.First(userEntity, id)

	return &dto.UserInfo{
		Id:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
}

func (ur *UserRepositoryImpl) GetUserByEmail(email string) *dto.UserInfo {
	userEntity := &db.User{}
	ur.db.Where("email = ?", email).First(userEntity)

	return &dto.UserInfo{
		Id:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
}

func (ur *UserRepositoryImpl) VerifyUser(user *dto.User) bool {
	userEntity := &db.User{}
	ur.db.Where("email = ?", user.Email).First(userEntity)

	return ur.util.VerifyPassword(userEntity.Password, user.Password)
}
