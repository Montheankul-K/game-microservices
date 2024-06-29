package database

import (
	"context"
	"github.com/Montheankul-K/game-microservices/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url))
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("failed to pinging to database")
	}
	return client
}
