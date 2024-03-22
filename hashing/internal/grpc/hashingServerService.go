package grpc

import (
	"context"
	"encoding/base64"
	"errors"
	"hash-system/hashing/internal/grpc/pb"
	"sync"
)

type hashingServerService struct {
	pb.UnimplementedHashServiceServer
}

var cache map[string]string = make(map[string]string, 64)
var hashes map[string]struct{} = make(map[string]struct{}, 64)
var mu *sync.RWMutex = &sync.RWMutex{}

func (hss *hashingServerService) GetHash(ctx context.Context, payload *pb.HashingPayload) (*pb.Hash, error) {
	mu.RLock()
	defer mu.Unlock()

	hash, ok := cache[payload.Src]

	if !ok {
		return nil, errors.New("hash not yet created")
	}

	return &pb.Hash{Value: hash}, nil
}

func (hss *hashingServerService) CheckHash(ctx context.Context, payload *pb.HashedPayload) (*pb.Exists, error) {
	mu.RLock()
	defer mu.Unlock()

	_, ok := hashes[payload.Hash]

	return &pb.Exists{Value: ok}, nil
}

func (hss *hashingServerService) CreateHash(ctx context.Context, payload *pb.HashingPayload) (*pb.Hash, error) {
	hash := base64.StdEncoding.EncodeToString([]byte(payload.Src))

	mu.Lock()
	defer mu.Unlock()

	cache[payload.Src] = hash
	hashes[hash] = struct{}{}

	return &pb.Hash{Value: hash}, nil
}
