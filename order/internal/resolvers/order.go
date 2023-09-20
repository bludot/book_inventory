package resolvers

import (
	"context"
	"github.com/bludot/tempmee/order/graph/model"
	orderRepository "github.com/bludot/tempmee/order/internal/db/repositories/order"
	"github.com/bludot/tempmee/order/internal/services/order"
	"strconv"
)

func GetOrderByID(ctx context.Context, orderService order.OrderServiceImpl, id string) (*model.Order, error) {
	orderid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	order, err := orderService.FindOrderById(ctx, orderid)

	if err != nil {
		return nil, err
	}

	subOrders, err := orderService.FindOrdersByParentOrderId(ctx, order.ID)
	if err != nil {
		return nil, err
	}

	var products []string
	for _, subOrder := range *subOrders {
		products = append(products, *subOrder.ProductId)
	}

	return &model.Order{
		ID:       strconv.Itoa(order.ID),
		Products: products,
		UserID:   order.UserID,
		Status:   model.OrderStatus(order.Status),
	}, nil

}

func GetOrdersByUserID(ctx context.Context, orderService order.OrderServiceImpl, userID string) ([]*model.Order, error) {
	ordersByUserId, err := orderService.FindOrdersByUserId(ctx, userID)

	if err != nil {
		return nil, err
	}

	orderItems, err := MapOrdersAndSubOrders(ctx, *ordersByUserId)
	if err != nil {
		return nil, err
	}

	return orderItems, nil

}

func CreateOrder(ctx context.Context, orderService order.OrderServiceImpl, orderInput model.CreateOrderInput) (*model.Order, error) {

	order := orderRepository.Order{
		UserID: orderInput.UserID,
		Status: model.OrderStatusCreated.String(),
	}

	createdOrder, err := orderService.CreateOrder(ctx, &order)

	var orders []orderRepository.Order

	// for each product id create new order
	for index, productID := range orderInput.Products {
		productid := productID
		orderItem := orderRepository.Order{
			UserID:        orderInput.UserID,
			ParentOrderID: &createdOrder.ID,
			Status:        model.OrderStatusUnknown.String(),
			ProductId:     &productid,
			Quantity:      int64(orderInput.Quantity[index]),
		}

		if err != nil {
			return nil, err
		}
		orders = append(orders, orderItem)
	}

	for _, order := range orders {
		_, err := orderService.CreateOrder(ctx, &order)
		if err != nil {
			return nil, err
		}
	}

	return &model.Order{
		ID:       strconv.Itoa(createdOrder.ID),
		Products: orderInput.Products,
		UserID:   createdOrder.UserID,
		Status:   model.OrderStatus(createdOrder.Status),
	}, nil
}

func Map[T any, U any](arr []T, f func(T) U) []U {
	var newArr []U
	for _, v := range arr {
		newArr = append(newArr, f(v))
	}
	return newArr
}

func MapOrdersAndSubOrders(ctx context.Context, orders []orderRepository.Order) ([]*model.Order, error) {
	orderItems := make([]*model.Order, 0)
	for _, order := range orders {
		if order.ParentOrderID == nil {
			orderItems = append(orderItems, &model.Order{
				ID:       strconv.Itoa(order.ID),
				Products: make([]string, 0),
				UserID:   order.UserID,
				Status:   model.OrderStatus(order.Status),
			})

		}
	}

	for _, order := range orders {
		if order.ParentOrderID != nil {
			// add to products array to relevant parent order
			for _, orderItem := range orderItems {
				if orderItem.ID == strconv.Itoa(*order.ParentOrderID) {
					orderItem.Products = append(orderItem.Products, *order.ProductId)
				}
			}

		}
	}

	return orderItems, nil
}

func UpdateOrderStatus(ctx context.Context, orderService order.OrderServiceImpl, id string, status model.OrderStatus) (*model.Order, error) {
	orderid, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	order, err := orderService.FindOrderById(ctx, orderid)
	if err != nil {
		return nil, err
	}

	order.Status = status.String()

	_, err = orderService.UpdateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return GetOrderByID(ctx, orderService, id)

}
