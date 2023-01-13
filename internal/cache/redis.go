package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/mediocregopher/radix/v4"
	"github.com/rs/zerolog/log"
	"github.com/syalsr/GIS/internal/config"
)

const (
	getRoadTrip string = "SMEMBERS"
	addRoadTrip string = "SADD"
)

// Interface - cache
type Interafce interface {
	SetRoadTrip(ctx context.Context, from, to string, stops []string) error
	GetRoadTrip(ctx context.Context, from, to string) ([]string, error)
}

// RedisClient - client for redis
type RedisClient struct {
	client radix.Client
}

// NewCache creates new cache
func NewCache(cfg *config.App) (Interafce, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	client, err := (radix.PoolConfig{}).New(ctx, "tcp", cfg.RedisURL)
	if err != nil {
		return nil, fmt.Errorf("can't connect to redis, %w", err)
	}

	return &RedisClient{client: client}, nil
}

// SetRoadTrip - add road trip to redis
func (r *RedisClient) SetRoadTrip(ctx context.Context, from, to string, stops []string) error {
	err := r.client.Do(ctx, radix.FlatCmd(nil, addRoadTrip, from+to, stops))
	if err != nil {
		log.Err(err).Msg("cant add road trip to redis")
		return fmt.Errorf("cant add road trip to redis: %w", err)
	}
	return nil
}

// GetRoadTrip - get road trip from redis
func (r *RedisClient) GetRoadTrip(ctx context.Context, from, to string) (road []string, err error) {
	err = r.client.Do(ctx, radix.Cmd(&road, getRoadTrip, from+to))
	if err != nil {
		log.Err(err).Msg("cant get road trip from redis")
		return nil, fmt.Errorf("cant get road trip from redis: %w", err)
	}
	return
}
