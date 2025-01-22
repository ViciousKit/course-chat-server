package model

import "time"

type Message struct {
	AuthorId  int64
	Text      string
	Timestamp time.Time
	Id        int64
	ChatId    int64
}
