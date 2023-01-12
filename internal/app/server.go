package app

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/syalsr/GIS/internal/app/servicegis"
	"github.com/syalsr/GIS/internal/config"

	"google.golang.org/grpc"

	api "github.com/syalsr/GIS/pkg/GIS-api/GIS/v1"
)

func Run(ctx context.Context, cfg *config.Config) error {

	_, cancel := context.WithCancel(ctx)

	listener, err := net.Listen("tcp", cfg.GRPCAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	server := grpc.NewServer()
	api.RegisterGISServer(server, &servicegis.GIS{})
	log.Printf("gRPC server listening at %v", listener.Addr())
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

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
