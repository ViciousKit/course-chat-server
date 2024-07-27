package models

import (
	"time"
)

type Message struct {
	From      int64      `json:"from"`
	Text      string     `json:"text"`
	Timestamp *time.Time `json:"timestamp"`
	Id        int64      `json:"id"`
	ChatId    int64      `json:"chat_id"`
}
