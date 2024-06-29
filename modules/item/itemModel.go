package item

import "github.com/Montheankul-K/game-microservices/modules/model"

type (
	CreateItemReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		Damage   int     `json:"damage" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
	}

	ShowCase struct {
		ItemId   string `json:"item_id"`
		Title    string `json:"title"`
		Price    string `json:"price"`
		Damage   int    `json:"damage"`
		ImageUrl string `json:"image_url"`
	}

	SearchReq struct {
		Title string `json:"title"`
		model.PaginateReq
	}

	UpdateItemReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		Damage   int     `json:"damage" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
	}

	EnableOrDisableItemReq struct {
		UsageStatus bool `json:"status"`
	}
)
