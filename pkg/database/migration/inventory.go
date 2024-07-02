package migration

import (
	"context"
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func inventoryDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("inventory_db")
}

func InventoryMigrate(pctx context.Context, cfg *config.Config) {
	db := inventoryDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("players_inventory")
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"player_id", 1}, {"item_id", 1}}},
	})
	for _, index := range indexs {
		log.Printf("index: %s", index)
	}

	col = db.Collection("player_inventory_queue")
	results, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("migration inventory success: ", results)
}
