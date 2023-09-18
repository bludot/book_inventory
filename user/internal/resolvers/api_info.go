package resolvers

import (
	"github.com/bludot/tempmee/user/config"
	"github.com/bludot/tempmee/user/graph/model"
)

func APIInfo(cfg config.Config) (*model.APIInfo, error) {
	return &model.APIInfo{
		Name: cfg.AppConfig.APPName,
		UserAPI: &model.UserAPI{
			Version: cfg.AppConfig.Version,
		},
	}, nil
}

func AnimeAPI(cfg config.Config) (*model.UserAPI, error) {
	return &model.UserAPI{
		Version: cfg.AppConfig.Version,
	}, nil
}
