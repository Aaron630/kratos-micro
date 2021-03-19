package service

import (
	"context"
	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	v1 "gateway/api/helloworld/v1"
	"gateway/internal/biz"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	g "google.golang.org/grpc"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper("service/greeter", logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.Infof("SayHello Received: %v", in.GetName())
	if in.GetName() == "error" {
		return nil, errors.NotFound("ErrorReason", in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *GreeterService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginReply, error) {
	//s.log.Infof("SayHello Received: %v", in.Username)
	//
	//return &v1.LoginReply{Message: "Hello " + in.Username}, nil
	r := registry.New(EtcdCli)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///user"),
		grpc.WithOptions(g.WithBalancerName("round_robin")),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		//log.Fatal(err)
	}
	client := v1.NewGreeterClient(conn)
	reply, err := client.Login(context.Background(), in)
	if err != nil {
		//log.Fatal(err)
	}
	return reply, nil
}

func (s *GreeterService) GetOrderForm(ctx context.Context, in *v1.GetOrderFormRequest) (*v1.GetOrderFormReply, error) {
	//s.log.Infof("SayHello Received: %v", in.GetId())
	//
	//return &v1.GetOrderFormReply{Id: in.GetId()}, nil
	r := registry.New(EtcdCli)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///orderForm"),
		grpc.WithOptions(g.WithBalancerName("round_robin")),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		//log.Fatal(err)
	}
	client := v1.NewGreeterClient(conn)
	reply, err := client.GetOrderForm(context.Background(), in)
	if err != nil {
		//log.Fatal(err)
	}
	return reply, nil
}
