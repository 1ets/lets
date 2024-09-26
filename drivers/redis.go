package drivers

import (
	"github.com/1ets/lets"
	"github.com/1ets/lets/types"

	"github.com/go-redis/redis/v8"
)

var RedisConfig types.IRedis

type redisProvider struct {
	dsn      string
	username string
	password string
	database int
	redis    *redis.Client
}

func (m *redisProvider) Connect() {
	m.redis = redis.NewClient(&redis.Options{
		Addr:     m.dsn,
		Username: m.username,
		Password: m.password,
		DB:       m.database,
	})
}

// Define MySQL service host and port
func Redis() {
	if RedisConfig == nil {
		return
	}

	lets.LogI("Redis Starting ...")

	redis := redisProvider{
		dsn:      RedisConfig.GetDsn(),
		username: RedisConfig.GetUsername(),
		password: RedisConfig.GetPassword(),
		database: RedisConfig.GetDatabase(),
	}
	redis.Connect()

	// Inject Gorm into repository
	for _, repository := range RedisConfig.GetRepositories() {
		repository.SetDriver(redis.redis)
	}
}
