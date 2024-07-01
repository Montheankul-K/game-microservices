package itemHandler

import (
	"context"
	itemPb "github.com/Montheankul-K/game-microservices/modules/item/itemPb"
	"github.com/Montheankul-K/game-microservices/modules/item/itemUsecase"
)

type (
	ItemGrpcHandler struct {
		itemPb.UnimplementedItemGrpcServiceServer
		itemUsecase itemUsecase.ItemUsecase
	}
)

func NewItemGrpcHandler(itemUsecase itemUsecase.ItemUsecase) *ItemGrpcHandler {
	return &ItemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}

func (g *ItemGrpcHandler) FindItemInIds(ctx context.Context, req *itemPb.FindItemInIdsReq) (*itemPb.FindItemInIdsRes, error) {
	return nil, nil
}
