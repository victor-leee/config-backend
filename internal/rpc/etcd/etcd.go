package etcd

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var client *clientv3.Client

func MustInitETCDClient() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"config-etcd-infra:80"},
		DialTimeout: time.Second,
	})
	if err != nil {
		panic(err)
	}

	client = cli
}

func GetClient() *clientv3.Client {
	return client
}
