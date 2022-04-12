package service

import (
	"context"
	"errors"
	"github.com/victor-leee/config-backend/internal/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/proto"
)

type ConfigOperator interface {
	Get(ctx context.Context, serviceID, serviceKey, key string) (*string, error)
	Put(ctx context.Context, serviceID, serviceKey, key, value string) error
}

type ETCDConfigOperator struct {
	ServiceDao model.ServiceDao
	ETCDClient *clientv3.Client
}

var (
	ErrMismatchServiceKey = errors.New("mismatch service key")
	ErrKeyNotExist        = errors.New("key not exist")
)

func (E *ETCDConfigOperator) Get(ctx context.Context, serviceID, serviceKey, key string) (*string, error) {
	if err := E.validateService(ctx, serviceID, serviceKey); err != nil {
		return nil, err
	}

	resp, err := E.ETCDClient.Get(ctx, E.configKey(serviceID, key))
	if err != nil {
		return nil, err
	}
	if resp.Count == 0 {
		return nil, ErrKeyNotExist
	}

	return proto.String(string(resp.Kvs[0].Value)), nil
}

func (E *ETCDConfigOperator) Put(ctx context.Context, serviceID, serviceKey, key, value string) error {
	if err := E.validateService(ctx, serviceID, serviceKey); err != nil {
		return err
	}
	_, err := E.ETCDClient.Put(ctx, E.configKey(serviceID, key), value)
	return err
}

func (E *ETCDConfigOperator) validateService(ctx context.Context, serviceID, serviceKey string) error {
	serviceEntity, err := E.ServiceDao.Get(ctx, serviceID)
	if err != nil {
		return err
	}
	if serviceEntity.Key != serviceKey {
		return ErrMismatchServiceKey
	}

	return nil
}

func (E *ETCDConfigOperator) configKey(serviceID, key string) string {
	return serviceID + "." + key
}
