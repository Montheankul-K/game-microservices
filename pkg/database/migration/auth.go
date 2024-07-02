package migration

import (
	"context"
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/modules/auth"
	"github.com/Montheankul-K/game-microservices/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func authDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("auth_db")
}

func AuthMigrate(pctx context.Context, cfg *config.Config) {
	db := authDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("auth")
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"id", 1}}},
		{Keys: bson.D{{"player_id", 1}}},
		{Keys: bson.D{{"refresh_token", 1}}},
	})
	for _, index := range indexs {
		log.Printf("index: %s", index)
	}

	col = db.Collection("roles")
	indexs, _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"id", 1}}},
		{Keys: bson.D{{"code", 1}}},
	})
	for _, index := range indexs {
		log.Printf("index: %s", index)
	}

	documents := func() []any {
		roles := []*auth.Role{
			{Title: "player", Code: 0},
			{Title: "admin", Code: 1},
		}

		docs := make([]any, len(roles))
		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("migration auth success: ", results)
}
