package models

type Chat struct {
	Id      int64   `json:"id"`
	UserIds []int64 `json:"user_ids"`
}
