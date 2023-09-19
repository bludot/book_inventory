package resolvers

import (
	"github.com/bludot/tempmee/inventory/config"
	"github.com/bludot/tempmee/inventory/graph/model"
)

func APIInfo(cfg config.Config) (*model.APIInfo, error) {
	return &model.APIInfo{
		Name: cfg.AppConfig.APPName,
		InventoryAPI: &model.InventoryAPI{
			Version: cfg.AppConfig.Version,
		},
	}, nil
}

func InventoryAPI(cfg config.Config) (*model.InventoryAPI, error) {
	return &model.InventoryAPI{
		Version: cfg.AppConfig.Version,
	}, nil
}
