package models

import "github.com/golang/protobuf/ptypes/timestamp"

type Message struct {
	From      string               `json:"from"`
	Text      string               `json:"text"`
	Timestamp *timestamp.Timestamp `json:"timestamp"`
}
