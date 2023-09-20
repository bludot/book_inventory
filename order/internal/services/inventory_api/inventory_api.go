package inventory_api

import (
	"context"
	"github.com/Khan/genqlient/graphql"
	"github.com/bludot/tempmee/order/config"
)

type InventoryApiImpl interface {
	GetBookByID(ctx context.Context, id string) (*GetBookByIDResponse, error)
}

type InventoryApi struct {
	Client graphql.Client
}

func NewInventoryApi(cfg config.InventoryAPIConfig) InventoryApiImpl {
	return &InventoryApi{
		Client: graphql.NewClient("http://"+cfg.Host+":"+string(cfg.Port), nil),
	}
}

func (i *InventoryApi) GetBookByID(ctx context.Context, id string) (*GetBookByIDResponse, error) {
	return GetBookByID(ctx, i.Client, id)
}
