package resolvers_test

import (
	"context"
	"github.com/bludot/tempmee/order/graph/model"
	orderRepository "github.com/bludot/tempmee/order/internal/db/repositories/order"
	"github.com/bludot/tempmee/order/internal/resolvers"
	"github.com/bludot/tempmee/order/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestGetOrdersByUserID(t *testing.T) {
	t.Run("should return all orders by user id", func(t *testing.T) {
		a := assert.New(t)
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		ctx := context.Background()

		orderService := mocks.NewMockOrderServiceImpl(ctrl)
		orderService.EXPECT().FindOrdersByUserId(gomock.Any(), "1").Return(&[]orderRepository.Order{
			{
				ID:     1,
				UserID: "1",
				Status: "CREATED",
			},
			{
				ID:            2,
				UserID:        "1",
				Status:        "CREATED",
				ProductId:     &[]string{"1"}[0],
				ParentOrderID: &[]int{1}[0],
			},
			{
				ID:            3,
				UserID:        "1",
				Status:        "CREATED",
				ProductId:     &[]string{"2"}[0],
				ParentOrderID: &[]int{1}[0],
			},
		}, nil)

		orders, err := resolvers.GetOrdersByUserID(ctx, orderService, "1")
		a.NoError(err)

		a.Len(orders, 1)
		a.Equal("1", orders[0].ID)
		a.Equal("1", orders[0].UserID)
		a.Equal(model.OrderStatusCreated, orders[0].Status)
		a.Len(orders[0].Products, 2)

	})
}
