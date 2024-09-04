package usecase

import (
	"github.com/desvioow/goexpert-desafio-3/internal/entity"
)

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	//OrderCreated    events.EventInterface
	//EventDispatcher events.EventDispatcherInterface
}

func NewCreateOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	// OrderCreated events.EventInterface,
	// EventDispatcher events.EventDispatcherInterface,
) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: OrderRepository,
		//OrderCreated:    OrderCreated,
		//EventDispatcher: EventDispatcher,
	}
}
