package registry

import (
	"time"
)

type Option func(*Options)

type Options struct {
	dialTimeout time.Duration
	timeout     time.Duration
	addrs       []string

	password string
	username string
}

func Address(addrs []string) Option {
	return func(opts *Options) {
		opts.addrs = addrs
	}
}

func Timeout(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.timeout = timeout
	}
}
