package db

type Client interface {
	DB() DB
	Close() error
}

type DB interface {
	Close() error
}
