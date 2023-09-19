package order_test

import (
	orderRepository "github.com/bludot/tempmee/order/internal/db/repositories/order"
	"github.com/bludot/tempmee/order/internal/services/order"
	"github.com/bludot/tempmee/order/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestOrderService_CreateOrder(t *testing.T) {
	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(&orderRepository.Order{
			ID:            1,
			ParentOrderID: nil,
			UserID:        "user_id",
			Price:         1000,
			Quantity:      1,
		}, nil)

		orderService := order.NewOrderService(orderRepo)

		order, err := orderService.CreateOrder(nil, &orderRepository.Order{
			ID:            1,
			ParentOrderID: nil,
			UserID:        "user_id",
			Price:         1000,
			Quantity:      1,
		})

		a.NoError(err)
		a.Equal(1, order.ID)
		a.Equal("user_id", order.UserID)
		a.Equal(int64(1000), order.Price)
		a.Equal(int64(1), order.Quantity)
	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().Create(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

		orderService := order.NewOrderService(orderRepo)

		order, err := orderService.CreateOrder(nil, &orderRepository.Order{
			ID:            1,
			ParentOrderID: nil,
			UserID:        "user_id",
			Price:         1000,
			Quantity:      1,
		})

		a.Error(err)
		a.Nil(order)
	})

}

func TestOrderService_FindOrderById(t *testing.T) {
	t.Run("should return order", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(&orderRepository.Order{
			ID:            1,
			ParentOrderID: nil,
			UserID:        "user_id",
			Price:         1000,
			Quantity:      1,
		}, nil)

		orderService := order.NewOrderService(orderRepo)

		order, err := orderService.FindOrderById(nil, 1)

		a.NoError(err)
		a.Equal(1, order.ID)
		a.Equal("user_id", order.UserID)
		a.Equal(int64(1000), order.Price)
		a.Equal(int64(1), order.Quantity)
	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().FindById(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

		orderService := order.NewOrderService(orderRepo)

		order, err := orderService.FindOrderById(nil, 1)

		a.Error(err)
		a.Nil(order)
	})
}

func TestOrderService_FindOrdersByUserId(t *testing.T) {
	t.Run("should return orders", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().FindAllByUserId(gomock.Any(), gomock.Any()).Return(&[]orderRepository.Order{
			{
				ID:            1,
				ParentOrderID: nil,
				UserID:        "user_id",
				Price:         1000,
				Quantity:      1,
			},
			{
				ID:            2,
				ParentOrderID: nil,
				UserID:        "user_id",
				Price:         1000,
				Quantity:      1,
			},
		}, nil)

		orderService := order.NewOrderService(orderRepo)

		orders, err := orderService.FindOrdersByUserId(nil, "user_id")

		a.NoError(err)
		a.Equal(2, len(*orders))
		a.Equal(1, (*orders)[0].ID)
		a.Equal(2, (*orders)[1].ID)
	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().FindAllByUserId(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

		orderService := order.NewOrderService(orderRepo)

		orders, err := orderService.FindOrdersByUserId(nil, "user_id")

		a.Error(err)
		a.Nil(orders)
	})

}

func TestOrderService_FindOrdersByParentOrderId(t *testing.T) {
	t.Run("should return orders", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().FindAllByParentOrderId(gomock.Any(), gomock.Any()).Return(&[]orderRepository.Order{
			{
				ID:            1,
				ParentOrderID: nil,
				UserID:        "user_id",
				Price:         1000,
				Quantity:      1,
			},
			{
				ID:            2,
				ParentOrderID: nil,
				UserID:        "user_id",
				Price:         1000,
				Quantity:      1,
			},
		}, nil)

		orderService := order.NewOrderService(orderRepo)

		orders, err := orderService.FindOrdersByParentOrderId(nil, 1)

		a.NoError(err)
		a.Equal(2, len(*orders))
		a.Equal(1, (*orders)[0].ID)
		a.Equal(2, (*orders)[1].ID)
	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().FindAllByParentOrderId(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

		orderService := order.NewOrderService(orderRepo)

		orders, err := orderService.FindOrdersByParentOrderId(nil, 1)

		a.Error(err)
		a.Nil(orders)
	})
}

func TestOrderService_FindAllOrders(t *testing.T) {
	t.Run("should return orders", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().FindAll(gomock.Any()).Return(&[]orderRepository.Order{
			{
				ID:            1,
				ParentOrderID: nil,
				UserID:        "user_id",
				Price:         1000,
				Quantity:      1,
			},
			{
				ID:            2,
				ParentOrderID: nil,
				UserID:        "user_id",
				Price:         1000,
				Quantity:      1,
			},
		}, nil)

		orderService := order.NewOrderService(orderRepo)

		orders, err := orderService.FindAllOrders(nil)

		a.NoError(err)
		a.Equal(2, len(*orders))
		a.Equal(1, (*orders)[0].ID)
		a.Equal(2, (*orders)[1].ID)
	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().FindAll(gomock.Any()).Return(nil, assert.AnError)

		orderService := order.NewOrderService(orderRepo)

		orders, err := orderService.FindAllOrders(nil)

		a.Error(err)
		a.Nil(orders)
	})

}

func TestOrderService_UpdateOrder(t *testing.T) {
	t.Run("should return order", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(&orderRepository.Order{
			ID:            1,
			ParentOrderID: nil,
			UserID:        "user_id",
			Price:         1000,
			Quantity:      1,
		}, nil)

		orderService := order.NewOrderService(orderRepo)

		order, err := orderService.UpdateOrder(nil, &orderRepository.Order{
			ID:            1,
			ParentOrderID: nil,
			UserID:        "user_id",
			Price:         1000,
			Quantity:      1,
		})

		a.NoError(err)
		a.Equal(1, order.ID)
		a.Equal("user_id", order.UserID)
		a.Equal(int64(1000), order.Price)
		a.Equal(int64(1), order.Quantity)
	})

	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)

		orderService := order.NewOrderService(orderRepo)

		order, err := orderService.UpdateOrder(nil, &orderRepository.Order{
			ID:            1,
			ParentOrderID: nil,
			UserID:        "user_id",
			Price:         1000,
			Quantity:      1,
		})

		a.Error(err)
		a.Nil(order)
	})

}

func TestOrderService_DeleteOrder(t *testing.T) {
	t.Run("should return error if error", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orderRepo := mocks.NewMockOrderRepositoryImpl(ctrl)
		orderRepo.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(assert.AnError)

		orderService := order.NewOrderService(orderRepo)

		err := orderService.DeleteOrder(nil, 1)

		a.Error(err)
	})

}
