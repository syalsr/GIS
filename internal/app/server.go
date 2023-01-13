package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"github.com/syalsr/GIS/internal/app/servicegis"
	"github.com/syalsr/GIS/internal/config"
	api "github.com/syalsr/GIS/pkg/GIS-api/GIS/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Run - func which run grpc and grpc-gateway server
func Run(ctx context.Context, cfg *config.Config) error {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	listener, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		log.Err(err).Msgf("cant connected to %s", cfg.GrpcAddr)
	}

	server := grpc.NewServer()

	gtw := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = api.RegisterGISHandlerFromEndpoint(ctx, gtw, cfg.GrpcAddr, opts)
	if err != nil {
		log.Err(err).Msg("cant register handlers")
	}

	api.RegisterGISServer(server, servicegis.NewGrcpGIS())
	go func() {
		if err = server.Serve(listener); err != nil {
			log.Fatal().Msgf("cant start gRPC server: %w", err)
		}
	}()

	go func() {
		if err = http.ListenAndServe(cfg.GrpcGateway, gtw); err != nil {
			log.Fatal().Msgf("cant start gRPC-gateway server: %w", err)
		}
	}()

	gracefulShutDown(server, cancel)

	return nil
}

func gracefulShutDown(s *grpc.Server, cancel context.CancelFunc) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

LOOP:
	for {
		select {
		case <-ch:
			break LOOP
		default:
		}
	}

	s.GracefulStop()
	cancel()
}
