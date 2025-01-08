package service

import (
	"splitwise/domain/dto"
	"splitwise/repository"
)

type Group struct {
	groupRepository repository.GroupRepository
}

func NewGroupService(groupRepository repository.GroupRepository) *Group {
	return &Group{groupRepository: groupRepository}
}

func (g *Group) CreateGroup(group *dto.Group) (uint, error) {
	return g.groupRepository.CreateGroup(group)
}

func (g *Group) GetGroupById(id uint) *dto.Group {
	group, _ := g.groupRepository.GetGroupById(id)
	return group
}

func (g *Group) AddUsers(groupId uint, userIds []uint) error {
	return g.groupRepository.AddUsers(groupId, userIds)
}
