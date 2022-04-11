package handler

import (
	"context"
	"github.com/victor-leee/config-backend/github.com/victor-leee/config-backend"
	"github.com/victor-leee/config-backend/internal/service"
)

type ConfigBackendServiceImpl struct {
	ConfigOperator service.ConfigOperator
}

func (c *ConfigBackendServiceImpl) GetConfig(ctx context.Context, req *config_backend.GetConfigRequest) (*config_backend.GetConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ConfigBackendServiceImpl) PutConfig(ctx context.Context, req *config_backend.PutConfigRequest) (*config_backend.PutConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}
