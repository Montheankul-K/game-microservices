package itemHandler

import "github.com/Montheankul-K/game-microservices/modules/item/itemUsecase"

type (
	ItemGrpcHandler struct {
		itemUsecase itemUsecase.ItemUsecase
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecase) *ItemGrpcHandler {
	return &ItemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}
