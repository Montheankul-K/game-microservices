package inventoryUsecase

import "github.com/Montheankul-K/game-microservices/modules/inventory/inventoryRepository"

type (
	InventoryUsecase interface{}

	inventoryUsecase struct {
		inventoryRepository inventoryRepository.InventoryRepository
	}
)

func NewInventoryUsecase(inventoryRepository inventoryRepository.InventoryRepository) InventoryUsecase {
	return &inventoryUsecase{
		inventoryRepository: inventoryRepository,
	}
}
