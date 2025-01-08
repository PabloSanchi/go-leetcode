package dto

type Group struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserIds     []uint `json:"user_ids"`
}

type AddGroupUserRequest struct {
	UserIds []uint `json:"user_ids"`
}
