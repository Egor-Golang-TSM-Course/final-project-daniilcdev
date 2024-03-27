package grpcHashing

import (
	"context"
	adapters "shared/adapters"
	pb "shared/grpc/pb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHash(t *testing.T) {
	s := createService()
	hash, err := s.CreateHash(context.Background(), &pb.HashingPayload{
		Src: "payload",
	})

	assert.NoError(t, err)
	assert.NotNil(t, hash)
}

func TestCreateHashForEmptyPayload(t *testing.T) {
	s := createService()
	hash, err := s.CreateHash(context.Background(), &pb.HashingPayload{
		Src: "",
	})

	assert.Error(t, err)
	assert.Nil(t, hash)
}

func TestCheckHash(t *testing.T) {
	s := createService()
	s.storage.Add("some-payload", "some-hash")

	exists, err := s.CheckHash(context.Background(), &pb.HashedPayload{Hash: "some-hash"})

	assert.NoError(t, err)
	assert.True(t, exists.Value)
}

func TestGetHash(t *testing.T) {
	s := createService()
	s.storage.Add("some-payload", "some-hash")

	hash, err := s.GetHash(context.Background(), &pb.HashingPayload{Src: "some-payload"})

	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, "some-hash", hash.Value)
}

func TestGetHashReturnsErrorWhenNoHash(t *testing.T) {
	s := createService()

	hash, err := s.GetHash(context.Background(), &pb.HashingPayload{Src: "some-payload"})

	assert.Error(t, err)
	assert.Nil(t, hash)
}

func createService() *hashingServerService {
	return &hashingServerService{
		storage: adapters.NewInMemoryStorage(),
	}
}
