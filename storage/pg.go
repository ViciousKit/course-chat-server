package storage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/lib/pq"
)

type Storage struct {
	db *sql.DB
}

func New(user string, password string, dbname string, host string, port int) *Storage {
	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Cant connect pg" + err.Error())
		panic(err)
	}
	if err := db.Ping(); err != nil {
		fmt.Println("Cant ping pg" + err.Error())
		panic(err)
	}
	fmt.Println("Connected!")

	return &Storage{db: db}
}

func (s *Storage) CreateChat(ctx context.Context, userIds []int64) (int64, error) {
	stmt, err := s.db.Prepare("INSERT INTO chats (user_ids) VALUES ($1)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.ExecContext(ctx, userIds)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (s *Storage) DeleteChat(ctx context.Context, chatId int64) error {
	stmt, err := s.db.Prepare("DELETE from chats WHERE id = $1")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, chatId)
	if err != nil {
		return err
	}

	return nil

}

func (s *Storage) SendMessage(ctx context.Context, from int64, text string, chatId int64, timestamp *timestamp.Timestamp) error {

	stmt, err := s.db.Prepare("INSERT INTO messages (from, text, chat_id, created_at) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, from, text, chatId, timestamp)
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
