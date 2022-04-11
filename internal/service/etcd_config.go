package service

import "context"

type ConfigOperator interface {
	Get(ctx context.Context, serviceID, serviceKey, key string) (string, error)
	Put(ctx context.Context, serviceID, serviceKey, key, value string) error
}

type ETCDConfigOperator struct {
}

func (E *ETCDConfigOperator) Get(ctx context.Context, serviceID, serviceKey, key string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (E *ETCDConfigOperator) Put(ctx context.Context, serviceID, serviceKey, key, value string) error {
	//TODO implement me
	panic("implement me")
}
