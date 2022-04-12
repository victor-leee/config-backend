package main

import (
	"github.com/sirupsen/logrus"
	config_backend "github.com/victor-leee/config-backend/github.com/victor-leee/config-backend"
	"github.com/victor-leee/config-backend/internal/config"
	"github.com/victor-leee/config-backend/internal/handler"
	"github.com/victor-leee/config-backend/internal/model"
	"github.com/victor-leee/config-backend/internal/rpc/etcd"
	"github.com/victor-leee/config-backend/internal/service"
	"time"
)

func main() {
	time.Sleep(3 * time.Second)
	cfg, err := config.Init()
	if err != nil {
		logrus.Fatal(err)
	}
	if err = model.Init(cfg); err != nil {
		logrus.Fatal(err)
	}

	etcd.MustInitETCDClient()
	config_backend.RegisterConfigBackendService(&handler.ConfigBackendServiceImpl{
		ConfigOperator: &service.ETCDConfigOperator{
			ServiceDao: &model.ServiceDaoMysqlImpl{},
			ETCDClient: etcd.GetClient(),
		},
	})
	server := config_backend.GetServer()
	server.Start()
	server.WaitTermination()
}
