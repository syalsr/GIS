package main

import (
	"github.com/syalsr/GIS/internal/app"
	"github.com/syalsr/GIS/internal/config"

	"context"
	"log"

	"github.com/caarlos0/env/v6"
)

func main() {
	cfg := &config.App{}

	if err := env.Parse(cfg); err != nil {
		log.Fatalf("failed to retrieve env variables, %v", err)
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		log.Fatal("error running grpc server ", err)
	}
}
