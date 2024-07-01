package inventoryHandler

import (
	"context"
	inventoryPb "github.com/Montheankul-K/game-microservices/modules/inventory/inventoryPb"
	"github.com/Montheankul-K/game-microservices/modules/inventory/inventoryUsecase"
)

type (
	InventoryGrpcHandler struct {
		inventoryPb.UnimplementedInventoryGrpcServiceServer
		inventoryUsecase inventoryUsecase.InventoryUsecase
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecase) *InventoryGrpcHandler {
	return &InventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}

func (g *InventoryGrpcHandler) IsAvailableToSell(ctx context.Context, req *inventoryPb.IsAvailableToSellReq) (*inventoryPb.IsAvailableToSellRes, error) {
	return nil, nil
}
