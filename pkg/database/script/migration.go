package main

import (
	"context"
	"github.com/Montheankul-K/game-microservices/config"
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
	case "auth":
	case "item":
	case "inventory":
	case "payment":
	}
}
