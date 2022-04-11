package model

type Service struct {
	Name string `gorm:"column:name"`
	Key  string `gorm:"column:key"`
}

type ServiceDao interface {
	Put(service *Service) error
	Get(serviceName string) (*Service, error)
}

type ServiceDaoMysqlImpl struct {
}

func (s *ServiceDaoMysqlImpl) Put(service *Service) error {
}

func (s *ServiceDaoMysqlImpl) Get(serviceName string) (*Service, error) {
	//TODO implement me
	panic("implement me")
}
