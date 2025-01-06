package repository

import (
	"gorm.io/gorm"
	"log/slog"
	models "splitwise/domain/db"
	"splitwise/domain/dto"
)

type GroupRepositoryImpl struct {
	db *gorm.DB
}

func NewGroupRepositoryImpl(db *gorm.DB) GroupRepository {
	return &GroupRepositoryImpl{
		db: db,
	}
}

func (gr *GroupRepositoryImpl) CreateGroup(group *dto.Group) (uint, error) {
	users := gr.fetchUsers(group.UserIds)
	groupEntity := &models.Group{
		Name:        group.Name,
		Description: group.Description,
		Users:       users,
	}

	if err := gr.db.Create(groupEntity).Error; err != nil {
		return 0, err
	}
	return groupEntity.ID, nil
}

func (gr *GroupRepositoryImpl) GetGroupById(id uint) (*dto.Group, error) {
	groupEntity := &models.Group{}
	if err := gr.db.Preload("Users", func(db *gorm.DB) *gorm.DB { return db.Select("id") }).First(groupEntity, id).Error; err != nil {
		slog.Error("group not found", "group_id", id)
		return nil, err
	}

	var userIds []uint
	for _, user := range groupEntity.Users {
		userIds = append(userIds, user.ID)
	}

	return &dto.Group{
		Id:          groupEntity.ID,
		Name:        groupEntity.Name,
		Description: groupEntity.Description,
		UserIds:     userIds,
	}, nil
}

func (gr *GroupRepositoryImpl) AddUsersToGroup(groupId uint, userIds []uint) error {
	groupEntity := &models.Group{}
	if err := gr.db.Preload("Users").First(groupEntity, groupId).Error; err != nil {
		slog.Error("group not found", "group_id", groupId)
		return err
	}

	users := gr.fetchUsers(userIds)
	if err := gr.db.Model(groupEntity).Association("Users").Append(users); err != nil {
		slog.Error("error adding users to group", "error", err)
		return err
	}

	return nil
}

func (gr *GroupRepositoryImpl) fetchUsers(userIds []uint) []*models.User {
	var users []*models.User
	gr.db.Where(len(userIds) > 0, "id IN ?", userIds).Find(&users)
	return users
}
