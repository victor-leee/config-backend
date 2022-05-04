package service

import (
	"context"
	"errors"
	"github.com/victor-leee/config-backend/internal/model"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/proto"
	"strings"
)

type ConfigOperator interface {
	Get(ctx context.Context, serviceID, serviceKey, key string) (*string, error)
	Put(ctx context.Context, serviceID, serviceKey, key, value string) error
	Keys(ctx context.Context, serviceID, serviceKey string) ([]string, error)
}

type ETCDConfigOperator struct {
	ServiceConfigKeyDao model.ServiceConfigKeyDao
	ServiceDao          model.ServiceDao
	ETCDClient          *clientv3.Client
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
	E.ServiceConfigKeyDao.Put(ctx, serviceID, E.configKey(serviceID, key))
	return err
}

func (E *ETCDConfigOperator) Keys(ctx context.Context, serviceID, serviceKey string) ([]string, error) {
	if err := E.validateService(ctx, serviceID, serviceKey); err != nil {
		return nil, err
	}

	keys, err := E.ServiceConfigKeyDao.GetAllKeys(ctx, serviceID)
	if err != nil {
		return nil, err
	}
	for i, k := range keys {
		keys[i] = E.refineOriginalKey(k)
	}

	return keys, nil
}

func (E *ETCDConfigOperator) validateService(ctx context.Context, serviceID, serviceKey string) error {
	serviceEntity, err := E.ServiceDao.Get(ctx, serviceID)
	if err != nil {
		return err
	}
	if serviceEntity.ServiceKey != serviceKey {
		return ErrMismatchServiceKey
	}

	return nil
}

func (E *ETCDConfigOperator) configKey(keys ...string) string {
	return strings.Join(keys, ".")
}

func (E *ETCDConfigOperator) refineOriginalKey(mixKey string) string {
	keys := strings.Split(mixKey, ".")
	return keys[1]
}
