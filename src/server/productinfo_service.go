package main

import (
	"context"
	"grpc_soldshop/src/proto/pb"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedProductInfoServer
	products map[string]*pb.Product
}

func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductID, error) {
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.products == nil {
		s.products = make(map[string]*pb.Product)
	}
	s.products[in.Id] = in
	return &pb.ProductID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetProduct(ctx context.Context, in *pb.ProductID) (*pb.Product, error) {
	value, existed := s.products[in.Value]
	if !existed {
		return nil, status.Errorf(codes.NotFound, "Product does not existed.")
	}
	return value, status.New(codes.OK, "").Err()
}
