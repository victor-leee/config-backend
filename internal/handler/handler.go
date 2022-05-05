package handler

import (
	"context"
	"errors"
	"github.com/victor-leee/config-backend/github.com/victor-leee/config-backend"
	"github.com/victor-leee/config-backend/internal/service"
)

type ConfigBackendServiceImpl struct {
	ConfigOperator service.ConfigOperator
}

func (c *ConfigBackendServiceImpl) GetAllKeys(ctx context.Context, req *config_backend.GetAllKeysRequest) (*config_backend.GetAllKeysResponse, error) {
	keys, err := c.ConfigOperator.Keys(ctx, req.ServiceId, req.ServiceKey)
	if err == nil {
		err = errors.New("")
	}
	return &config_backend.GetAllKeysResponse{
		BaseResponse: &config_backend.BaseResponse{
			ErrCode: config_backend.ErrorCode_SUCCESS,
			ErrMsg:  err.Error(),
		},
		Keys: keys,
	}, nil
}

func (c *ConfigBackendServiceImpl) GetConfig(ctx context.Context, req *config_backend.GetConfigRequest) (*config_backend.GetConfigResponse, error) {
	value, err := c.ConfigOperator.Get(ctx, req.ServiceId, req.ServiceKey, req.Key)
	resp := &config_backend.GetConfigResponse{
		BaseResponse: &config_backend.BaseResponse{
			ErrCode: config_backend.ErrorCode_SUCCESS,
		},
		KeyExist: value != nil,
	}
	if resp.KeyExist {
		resp.Value = *value
	}
	if err != nil {
		resp.BaseResponse.ErrMsg = err.Error()
	}

	return resp, nil
}

func (c *ConfigBackendServiceImpl) PutConfig(ctx context.Context, req *config_backend.PutConfigRequest) (*config_backend.PutConfigResponse, error) {
	errCode := config_backend.ErrorCode_SUCCESS
	errMsg := ""
	if err := c.ConfigOperator.Put(ctx, req.ServiceId, req.ServiceKey, req.Key, req.Value); err != nil {
		errCode = config_backend.ErrorCode_ERR_INTERNAL_SERVER_ERROR
		errMsg = err.Error()
	}

	return &config_backend.PutConfigResponse{
		BaseResponse: &config_backend.BaseResponse{
			ErrCode: errCode,
			ErrMsg:  errMsg,
		},
	}, nil
}
