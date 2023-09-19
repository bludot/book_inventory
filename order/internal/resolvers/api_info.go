package resolvers

import (
	"github.com/bludot/tempmee/order/config"
	"github.com/bludot/tempmee/order/graph/model"
)

func APIInfo(cfg config.Config) (*model.APIInfo, error) {
	return &model.APIInfo{
		Name: cfg.AppConfig.APPName,
		OrderAPI: &model.OrderAPI{
			Version: cfg.AppConfig.Version,
		},
	}, nil
}

func OrderAPI(cfg config.Config) (*model.OrderAPI, error) {
	return &model.OrderAPI{
		Version: cfg.AppConfig.Version,
	}, nil
}
