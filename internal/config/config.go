package config

type App struct {
	GrpcAddr    string `env:"GRPC_ADDR" envDefault:":9099"`
	GrpcGateway string `env:"GRPC_GATEWAY" envDefault:":9098"`
	PostgresURL string `env:"PG_DSN" envDefault:"postgres://postgres:postgrespw@localhost:55000/postgres?sslmode=disable"`
	RedisURL    string `env:"REDIS_DSN" envDefault:"redis://default:redispw@localhost:55001"`
}
