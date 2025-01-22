package chat

import (
	"context"
	"github.com/ViciousKit/course-chat-server/internal/repository"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{db: db}
}

func (r repo) CreateChat(ctx context.Context, userIds []int64) (int64, error) {
	rows := r.db.QueryRow(ctx, "INSERT INTO chats (user_ids) VALUES ($1) RETURNING id", userIds)

	var lastId int64
	err := rows.Scan(&lastId)
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (r repo) DeleteChat(ctx context.Context, chatId int64) error {
	_, err := r.db.Exec(ctx, "DELETE from chats WHERE id = $1", chatId)

	if err != nil {
		return err
	}

	return nil
}

func (r repo) SendMessage(ctx context.Context, from int64, text string, chatId int64, timestamp *timestamp.Timestamp) error {
	_, err := r.db.Exec(ctx, "INSERT INTO messages (from, text, chat_id, created_at) VALUES ($1, $2, $3, $4)", from, text, chatId, timestamp)

	if err != nil {
		return err
	}

	return nil
}

func (r repo) AddUsers(ctx context.Context, userIds []int64, chatId int64) error {

	return nil
}

func (r repo) RemoveUsers(ctx context.Context, userIds []int64, chatId int64) error {

	return nil
}

func (r repo) GetMessages(ctx context.Context, chatId int64, lastId int64, limit int64, offset int64) error {

	return nil
}
