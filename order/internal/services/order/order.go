package order

import (
	"context"
	orderRepository "github.com/bludot/tempmee/order/internal/db/repositories/order"
)

type OrderServiceImpl interface {
	FindAllOrders(ctx context.Context) (*[]orderRepository.Order, error)
	FindOrderById(ctx context.Context, id int) (*orderRepository.Order, error)
	FindOrdersByUserId(ctx context.Context, userId string) (*[]orderRepository.Order, error)
	FindOrdersByParentOrderId(ctx context.Context, parentOrderId int) (*[]orderRepository.Order, error)
	CreateOrder(ctx context.Context, order *orderRepository.Order) (*orderRepository.Order, error)
	UpdateOrder(ctx context.Context, order *orderRepository.Order) (*orderRepository.Order, error)
	DeleteOrder(ctx context.Context, id int) error
}

type OrderService struct {
	Repo orderRepository.OrderRepositoryImpl
}

func NewOrderService(repo orderRepository.OrderRepositoryImpl) OrderServiceImpl {
	return &OrderService{
		Repo: repo,
	}
}

func (o *OrderService) FindAllOrders(ctx context.Context) (*[]orderRepository.Order, error) {
	return o.Repo.FindAll(ctx)
}

func (o *OrderService) FindOrderById(ctx context.Context, id int) (*orderRepository.Order, error) {
	return o.Repo.FindById(ctx, id)
}

func (o *OrderService) FindOrdersByUserId(ctx context.Context, userId string) (*[]orderRepository.Order, error) {
	return o.Repo.FindAllByUserId(ctx, userId)
}

func (o *OrderService) FindOrdersByParentOrderId(ctx context.Context, parentOrderId int) (*[]orderRepository.Order, error) {
	return o.Repo.FindAllByParentOrderId(ctx, parentOrderId)
}

func (o *OrderService) CreateOrder(ctx context.Context, order *orderRepository.Order) (*orderRepository.Order, error) {
	return o.Repo.Create(ctx, order)
}

func (o *OrderService) UpdateOrder(ctx context.Context, order *orderRepository.Order) (*orderRepository.Order, error) {
	return o.Repo.Update(ctx, order)
}

func (o *OrderService) DeleteOrder(ctx context.Context, id int) error {
	return o.Repo.Delete(ctx, id)
}
