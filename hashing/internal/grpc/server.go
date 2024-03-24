package grpc

import (
	"net"

	"hash-system/hashing/internal/grpc/pb"
	client "hash-system/hashing/pkg"

	"google.golang.org/grpc"
)

func Serve(addr string, storage client.HashStorage) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	hss := &hashingServerService{storage: storage}
	grpcServer := grpc.NewServer()
	pb.RegisterHashServiceServer(grpcServer, hss)

	return grpcServer.Serve(listener)
}
