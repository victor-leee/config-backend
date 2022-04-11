package main

import (
	config_backend "github.com/victor-leee/config-backend/github.com/victor-leee/config-backend"
	"github.com/victor-leee/config-backend/internal/handler"
	"github.com/victor-leee/config-backend/internal/service"
)

func main() {
	config_backend.RegisterConfigBackendService(&handler.ConfigBackendServiceImpl{
		ConfigOperator: &service.ETCDConfigOperator{},
	})
	server := config_backend.GetServer()
	server.Start()
	server.WaitTermination()
}
