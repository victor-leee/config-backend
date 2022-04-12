package model

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"strings"
)

type Service struct {
	Name string `gorm:"column:name"`
	Key  string `gorm:"column:service_key"`
}

func (*Service) TableName() string {
	return "service_tab"
}

var (
	ErrDuplicateService = errors.New("duplicate service name")
	ErrServiceNotFound  = errors.New("cannot find the service")
)

type ServiceDao interface {
	Put(ctx context.Context, service *Service) error
	Get(ctx context.Context, serviceName string) (*Service, error)
}

type ServiceDaoMysqlImpl struct {
}

func (s *ServiceDaoMysqlImpl) Put(ctx context.Context, service *Service) error {
	if service == nil {
		return errors.New("service cannot be nil")
	}
	if len(strings.TrimSpace(service.Name)) == 0 {
		return errors.New("service cannot be empty")
	}
	if err := s.create(ctx, service); err == nil || !errors.Is(err, ErrDuplicateService) {
		return err
	}

	return s.update(ctx, service)
}

func (s *ServiceDaoMysqlImpl) create(ctx context.Context, service *Service) error {
	if err := GetDatabase(ctx).Create(service).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr); mysqlErr != nil {
			if mysqlErr.Number == 1062 {
				return ErrDuplicateService
			}
		}

		return err
	}

	return nil
}

func (s *ServiceDaoMysqlImpl) update(ctx context.Context, service *Service) error {
	whereMap := map[string]interface{}{
		"name": service.Name,
	}
	updatesMap := map[string]interface{}{
		"key": service.Key,
	}
	tx := GetDatabase(ctx).Model(&Service{}).Where(whereMap).Updates(updatesMap)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return ErrServiceNotFound
	}

	return nil
}

func (s *ServiceDaoMysqlImpl) Get(ctx context.Context, serviceName string) (*Service, error) {
	whereMap := map[string]interface{}{
		"name": serviceName,
	}
	var svc *Service
	if err := GetDatabase(ctx).Model(&Service{}).First(&svc, whereMap).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrServiceNotFound
		}
		return nil, err
	}

	return svc, nil
}
