package inventoryHandler

import "github.com/Montheankul-K/game-microservices/modules/inventory/inventoryUsecase"

type (
	InventoryGrpcHandler struct {
		inventoryUsecase inventoryUsecase.InventoryUsecase
	}
)

func NewInventoryGrpcHandler(inventoryUsecase inventoryUsecase.InventoryUsecase) *InventoryGrpcHandler {
	return &InventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
