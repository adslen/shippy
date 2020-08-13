package client

import "google.golang.org/grpc"

type Client struct {
	opts *Options
}

func MakeClientConn() *grpc.ClientConn {
	conn := &grpc.ClientConn{}

	return conn
}
