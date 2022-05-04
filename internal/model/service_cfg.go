package model

import (
	"context"
)

type ServiceConfigKey struct {
	ID        uint64 `json:"id" gorm:"column:id"`
	ServiceID string `json:"serviceID" gorm:"column:service_id"`
	Key       string `json:"key" gorm:"column:key"`
}

func (*ServiceConfigKey) TableName() string {
	return "service_config_key_tab"
}

type ServiceConfigKeyDao interface {
	Put(ctx context.Context, ID string, key string) error
	GetAllKeys(ctx context.Context, ID string) ([]string, error)
}

type DefaultServiceConfigKeyDao struct {
}

func (d *DefaultServiceConfigKeyDao) Put(ctx context.Context, serviceID string, key string) error {
	return GetDatabase(ctx).Create(&ServiceConfigKey{
		ServiceID: serviceID,
		Key:       key,
	}).Error
}

func (d *DefaultServiceConfigKeyDao) GetAllKeys(ctx context.Context, serviceID string) ([]string, error) {
	var configKeys []*ServiceConfigKey
	if err := GetDatabase(ctx).Model(&ServiceConfigKey{}).Where(map[string]interface{}{
		"service_id": serviceID,
	}).Find(&configKeys).Error; err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(configKeys))
	for _, cfgKey := range configKeys {
		keys = append(keys, cfgKey.Key)
	}

	return keys, nil
}
