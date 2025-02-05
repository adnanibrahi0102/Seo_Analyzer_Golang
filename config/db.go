package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var (
	DB          *gorm.DB      // PostgreSQL connection
	RedisClient *redis.Client // Redis connection
	Ctx         = context.Background() // Changed ctx to Ctx to make it exported
)

// ConnectToDB initializes the PostgreSQL database connection
func ConnectToDB() {
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("DB_URL is not set in the environment variables")
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")
}

// ConnectToRedis initializes the Redis connection

func ConnectToRedis() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		log.Fatal("REDIS_URL is not set in the environment variables")
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr: redisURL,
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	fmt.Println("✅ Connected to Redis successfully!")

}
