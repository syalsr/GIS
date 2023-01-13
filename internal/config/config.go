package config

type Config struct {
	GrpcAddr    string `env:"GRPC_ADDR" envDefault:":9099"`
	GrpcGateway string `env:"GRPC_GATEWAY" envDefault:":9098"`
	PostgresURL string `env:"PG_DSN"`
}
