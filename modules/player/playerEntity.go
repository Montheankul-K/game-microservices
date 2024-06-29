package player

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type (
	Player struct {
		Id         primitive.ObjectID `json:"id" bson:"id,omitempty"`
		Email      string             `json:"email" bson:"email"`
		Password   string             `json:"password" bson:"password"`
		Username   string             `json:"username" bson:"username"`
		CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
		PlayerRole []Role             `bson:"player_role"`
	}

	Role struct {
		RoleTitle string `json:"role_title" bson:"role_title"`
		RoleCode  int    `json:"role_code" bson:"role_code"`
	}

	ProfileBson struct {
		Id        primitive.ObjectID `json:"id" bson:"id,omitempty"`
		Email     string             `json:"email" bson:"email"`
		Username  string             `json:"username" bson:"username"`
		CreatedAt time.Time          `json:"created_at" bson:"created_at"`
		UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	}

	SavingAccount struct {
		PlayerId string  `json:"player_id" bson:"player_id"`
		Balance  float64 `json:"balance" bson:"balance"`
	}
)
