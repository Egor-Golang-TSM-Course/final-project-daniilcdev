package grpcClient

import (
	"context"
	"hash-system/hashing/internal/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HashServiceClient interface {
	GetHash(ctx context.Context, in string) (string, error)
	CheckHash(ctx context.Context, hash string) (bool, error)
	CreateHash(ctx context.Context, in string) (string, error)

	CloseConn()
}

type hashServiceClient struct {
	rpc pb.HashServiceClient
	cc  *grpc.ClientConn
}

func New(addr string) (HashServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewHashServiceClient(conn)

	return &hashServiceClient{rpc: grpcClient, cc: conn}, err
}

func (hsc *hashServiceClient) GetHash(ctx context.Context, in string) (string, error) {
	if hash, err := hsc.rpc.GetHash(ctx, &pb.HashingPayload{Src: in}); err != nil {
		return "", err
	} else {
		return hash.Value, nil
	}
}

func (hsc *hashServiceClient) CheckHash(ctx context.Context, hash string) (bool, error) {
	if exist, err := hsc.rpc.CheckHash(ctx, &pb.HashedPayload{Hash: hash}); err != nil {
		return false, err
	} else {
		return exist.Value, nil
	}
}

func (hsc *hashServiceClient) CreateHash(ctx context.Context, in string) (string, error) {
	if hash, err := hsc.rpc.CreateHash(ctx, &pb.HashingPayload{Src: in}); err != nil {
		return "", err
	} else {
		return hash.Value, nil
	}
}

func (hsc *hashServiceClient) CloseConn() {
	hsc.cc.Close()
}
