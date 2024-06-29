package itemRepository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ItemRepository interface{}

	itemRepository struct {
		db *mongo.Client
	}
)

func NewItemRepository(db *mongo.Client) ItemRepository {
	return &itemRepository{
		db: db,
	}
}

func (r *itemRepository) itemDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("item_db")
}
