package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Message struct {
	From      int64                `json:"from"`
	Text      string               `json:"text"`
	Timestamp *timestamp.Timestamp `json:"timestamp"`
	Id        int64                `json:"id"`
	ChatId    int64                `json:"chat_id"`
}
