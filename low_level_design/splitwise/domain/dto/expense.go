package dto

type Share struct {
	UserId uint    `json:"user_id"`
	Amount float64 `json:"amount"`
}

type Expense struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	PayerId     uint    `json:"payer_id"`
	GroupId     uint    `json:"group_id"`
	Shares      []Share `json:"shares"`
}
