package server

type Option func(*Options)

type RegisterOption func(*RegisterOptions)

type Options struct {
	Addrs string
}

type RegisterOptions struct {
}

func Addr(addr string) Option {
	return func(opts *Options) {
		opts.Addrs = addr
	}
}
