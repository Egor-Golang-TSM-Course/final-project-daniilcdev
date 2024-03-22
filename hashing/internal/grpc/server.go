package grpc

import (
	"net"

	"hash-system/hashing/internal/grpc/pb"

	"google.golang.org/grpc"
)

func Serve(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	hss := &hashingServerService{}
	grpcServer := grpc.NewServer()
	pb.RegisterHashServiceServer(grpcServer, hss)

	return grpcServer.Serve(listener)
}
