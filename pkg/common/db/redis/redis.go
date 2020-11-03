package redis

import (
	"github.com/go-redis/redis"
)

// NewConnection creates a new connection to mySQL DB
func NewConnection(connectionURL string) (*redis.Client, error) {
	conn := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	res := conn.Ping()

	return conn, res.Err()
}
