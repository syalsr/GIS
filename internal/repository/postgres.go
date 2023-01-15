package repository

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
	"github.com/syalsr/GIS/internal/config"
	"github.com/syalsr/GIS/internal/model"

	sq "github.com/Masterminds/squirrel"
)

const (
	startTransaction = "START TRANSACTION"
	commit           = "COMMIT"
	rollback         = "ROLLBACK"
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

// IsStrExist - check if stop is exist in db
func (p *Postgres) IsStrExist(ctx context.Context, table, column, value string) bool {
	log.Info().Msgf("Check if value: %s is in the table: %s, column: %s", value, table, column)

	sql, args, err := sq.Select(column).From(table).Where(sq.Eq{column: value}).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return true
	}
	sql, err = sq.Dollar.ReplacePlaceholders(sql)
	if err != nil {
		log.Err(err).Msg("cano replace the question mark with a dollar")
	}

	var existValue string
	err = p.Conn.QueryRow(ctx, sql, args...).Scan(&existValue)

	if err != nil {
		log.Err(err).Msgf("in table %s value %s not found", table, value)
		return false
	}
	return true
}

// CreateStop - create stop
func (p *Postgres) CreateStop(ctx context.Context, stop model.Stop) {
	log.Info().Msgf("Create stop %s %d %d", stop.Name, stop.Longitude, stop.Latitude)
	sql, args, err := sq.Insert("stop").Columns("name", "longitude", "latitude").Values(stop.Name, stop.Longitude, stop.Latitude).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return
	}
	sql, err = sq.Dollar.ReplacePlaceholders(sql)
	if err != nil {
		log.Err(err).Msg("cano replace the question mark with a dollar")
	}
	p.Conn.QueryRow(ctx, sql, args...)
}

// UpdateStop - update info stop
func (p *Postgres) UpdateStop(ctx context.Context, stop model.Stop) {
	log.Info().Msgf("Update stop %s %d %d", stop.Name, stop.Longitude, stop.Latitude)

	sql, args, err := sq.Update("stop").
		Set("longitude", stop.Longitude).
		Set("latitude", stop.Latitude).
		Where("name", stop.Name).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return
	}
	sql, err = sq.Dollar.ReplacePlaceholders(sql)
	if err != nil {
		log.Err(err).Msg("cano replace the question mark with a dollar")
	}
	p.Conn.QueryRow(ctx, sql, args...)
}

// CreateCurvature - insert curvature between stops in table curvature
func (p *Postgres) CreateCurvature(ctx context.Context, stop model.Stop, roadDistance model.RoadDistance) {
	p.Conn.QueryRow(ctx, startTransaction)

	stopIDFrom, err := p.GetIDByName(ctx, "stop", "stop_id", stop.Name)
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msgf("cant get stopIDFrom for: %s", stop.Name)
		return
	}

	stopIDTo, err := p.GetIDByName(ctx, "stop", "stop_id", roadDistance.Name)
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msgf("cant get stopIDTo for: %s", roadDistance.Name)
		return
	}

	log.Info().Msgf("Insert curvature between %d and %d - %d", stopIDFrom, stopIDTo, roadDistance.Curvature)
	insertCurvature, args, err := sq.Insert("curvature").Columns("stop_id_from", "stop_id_to", "length").Values(stopIDFrom, stopIDTo, roadDistance.Curvature).ToSql()
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msg("cant build sql")
		return
	}
	
	insertCurvature, err = sq.Dollar.ReplacePlaceholders(insertCurvature)
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msgf("cant replace the question mark with a dollar, %s", rollback)
	}
	p.Conn.QueryRow(ctx, insertCurvature, args...)

	p.Conn.QueryRow(ctx, commit)

}

// GetIDByName - get id from bus, stop tables
func (p *Postgres) GetIDByName(ctx context.Context, table, column, name string) (int, error) {
	log.Info().Msgf("Get id for table: %s, column: %s, name: %s", table, column, name)
	getStopID, args, err := sq.Select(column).From(table).Where(sq.Eq{"name": name}).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return -1, err
	}
	getStopID, err = sq.Dollar.ReplacePlaceholders(getStopID)
	if err != nil {
		log.Err(err).Msgf("cant replace the question mark with a dollar, %s", rollback)
		return -1, err
	}

	var id int
	err = p.Conn.QueryRow(ctx, getStopID, args...).Scan(&id)
	if err != nil {
		log.Err(err).Msg("cant scan value into id")
		return -1, err
	}
	log.Info().Msgf("stop id: %d", id)

	return id, nil
}

// CreateBus - create bus
func (p *Postgres) CreateBus(ctx context.Context, bus model.Bus) {
	p.Conn.QueryRow(ctx, startTransaction)

	sql, args, err := sq.Insert("bus").Columns("name", "is_roundtrip").Values(bus.Name, bus.IsRoundtrip).ToSql()
	if err != nil {
		log.Err(err).Msgf("cant build sql, %s", rollback)
		p.Conn.QueryRow(ctx, rollback)
		return
	}

	sql, err = sq.Dollar.ReplacePlaceholders(sql)
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msgf("cant replace the question mark with a dollar, %s", rollback)
	}

	p.Conn.QueryRow(ctx, sql, args...)

	for _, item := range bus.Stop {
		sql, args, err := sq.Insert("bus_stop").Columns("stop_name", "bus_name").Values(item, bus.Name).ToSql()
		if err != nil {
			p.Conn.QueryRow(ctx, rollback)
			log.Err(err).Msgf("cant build sql, %s", rollback)
		}

		sql, err = sq.Dollar.ReplacePlaceholders(sql)
		if err != nil {
			p.Conn.QueryRow(ctx, rollback)
			log.Err(err).Msgf("cant replace the question mark with a dollar, %s", rollback)
		}
		p.Conn.QueryRow(ctx, sql, args...)
	}
	p.Conn.QueryRow(ctx, commit)
}

// UpdateBus - update info bus
func (p *Postgres) UpdateBus(ctx context.Context, bus model.Bus) {
	sql, args, err := sq.Update("bus").Set("is_roundtrip", bus.IsRoundtrip).ToSql()
	if err != nil {
		log.Err(err).Msg("cant build sql")
		return
	}
	sql, err = sq.Dollar.ReplacePlaceholders(sql)
	if err != nil {
		log.Err(err).Msg("cant replace the question mark with a dollar")
	}
	p.Conn.QueryRow(ctx, sql, args...)
}

// CreateTrip - set bus route
func (p *Postgres) CreateTrip(ctx context.Context, stopName, busName string) {
	p.Conn.QueryRow(ctx, startTransaction)

	stopIDFrom, err := p.GetIDByName(ctx, "bus", "bus_id", stop.Name)
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msgf("cant get stopIDFrom for: %s", stop.Name)
		return
	}

	stopIDTo, err := p.GetIDByName(ctx, "stop", "stop_id", roadDistance.Name)
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msgf("cant get stopIDTo for: %s", roadDistance.Name)
		return
	}

	log.Info().Msgf("Insert curvature between %d and %d - %d", stopIDFrom, stopIDTo, roadDistance.Curvature)
	insertCurvature, args, err := sq.Insert("curvature").Columns("stop_id_from", "stop_id_to", "length").Values(stopIDFrom, stopIDTo, roadDistance.Curvature).ToSql()
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msg("cant build sql")
		return
	}
	
	insertCurvature, err = sq.Dollar.ReplacePlaceholders(insertCurvature)
	if err != nil {
		p.Conn.QueryRow(ctx, rollback)
		log.Err(err).Msgf("cant replace the question mark with a dollar, %s", rollback)
	}
	p.Conn.QueryRow(ctx, insertCurvature, args...)

	p.Conn.QueryRow(ctx, commit)
}