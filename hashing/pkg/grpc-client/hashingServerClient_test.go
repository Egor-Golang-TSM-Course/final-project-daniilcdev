package grpcClient

import (
	"context"
	adapters "hash-system/hashing/cmd/adapters"
	grpcServer "hash-system/hashing/internal/grpc"
	client "hash-system/hashing/pkg"
	"log"

	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	const addr = ":9001"
	defer bootServer(addr, adapters.NewInMemoryStorage())()

	cl, err := New(addr)
	defer cl.CloseConn()

	assert.NoError(t, err)
	assert.NotNil(t, cl)
}

func TestNewClientFailed(t *testing.T) {
	const addr = ":9001"
	defer bootServer(addr, adapters.NewInMemoryStorage())()

	cl, err := New("")

	assert.Error(t, err)
	assert.Nil(t, cl)
}

func TestCreateHash(t *testing.T) {
	const addr = ":9001"
	tearDown := bootServer(addr, adapters.NewInMemoryStorage())
	defer tearDown()

	cl, _ := New(addr)
	defer cl.CloseConn()

	hash, err := cl.CreateHash(context.Background(), "some-payload")

	assert.NoError(t, err)
	assert.NotNil(t, hash)
}

func TestCreateHashForEmptyPayload(t *testing.T) {
	const addr = ":9001"
	tearDown := bootServer(addr, adapters.NewInMemoryStorage())
	defer tearDown()

	cl, _ := New(addr)
	defer cl.CloseConn()

	hash, err := cl.CreateHash(context.Background(), "")

	assert.Error(t, err)
	assert.Empty(t, hash)

	tearDown()
}

func TestCheckHash(t *testing.T) {
	const addr = ":9001"
	storage := adapters.NewInMemoryStorage()
	tearDown := bootServer(addr, storage)
	defer tearDown()

	cl, _ := New(addr)
	defer cl.CloseConn()

	hash, _ := cl.CreateHash(context.Background(), "some-payload")
	exists, err := cl.CheckHash(context.Background(), hash)

	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestCheckHashThatDoesNotCreated(t *testing.T) {
	const addr = ":9001"
	storage := adapters.NewInMemoryStorage()
	tearDown := bootServer(addr, storage)
	defer tearDown()

	cl, _ := New(addr)
	defer cl.CloseConn()

	exists, err := cl.CheckHash(context.Background(), "some-hash")

	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestGetHash(t *testing.T) {
	const addr = ":9001"
	storage := adapters.NewInMemoryStorage()
	tearDown := bootServer(addr, storage)
	defer tearDown()

	cl, _ := New(addr)
	defer cl.CloseConn()
	created, _ := cl.CreateHash(context.Background(), "some-payload")
	hash, err := cl.GetHash(context.Background(), "some-payload")

	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, created, hash)
}

func TestGetHashReturnsErrorWhenNoHash(t *testing.T) {
	const addr = ":9001"
	tearDown := bootServer(addr, adapters.NewInMemoryStorage())
	defer tearDown()

	cl, _ := New(addr)
	defer cl.CloseConn()

	hash, err := cl.GetHash(context.Background(), "some-payload")

	assert.Error(t, err)
	assert.Empty(t, hash)
}

func bootServer(addr string, storage client.HashStorage) func() {
	go func(addr string) {
		err := grpcServer.Serve(addr, storage)
		if err != nil {
			log.Fatalln(err)
		}
	}(addr)

	<-time.After(10 * time.Millisecond)

	return func() {
		grpcServer.Stop()
	}
}
