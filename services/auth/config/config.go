package config

import (
	"booking/libs/utils"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gopkg.in/yaml.v3"

	_ "github.com/lib/pq"
)

func LoadConfig() (*Config, error) {
	path, err := utils.GetPath("services/auth/config/config.yaml")
	if err != nil {
		return nil, errors.Wrap(err, "cant get path")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "cant open config file")
	}

	cfg := &Config{}

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "cant decode config file")
	}

	return cfg, nil
}

func NewPostgresDb(cfg *Config) (*sqlx.DB, error) {
	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DbName,
		cfg.Database.SslMode,
	)
	db, err := sqlx.Open("postgres", conn)
	if err != nil {
		return nil, errors.Wrap(err, "trying to conn to db")
	}

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "pinging db")
	}

	//need to be deleted asap
	table := `CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(225) NOT NULL,
		is_consultant BOOLEAN,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(table)
	if err != nil {
		return nil, err
	}
	// ...

	return db, nil
}

func NewRedisDb(cfg *Config) (*redis.Client, error) {
	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Cache.Host, cfg.Cache.Port),
		Password: cfg.Cache.Password,
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := db.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return db, nil
}
