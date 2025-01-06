package repository

import "splitwise/domain/dto"

type GroupRepository interface {
	CreateGroup(group *dto.Group) (uint, error)
	GetGroupById(id uint) (*dto.Group, error)
	AddUsersToGroup(groupId uint, userIds []uint) error
}
