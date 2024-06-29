package main

import (
	"context"
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/pkg/database"
	"github.com/Montheankul-K/game-microservices/server"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal(".env path is required")
		}
		return os.Args[1]
	}())

	db := database.DbConn(ctx, cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, cfg, db)
}
