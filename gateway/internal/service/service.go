package service

import (
	"fmt"

	"github.com/go-kratos/etcd/registry"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

var EtcdCli *clientv3.Client

func init() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	EtcdCli = cli
	if err != nil {
		panic(err)
	}
	fmt.Println("xxxxxxxxxxxxx")
}
