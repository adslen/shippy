package server

type IServer interface {
	RegisterService(register interface{}, handler ServiceInstance, Options ...RegisterOption) error
	Run(...Option) error
}
