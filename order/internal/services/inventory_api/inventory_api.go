package inventory_api

import (
	"context"
	"encoding/json"
	"github.com/bludot/tempmee/order/config"
	"net/http"
	"strconv"
)

type InventoryApiImpl interface {
	GetBookByID(ctx context.Context, id string) (map[string]interface{}, error)
}

type InventoryApi struct {
	Host   string
	Client *http.Client
}

func NewInventoryApi(cfg config.InventoryAPIConfig) InventoryApiImpl {

	return &InventoryApi{
		Client: &http.Client{},
		Host:   "http://" + cfg.Host + ":" + strconv.Itoa(cfg.Port) + cfg.Path,
	}
}

func (i *InventoryApi) GetBookByID(ctx context.Context, id string) (map[string]interface{}, error) {
	resp, err := i.Client.Get(i.Host + "/" + id)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var book map[string]interface{}

	err = json.NewDecoder(resp.Body).Decode(&book)
	if err != nil {
		return nil, err
	}

	return book, nil
}
