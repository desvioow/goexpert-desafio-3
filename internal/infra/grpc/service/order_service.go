package service

import (
	"context"
	"github.com/desvioow/goexpert-desafio-3/internal/infra/grpc/pb"
	"github.com/desvioow/goexpert-desafio-3/internal/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase   usecase.CreateOrderUseCase
	ListAllOrdersUseCase usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listAllOrdersUsecase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase:   createOrderUseCase,
		ListAllOrdersUseCase: listAllOrdersUsecase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := usecase.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *emptypb.Empty) (*pb.ListOrdersResponse, error) {
	orders, err := s.ListAllOrdersUseCase.ListAll()
	if err != nil {
		return nil, err
	}

	allOrders := make([]*pb.Order, 0, len(orders))
	for _, order := range orders {
		allOrders = append(allOrders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}
	return &pb.ListOrdersResponse{
		Orders: allOrders,
	}, nil
}
