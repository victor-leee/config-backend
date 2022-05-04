package model

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrServiceNotFound = errors.New("service doesn't exist")
)

type Service struct {
	ServiceID  string `json:"serviceID" gorm:"column:complete_path"`
	ServiceKey string `json:"serviceKey" gorm:"column:serviceKey"`
}

type ServiceDao interface {
	Get(ctx context.Context, serviceID string) (*Service, error)
}

type DefaultServiceDao struct {
}

func (d *DefaultServiceDao) Get(ctx context.Context, serviceID string) (*Service, error) {
	s := &Service{}
	whereMap := map[string]interface{}{
		"unique_complete_path": serviceID,
	}
	if err := GetDatabase(ctx).Model(s).Where(whereMap).First(s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceNotFound
		}
		return nil, err
	}

	return s, nil
}
