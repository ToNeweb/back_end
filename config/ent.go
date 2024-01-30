package config

import (
	"context"
	"log"
	"server04/ent"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var (
	client      *ent.Client
	redisClient *redis.Client
)

func GetClients() (*ent.Client, *redis.Client) {
	return client, redisClient
}

func SetClients(newClient *ent.Client, newRedisClient *redis.Client) {
	client = newClient
	redisClient = newRedisClient
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

func NewEntClient() (*ent.Client, error) {

	client, err := ent.Open("postgres", "dbname=videosharing2  sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	return client, err
}
