package storage

import (
	"context"
	"os"

	"github.com/TropicalDog17/alert-crud/internal/model"
	"github.com/redis/go-redis/v9"
)

type StorageInterface interface {
	AddAlert(ctx context.Context, req *model.CreateAlertRequest) error
	GetAlert(ctx context.Context, req *model.GetAlertRequest) (*model.Alert, error)
	GetAllAlerts(ctx context.Context) ([]*model.Alert, error)
	UpdateAlert(ctx context.Context, req *model.UpdateAlertRequest) error
	DeleteAlert(ctx context.Context, req *model.DeleteAlertRequest) error
	DeleteAllAlerts(ctx context.Context) error
	GetAllUserAlerts(ctx context.Context, userId string) ([]*model.Alert, error)
	GetDB() *redis.Client
	// Close() error
	NewStorage(r *redis.Client) *Storage
}

// Storage contains an SQL db. Storage implements the StorageInterface.
type Storage struct {
	DB *redis.Client
}

func (s *Storage) GetDB() *redis.Client {
	return s.DB
}

func NewRedisClient() (*Storage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ENDPOINT") + ":" + os.Getenv("REDIS_PORT"),
		DB:       0,
		Password: os.Getenv("REDIS_PASSWORD"),
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &Storage{
		DB: client,
	}, nil
}

func NewLocalRedisClient() (*Storage, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		DB:       0,
		Password: "",
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return &Storage{
		DB: client,
	}, nil
}

func (s *Storage) NewStorage(r *redis.Client) *Storage {
	return &Storage{
		DB: r,
	}
}
