package migration

import (
	"context"
	"github.com/Montheankul-K/game-microservices/config"
	"github.com/Montheankul-K/game-microservices/modules/player"
	"github.com/Montheankul-K/game-microservices/pkg/database"
	"github.com/Montheankul-K/game-microservices/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func playerDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.DbConn(pctx, cfg).Database("player_db")
}

func PlayerMigrate(pctx context.Context, cfg *config.Config) {
	db := playerDbConn(pctx, cfg)
	defer db.Client().Disconnect(pctx)

	col := db.Collection("player_transactions")
	indexs, _ := col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"id", 1}}},
		{Keys: bson.D{{"player_id", 1}}},
	})
	for _, index := range indexs {
		log.Printf("index: %s", index)
	}

	col = db.Collection("players")
	indexs, _ = col.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"id", 1}}},
		{Keys: bson.D{{"email", 1}}},
	})
	for _, index := range indexs {
		log.Printf("index: %s", index)
	}

	documents := func() []any {
		players := []*player.Player{
			{
				Email:    "player001@gmail.com",
				Password: "password",
				Username: "player001",
				PlayerRoles: []player.Role{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "player002@gmail.com",
				Password: "password",
				Username: "player002",
				PlayerRoles: []player.Role{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "player003@gmail.com",
				Password: "password",
				Username: "player003",
				PlayerRoles: []player.Role{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
			{
				Email:    "admin001@gmail.com",
				Password: "password",
				Username: "admin001",
				PlayerRoles: []player.Role{
					{
						RoleTitle: "player",
						RoleCode:  0,
					},
					{
						RoleTitle: "admin",
						RoleCode:  1,
					},
				},
				CreatedAt: utils.LocalTime(),
				UpdatedAt: utils.LocalTime(),
			},
		}

		docs := make([]any, len(players))
		for _, r := range players {
			docs = append(docs, r)
		}
		return docs
	}()

	results, err := col.InsertMany(pctx, documents, nil)
	if err != nil {
		panic(err)
	}
	log.Println("migration players success: ", results)

	playerTransactions := make([]any, len(results.InsertedIDs))
	for _, p := range results.InsertedIDs {
		playerTransactions = append(playerTransactions, &player.Transaction{
			PlayerId:  "player:" + p.(primitive.ObjectID).Hex(),
			Amount:    1000,
			CreatedAt: utils.LocalTime(),
		})
	}

	col = db.Collection("player_transactions")
	results, err = col.InsertMany(pctx, playerTransactions, nil)
	if err != nil {
		panic(err)
	}
	log.Println("migration player_transactions success: ", results)

	col = db.Collection("player_transactions_queue")
	result, err := col.InsertOne(pctx, bson.M{"offset": -1}, nil)
	if err != nil {
		panic(err)
	}
	log.Println("migration player_transactions_queue success: ", result)
}
