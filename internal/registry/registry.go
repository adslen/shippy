package registry

import (
	"context"
	"time"

	"github.com/adslen/shippy/internal/log"
	"github.com/coreos/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"google.golang.org/grpc/naming"
)

type etcdRegistry struct {
	cli *clientv3.Client

	opts Options
}

func newEtcdRegistry(options ...Option) *etcdRegistry {
	opts := Options{}
	for _, o := range options {
		o(&opts)
	}

	if opts.addrs == nil {
		opts.addrs = []string{"127.0.0.1:2379"}
	}

	config := clientv3.Config{
		Endpoints:   opts.addrs,
		DialTimeout: opts.timeout,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		log.L().Fatal(err)
	}

	if opts.timeout == 0 {
		opts.timeout = 1 * time.Second
	}

	return &etcdRegistry{
		cli:  cli,
		opts: opts,
	}
}

func (e *etcdRegistry) Registry(srvName string, addr string, meta interface{}) error {
	log.Debug("servername: ", srvName)
	ctx, cancel := context.WithTimeout(context.Background(), e.opts.timeout)
	defer cancel()

	r := &etcdnaming.GRPCResolver{Client: e.cli}
	return r.Update(ctx, srvName, naming.Update{Op: naming.Add, Addr: addr, Metadata: ""})
}
