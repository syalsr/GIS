package config

type Config struct {
	GRPCAddr string `env:"GRPC_ADDR" envDefault:":9099"`
	PostgresURL string `env:"PG_DSN"`
}
