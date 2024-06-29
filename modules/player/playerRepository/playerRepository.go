package playerRepository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	PlayerRepository interface{}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepository {
	return &playerRepository{
		db: db,
	}
}

func (r *playerRepository) playerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}
