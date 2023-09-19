package order_test

import (
	"context"
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/internal/db"
	"github.com/bludot/tempmee/order/internal/db/repositories/order"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestOrderRepository_Create(t *testing.T) {
	t.Run("should create order", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)

			order, err := orderRepo.Create(ctx, &order.Order{
				UserID:   "user_id",
				Quantity: 1,
				Price:    1000,
			})

			a.NoError(err)
			a.NotNil(order)
			a.Equal("user_id", order.UserID)
			a.Nil(order.ProductId)
			a.Equal(int64(1), order.Quantity)
			a.Equal(int64(1000), order.Price)

			return nil
		})
	})

}

func TestOrderRepository_FindById(t *testing.T) {
	t.Run("should find order by id", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()
			orderRepo := order.NewOrderRepository(txdb)

			order, err := orderRepo.Create(ctx, &order.Order{
				UserID:   "user_id",
				Quantity: 1,
				Price:    1000,
			})

			a.NoError(err)
			a.NotNil(order)

			order, err = orderRepo.FindById(ctx, order.ID)

			a.NoError(err)
			a.NotNil(order)
			a.Equal("user_id", order.UserID)
			a.Nil(order.ProductId)
			a.Equal(int64(1), order.Quantity)
			a.Equal(int64(1000), order.Price)

			return nil
		})
	})

	t.Run("should return nil if no order with id", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)

			order, err := orderRepo.FindById(ctx, 1)

			a.Error(err)
			a.Nil(order)

			return nil
		})
	})

}

func TestOrderRepository_FindAllByUserId(t *testing.T) {
	t.Run("should find all orders by user id", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}

			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)

			order, err := orderRepo.Create(ctx, &order.Order{
				UserID:   "user_id",
				Quantity: 1,
				Price:    1000,
			})

			a.NoError(err)
			a.NotNil(order)

			orders, err := orderRepo.FindAllByUserId(ctx, "user_id")

			a.NoError(err)
			a.NotNil(orders)
			a.Equal(1, len(*orders))

			return nil
		})
	})

	t.Run("should return nil if no order with user id", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)

			orders, err := orderRepo.FindAllByUserId(ctx, "user_id")

			a.NoError(err)
			a.NotNil(orders)
			a.Equal(0, len(*orders))

			return nil
		})
	})
}

func TestOrderRepository_FindAllByParentOrderId(t *testing.T) {
	t.Run("should find all orders by parent order id", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}

			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)
			parentId := 1
			order, err := orderRepo.Create(ctx, &order.Order{
				UserID:        "user_id",
				ParentOrderID: &parentId,
				Quantity:      1,
				Price:         1000,
			})

			a.NoError(err)
			a.NotNil(order)

			orders, err := orderRepo.FindAllByParentOrderId(ctx, parentId)

			a.NoError(err)
			a.NotNil(orders)
			a.Equal(1, len(*orders))

			return nil
		})
	})

	t.Run("should return nil if no order with parent order id", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}
			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)

			orders, err := orderRepo.FindAllByParentOrderId(ctx, 1)

			a.NoError(err)
			a.NotNil(orders)
			a.Equal(0, len(*orders))

			return nil
		})
	})
}

func TestOrderRepository_Update(t *testing.T) {
	t.Run("should update order", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}

			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)

			createdOrder, err := orderRepo.Create(ctx, &order.Order{
				UserID:   "user_id",
				Quantity: 1,
				Price:    1000,
			})

			a.NoError(err)
			a.NotNil(createdOrder)

			updatedOrder, err := orderRepo.Update(ctx, &order.Order{
				ID:       createdOrder.ID,
				UserID:   "user_id",
				Quantity: 2,
				Price:    2000,
			})

			a.NoError(err)
			a.NotNil(updatedOrder)
			a.Equal(int64(2), updatedOrder.Quantity)
			a.Equal(int64(2000), updatedOrder.Price)

			a.NotEqual(createdOrder.UpdatedAt, updatedOrder.UpdatedAt)

			return nil
		})
	})

	t.Run("should return error if no order with id", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}

			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)

			updatedOrder, err := orderRepo.Update(ctx, &order.Order{
				ID:       1,
				UserID:   "user_id",
				Quantity: 2,
				Price:    2000,
			})

			a.Error(err)
			a.Nil(updatedOrder)

			return nil
		})
	})
}

func TestOrderRepository_Delete(t *testing.T) {
	t.Run("should delete order", func(t *testing.T) {
		a := assert.New(t)

		cfg := config.LoadConfigOrPanic()

		ctx := context.Background()

		database := db.NewDatabase(cfg.DBConfig)
		database.DB.Transaction(func(tx *gorm.DB) error {
			txdb := &db.DB{
				DB: tx,
			}

			defer tx.Rollback()

			orderRepo := order.NewOrderRepository(txdb)

			createdOrder, err := orderRepo.Create(ctx, &order.Order{
				UserID:   "user_id",
				Quantity: 1,
				Price:    1000,
			})

			a.NoError(err)
			a.NotNil(createdOrder)

			err = orderRepo.Delete(ctx, createdOrder.ID)

			a.NoError(err)

			orders, err := orderRepo.FindAll(ctx)

			a.NoError(err)
			a.NotNil(orders)
			a.Equal(0, len(*orders))

			return nil
		})
	})

}
