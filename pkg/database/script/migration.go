package main

import (
	"context"
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/pkg/database/migration"
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
	_ = ctx
	switch cfg.App.Name {
	case "player":
		migration.PlayerMigrate(ctx, cfg)
	case "auth":
		migration.AuthMigrate(ctx, cfg)
	case "item":
		migration.ItemMigrate(ctx, cfg)
	case "inventory":
		migration.InventoryMigrate(ctx, cfg)
	case "payment":
		migration.PaymentMigrate(ctx, cfg)
	}
}
