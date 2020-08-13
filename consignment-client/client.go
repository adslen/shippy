package main

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/adslen/shippy/internal/client"
	"github.com/adslen/shippy/internal/log"
	pb "github.com/adslen/shippy/proto/consignment"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
)

const (
	address         = "localhost:50051"
	defaultFilename = "/home/adslen/workspace/go/shippy/data/consignment.json"
)

func parseFile(file string) (*pb.CreateConsignmentRequest, error) {
	var consignment pb.CreateConsignmentRequest
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = jsonpb.UnmarshalString(string(data), &consignment); err != nil {
		return nil, err
	}

	return &consignment, err
}

func main() {
	conn := client.MakeClientConn()
	shippyClient := pb.NewShippingServiceClient(conn)

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	cr, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := shippyClient.CreateConsignment(context.Background(), cr)
	if err != nil {
		log.L().Fatalf("Could not greet: %v", err)
	}

	log.L().Infof("Created: %v", r.GetStatus())
}

func UnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.L().Debugf("before invoker. method: %+v, request:%+v", method, req)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.L().Debugf("after invoker. reply: %+v", reply)
	return err
}
