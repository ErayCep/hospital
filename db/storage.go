package db

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Storage struct {
	DB    *gorm.DB
	Redis *redis.Client
}

var HospitalStorage Storage

func (s *Storage) GetDB() *gorm.DB {
	return s.DB
}

func (s *Storage) GetRedisClient() *redis.Client {
	return s.Redis
}
