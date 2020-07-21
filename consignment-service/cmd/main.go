package main

import (
	"context"

	pb "github.com/adslen/shippy/proto/consignment"

	"github.com/adslen/shippy/internal/log"
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

func (s *service) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	// 接收承运的货物
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}
	resp := &pb.Response{Created: true, Consignment: consignment}
	return resp, nil
}
func (s *service) GetConsignments(ctx context.Context, req *pb.GetRequest) (*pb.Response, error) {
	allConsignments := s.repo.GetAll()
	resp := &pb.Response{Consignments: allConsignments}
	return resp, nil
}

func NewShippyService() *service {
	return &service{
		repo: Repository{},
	}
}

func main() {

	//New
	shipService := NewShippyService()
	srv := server.NewServer()
	//Register
	if err := srv.RegisterService(pb.RegisterShippingServiceServer, shipService); err != nil {
		log.L().Fatalf("register service failed. err: %v", err)
	}
	//Run
	log.L().Infof("consignment-service start...")
	srv.Run(server.Addr("10.44.202.200:50051"))
}
