package repository

import (
	"gorm.io/gorm"
	models "splitwise/domain/db"
	"splitwise/domain/dto"
	"splitwise/util"
)

type UserRepositoryImpl struct {
	db   *gorm.DB
	util *util.Util
}

func NewUserRepositoryImpl(db *gorm.DB, util *util.Util) UserRepository {
	return &UserRepositoryImpl{
		db:   db,
		util: util,
	}
}

func (ur *UserRepositoryImpl) CreateUser(user *dto.User) (uint, error) {
	hashedPassword, err := ur.util.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	userEntity := &models.User{
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
	userEntity := &models.User{}
	ur.db.First(userEntity, id)

	return &dto.UserInfo{
		Id:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
}

func (ur *UserRepositoryImpl) GetUserByEmail(email string) *dto.UserInfo {
	userEntity := &models.User{}
	ur.db.Where("email = ?", email).First(userEntity)

	return &dto.UserInfo{
		Id:    userEntity.ID,
		Name:  userEntity.Name,
		Email: userEntity.Email,
	}
}

func (ur *UserRepositoryImpl) VerifyUser(user *dto.User) bool {
	userEntity := &models.User{}
	ur.db.Where("email = ?", user.Email).First(userEntity)

	return ur.util.VerifyPassword(userEntity.Password, user.Password)
}
