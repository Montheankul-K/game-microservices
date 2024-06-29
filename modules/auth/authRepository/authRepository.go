package authRepository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	AuthRepository interface{}

	authRepository struct {
		db *mongo.Client
	}
)

func NewAuthRepository(db *mongo.Client) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) authDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("auth_db")
}
