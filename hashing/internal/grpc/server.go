package grpcHashing

import (
	"log"
	"net"

	pb "shared/grpc/pb"
	shared "shared/pkg"

	"google.golang.org/grpc"
)

var grpcServer *grpc.Server

func Serve(addr string, storage shared.HashStorage) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	hss := &hashingServerService{storage: storage}
	grpcServer = grpc.NewServer()
	pb.RegisterHashServiceServer(grpcServer, hss)

	log.Default().Println("starting gRPC...")
	return grpcServer.Serve(listener)
}

func Stop() {
	grpcServer.GracefulStop()
}
