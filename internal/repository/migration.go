package repository

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
	"github.com/syalsr/GIS/internal/config"
)

func Migrate(c *config.App) {
	fileMigrations := "file://migrations"

	m, err := migrate.New(fileMigrations, c.PostgresURL)
	if err != nil {
		log.Fatal().Msgf("cant create migrate instance: %w", err)
	}
	
	if err := m.Up(); err != nil {
		log.Fatal().Msgf("cant up migrate: %s", err.Error())
	}
}