package storage

import (
	"context"
	"database/sql"
	"fmt"

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

func (s *Storage) Create(ctx context.Context, userNames []string) (int64, error) {

	return 1, nil
}

func (s *Storage) Delete(ctx context.Context, id int64) error {
	return nil

}

func (s *Storage) SendMessage(ctx context.Context, from string, text string, timestamp int64) error {

	return nil
}
