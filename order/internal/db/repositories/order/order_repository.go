package order

import (
	"context"
	"github.com/bludot/tempmee/order/internal/db"
)

type OrderRepositoryImpl interface {
	FindAll(ctx context.Context) (*[]Order, error)
	FindById(ctx context.Context, id int) (*Order, error)
	FindAllByUserId(ctx context.Context, userId string) (*[]Order, error)
	FindAllByParentOrderId(ctx context.Context, parentOrderId int) (*[]Order, error)
	Create(ctx context.Context, order *Order) (*Order, error)
	Update(ctx context.Context, order *Order) (*Order, error)
	Delete(ctx context.Context, id int) error
}

type OrderRepository struct {
	db *db.DB
}

func NewOrderRepository(db *db.DB) OrderRepositoryImpl {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) FindAll(ctx context.Context) (*[]Order, error) {
	var orders []Order
	err := r.db.DB.Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (r *OrderRepository) FindById(ctx context.Context, id int) (*Order, error) {
	var order Order
	err := r.db.DB.Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) FindAllByUserId(ctx context.Context, userId string) (*[]Order, error) {
	var orders []Order
	err := r.db.DB.Where("user_id = ?", userId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (r *OrderRepository) FindAllByParentOrderId(ctx context.Context, parentOrderId int) (*[]Order, error) {
	var orders []Order
	err := r.db.DB.Where("parent_order_id = ?", parentOrderId).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return &orders, nil
}

func (r *OrderRepository) Create(ctx context.Context, order *Order) (*Order, error) {
	err := r.db.DB.Create(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) Update(ctx context.Context, order *Order) (*Order, error) {

	var orderToUpdate Order
	err := r.db.DB.Where("id = ?", order.ID).First(&orderToUpdate).Error
	if err != nil {
		return nil, err
	}
	err = r.db.DB.Model(&orderToUpdate).Updates(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *OrderRepository) Delete(ctx context.Context, id int) error {
	err := r.db.DB.Delete(&Order{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
