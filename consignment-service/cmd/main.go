package main

import (
	"context"

	pb "github.com/adslen/shippy/proto/consignment"
	"google.golang.org/grpc/codes"
	status_v1 "google.golang.org/grpc/status"

	"github.com/adslen/shippy/internal/log"
	"github.com/adslen/shippy/internal/server"
	"google.golang.org/genproto/googleapis/rpc/status"
)

const (
	PORT = "10.44.202.200:50051"
)

type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
}

//
// 我们存放多批货物的仓库，实现了 IRepository 接口
//
type Repository struct {
	consignments []*pb.Consignment
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

type service struct {
	repo Repository
}

// 托运一批货物
func (s *service) CreateConsignment(ctx context.Context, req *pb.CreateConsignmentRequest) (*pb.CreateConsignmentResponse, error) {
	size := len(req.GetConsignments())
	res := &pb.CreateConsignmentResponse{
		Consignments: make([]*pb.Consignment, size),
		Status:       make([]*status.Status, size),
	}

	for i, rc := range req.GetConsignments() {
		consignment, err := s.repo.Create(rc)
		res.Consignments[i] = consignment
		res.Status[i] = &status.Status{
			Code: int32(codes.OK),
		}
		if err != nil {
			res.Status[i] = &status.Status{
				Code:    int32(status_v1.Code(err)),
				Message: err.Error(),
			}
		}
	}
	return res, nil
}

//获取货物信息
func (s *service) GetConsignment(ctx context.Context, req *pb.GetConsignmentRequest) (*pb.GetConsignmentResponse, error) {

	return nil, nil
}

//展现托运货物
func (s *service) ListConsignments(ctx context.Context, req *pb.ListConsignmentRequest) (*pb.ListConsignmentResponse, error) {

	return nil, nil
}

func NewShippyService() *service {
	return &service{
		repo: Repository{},
	}
}

func main() {
	log.L().Infof("consignment-service start...")
	//New
	shipService := NewShippyService()
	srv := server.NewServer()
	if err := srv.RegisterService(pb.RegisterShippingServiceServer, shipService); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(server.Addr("0.0.0.0:50051")); err != nil {
		log.Fatal(err)
	}
}
