package usecase

import (
	"github.com/desvioow/goexpert-desafio-3/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(OrderRepository entity.OrderRepositoryInterface) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) ListAll() ([]OrderOutputDTO, error) {
	orders, err := l.OrderRepository.ListAll()

	if err != nil {
		return nil, err
	}

	var orderOutputDTOs []OrderOutputDTO
	for _, order := range orders {
		orderOutputDTOs = append(orderOutputDTOs, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return orderOutputDTOs, nil
}
