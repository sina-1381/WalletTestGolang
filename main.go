package main

import (
	"log"

	"WalletTest/internal"
	"WalletTest/pkg/app_error"
	"WalletTest/pkg/config"
	"WalletTest/pkg/database"
	"WalletTest/pkg/server"
	"WalletTest/proto"
)

func main() {
	// =============== CONFIG SETUP ===============
	env := config.NewEnv()                                            // ENV
	logger, err := config.NewLogger(env.LOG_FOLDER, env.APP_LOCATION) // LOG
	if err != nil {
		log.Fatal(app_error.LogFile)
	}
	defer logger.Close()
	postgresql, err := database.NewPostgresql(env.PostgresqlConfig()) // POSTGRESQL
	if err != nil {
		log.Fatal(app_error.PostgresqlConnection)
	}
	defer postgresql.Connection.Close()
	redis, err := database.NewRedis(env.RedisConfig()) // REDIS
	if err != nil {
		log.Fatal(app_error.RedisConnection)
	}
	defer redis.Close()
	if err := postgresql.MigrationRestart(); err != nil { // MIGRATION
		log.Fatal(err)
	}
	//memoryCache := config.NewMemoryCache()

	// =============== DEPENDENCIES ===============
	wallet := internal.NewWallet()
	grpc, err := server.NewGrpc(env.APP_PORT)
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterWalletServiceServer(grpc.Server, wallet)

	// =============== SERVER ===============
	grpc.GracefulClose()
	if err := grpc.Start(); err != nil {
		log.Fatal(err)
	}
}
