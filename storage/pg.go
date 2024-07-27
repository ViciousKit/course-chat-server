package storage

import (
	"context"
	"github.com/jackc/pgx/v5"

	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateChat(ctx context.Context, userIds []int64) (int64, error) {
	rows := s.db.QueryRow(ctx, "INSERT INTO chats (user_ids) VALUES ($1) RETURNING id", userIds)

	var lastId int64
	err := rows.Scan(&lastId)
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Storage) DeleteChat(ctx context.Context, chatId int64) error {
	_, err := s.db.Exec(ctx, "DELETE from chats WHERE id = $1", chatId)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) SendMessage(ctx context.Context, from int64, text string, chatId int64, timestamp *timestamp.Timestamp) error {
	_, err := s.db.Exec(ctx, "INSERT INTO messages (from, text, chat_id, created_at) VALUES ($1, $2, $3, $4)", from, text, chatId, timestamp)

	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddUsers(ctx context.Context, userIds []int64, chatId int64) error {

	return nil
}

func (s *Storage) RemoveUsers(ctx context.Context, userIds []int64, chatId int64) error {

	return nil
}

func (s *Storage) GetMessages(ctx context.Context, chatId int64, lastId int64, limit int64, offset int64) error {

	return nil
}
