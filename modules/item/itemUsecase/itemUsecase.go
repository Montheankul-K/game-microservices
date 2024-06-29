package itemUsecase

import "github.com/Montheankul-K/game-microservices/modules/item/itemRepository"

type (
	ItemUsecase interface{}

	itemUsecase struct {
		itemRepository itemRepository.ItemRepository
	}
)

func NewItemUsecase(itemRepository itemRepository.ItemRepository) ItemUsecase {
	return &itemUsecase{
		itemRepository: itemRepository,
	}
}
