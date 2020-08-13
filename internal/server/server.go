package server

import (
	"errors"
	"net"
	"reflect"

	"github.com/adslen/shippy/internal/log"
	"github.com/adslen/shippy/internal/registry"
	"github.com/adslen/shippy/internal/service"
	"google.golang.org/grpc"
)

type ServiceInstance service.Instance

type server struct {
	grpcSrv *grpc.Server

	register registry.IRegistry

	serviceInfo *serverInfo
}

type serverInfo struct {
	serviceName string
}

//func (srv *server) GetServiceInfo(){
//}

func (srv *server) RegisterService(register interface{}, handler ServiceInstance, Options ...RegisterOption) error {
	if register == nil || handler == nil {
		return errors.New("register nil ")
	}

	rt := reflect.TypeOf(register)

	if rt.Kind() != reflect.Func || rt.NumIn() != 2 || rt.In(0) != reflect.TypeOf(srv.grpcSrv) {
		return errors.New("")
	}

	if !reflect.TypeOf(handler).Implements(rt.In(1)) {
		return errors.New("")
	}
	reflect.ValueOf(register).Call([]reflect.Value{reflect.ValueOf(srv.grpcSrv), reflect.ValueOf(handler)})
	return nil
}

func (srv *server) Run(options ...Option) error {
	opt := Options{}
	for _, o := range options {
		o(&opt)
	}

	lis, err := net.Listen("tcp", opt.Addrs)
	if err != nil {
		return err
	}

	si := srv.grpcSrv.GetServiceInfo()
	for name := range si {
		srv.serviceInfo = &serverInfo{
			serviceName: name,
		}
	}

	log.L().Infof("grpc server addr %s", opt.Addrs)

	if err := srv.register.Registry(srv.serviceInfo.serviceName, opt.Addrs, nil); err != nil {
		return err
	}
	return srv.grpcSrv.Serve(lis)
}

func NewServer() IServer {
	return &server{
		grpcSrv: grpc.NewServer(),

		register: registry.NewRegistry(registry.Address([]string{"127.0.0.1:2379"})),
	}
}
