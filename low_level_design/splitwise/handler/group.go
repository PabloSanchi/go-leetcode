package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"splitwise/domain/dto"
	"splitwise/service"
	"strconv"
)

type Group struct {
	groupService *service.Group
}

func NewGroupHandler(groupService *service.Group) *Group {
	return &Group{groupService: groupService}
}

func (gh *Group) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var group dto.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	groupId, err := gh.groupService.CreateGroup(&group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Location", "/api/v1/groups/"+strconv.Itoa(int(groupId)))
}

func (gh *Group) GetGroup(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	groupId, err := strconv.Atoi(strId)
	if err != nil {
		slog.Error(" error converting string to int", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	group := gh.groupService.GetGroupById(uint(groupId))
	if group == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	serializedGroup, err := json.Marshal(group)
	if err != nil {
		slog.Error("error serializing group", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(serializedGroup)
}

func (gh *Group) AddUsersToGroup(w http.ResponseWriter, r *http.Request) {
	strId := mux.Vars(r)["id"]
	groupId, err := strconv.Atoi(strId)
	if err != nil {
		slog.Error(" error converting string to int", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var GroupUsers dto.AddGroupUserRequest
	if err := json.NewDecoder(r.Body).Decode(&GroupUsers); err != nil {
		slog.Error("error decoding request body", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := gh.groupService.AddUsers(uint(groupId), GroupUsers.UserIds); err != nil {
		slog.Error("error adding users to group", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
