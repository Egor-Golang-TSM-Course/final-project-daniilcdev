package grpcHashing

import (
	"context"
	"encoding/base64"
	"errors"
	pb "shared/grpc/pb"
	shared "shared/pkg"
)

type hashingServerService struct {
	pb.UnimplementedHashServiceServer
	storage shared.HashStorage
}

func (hss *hashingServerService) GetHash(ctx context.Context, payload *pb.HashingPayload) (*pb.Hash, error) {
	hash, err := hss.storage.Get(payload.Src)

	if err != nil {
		return nil, err
	}

	return &pb.Hash{Value: hash}, nil
}

func (hss *hashingServerService) CheckHash(ctx context.Context, payload *pb.HashedPayload) (*pb.Exists, error) {
	contains := hss.storage.Contains(payload.Hash)
	return &pb.Exists{Value: contains}, nil
}

func (hss *hashingServerService) CreateHash(ctx context.Context, payload *pb.HashingPayload) (*pb.Hash, error) {
	if payload.Src == "" {
		return nil, errors.New("can't create hash for empty string")
	}

	hash := base64.StdEncoding.EncodeToString([]byte(payload.Src))

	hss.storage.Add(payload.Src, hash)

	return &pb.Hash{Value: hash}, nil
}
