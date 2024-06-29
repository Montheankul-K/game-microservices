package inventory

import (
	"github.com/Montheankul-K/game-microservices/modules/item"
	"github.com/Montheankul-K/game-microservices/modules/model"
)

type (
	UpdateInventoryReq struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}

	ItemInventory struct {
		InventoryId string `json:"inventory_id"`
		*item.ShowCase
	}

	PlayerInventory struct {
		PlayerId string `json:"player_id"`
		*model.PaginateRes
	}
)
