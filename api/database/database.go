package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbNo int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DB_ADDRESS"),
		Password: os.Getenv("DB_PASS"),
		DB:       dbNo,
	})
	return rdb
}
func StoreShortenedURL(shortURL, originalURL string, expiration time.Duration) error {
	rdb := CreateClient(0)
	defer rdb.Close()

	// Store the URL and visits count as a Redis hash
	err := rdb.HSet(Ctx, shortURL, "url", originalURL, "visits", 0).Err()
	if err != nil {
		return err
	}

	// Set expiration time for the key
	err = rdb.Expire(Ctx, shortURL, expiration).Err()
	return err
}

func GetOriginalURL(shortURL string) (string, error) {
	rdb := CreateClient(0)
	defer rdb.Close()

	url, err := rdb.HGet(Ctx, shortURL, "url").Result()
	fmt.Printf("url: %s, err: %v", url, err)
	if err != nil {
		return "", err
	}

	rdb.HIncrBy(Ctx, shortURL, "visits", 1)

	return url, nil
}
