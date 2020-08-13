package registry

type IRegistry interface {
	//Registry(srvInfo *ServerInfo) error
	Registry(srvName string, addr string, meta interface{}) error
}

func NewRegistry(opts ...Option) IRegistry {

	return newEtcdRegistry(opts...)
}
