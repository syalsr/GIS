package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/syalsr/GIS/internal/config"
	"github.com/syalsr/GIS/internal/model"

	sq "github.com/Masterminds/squirrel"
)

// Postgres - type implement repository
type Postgres struct {
	Conn *pgxpool.Pool
}

// NewClient - create new obj which implement repository
func NewClient(cfg *config.App) Repository {
	log.Info().Msg("Create database connection")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	pool, err := pgxpool.Connect(ctx, cfg.PostgresURL)
	if err != nil {
		log.Fatal().Msgf("cant connect to postgres: %s", err)
	}

	return &Postgres{
		Conn: pool,
	}
}

// IsStopExist - check if stop is exist in db
func (p *Postgres) IsStrExist(ctx context.Context, table, column, value string) bool {
	sql, args, err := sq.Select(column).From(table).Where(sq.Eq{column: value}).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return true
	}
	var stop string
	err = p.Conn.QueryRow(ctx, sql, args).Scan(&stop)
	if err != nil {
		log.Err(err).Msg("cant exec scan in var stop")
	}

	log.Debug().Msgf("Stop name: %s", stop)

	return stop != ""
}

func (p *Postgres) CreateStop(ctx context.Context, stop model.Stop) {
	sql, args, err := sq.Insert("stop").Columns("name", "longitude", "latitude").Values(stop.Name, stop.Longitude, stop.Latitude).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return
	}

	p.Conn.QueryRow(ctx, sql, args)
}

// UpdateStop - update info stop
func (p *Postgres) UpdateStop(ctx context.Context, stop model.Stop) {
	sql, args, err := sq.Update("stop").
								Set("longitude", stop.Longitude).
								Set("latitude", stop.Latitude).
								Where("name", stop.Name).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return
	}

	p.Conn.QueryRow(ctx, sql, args)
}

func (p *Postgres) CreateBus(ctx context.Context, bus model.Bus) {
	sql, args, err := sq.Insert("bus").Columns("name", "is_roundtrip").Values(bus.Name,bus.IsRoundtrip).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return
	}

	p.Conn.QueryRow(ctx, sql, args)
}

// UpdateBus - update info bus
func (p *Postgres) UpdateBus(ctx context.Context, bus model.Bus) {
	sql, args, err := sq.Update("bus").Set("is_roundtrip", bus.IsRoundtrip).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return
	}

	p.Conn.QueryRow(ctx, sql, args)
}
