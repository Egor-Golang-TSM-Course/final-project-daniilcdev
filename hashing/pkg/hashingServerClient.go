package client

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

func New() (HashServiceClient, error) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	conn, err := grpc.Dial(":9000", opts...)

	if err != nil {
		return nil, err
	}

	grpcClient := pb.NewHashServiceClient(conn)

	return &hashServiceClient{rpc: grpcClient}, err
}

type hashServiceClient struct {
	rpc pb.HashServiceClient
	cc  grpc.ClientConn
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
