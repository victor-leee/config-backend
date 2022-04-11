package handler

import (
	"context"
	"github.com/victor-leee/config-backend/github.com/victor-leee/config-backend"
	"github.com/victor-leee/config-backend/internal/service"
	"google.golang.org/protobuf/proto"
)

type ConfigBackendServiceImpl struct {
	ConfigOperator service.ConfigOperator
}

func (c *ConfigBackendServiceImpl) GetConfig(ctx context.Context, req *config_backend.GetConfigRequest) (*config_backend.GetConfigResponse, error) {
	value, err := c.ConfigOperator.Get(ctx, req.ServiceId, req.ServiceKey, req.Key)
	if value == nil {
		value = proto.String("")
	}
	resp := &config_backend.GetConfigResponse{
		BaseResponse: &config_backend.BaseResponse{
			ErrCode: config_backend.ErrorCode_SUCCESS,
			ErrMsg:  err.Error(),
		},
		Value: *value,
	}

	return resp, nil
}

func (c *ConfigBackendServiceImpl) PutConfig(ctx context.Context, req *config_backend.PutConfigRequest) (*config_backend.PutConfigResponse, error) {
	//TODO implement me
	panic("implement me")
}
