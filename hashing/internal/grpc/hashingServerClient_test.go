package grpcHashing

import (
	"context"
	"log"
	shared "shared/pkg"
	sharedadapters "shared/pkg/adapters"

	"time"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	const addr = ":9001"
	defer bootServer(addr, sharedadapters.NewInMemoryStorage())()

	cl, err := shared.New(addr)
	defer cl.CloseConn()

	assert.NoError(t, err)
	assert.NotNil(t, cl)
}

func TestNewClientFailed(t *testing.T) {
	const addr = ":9001"
	defer bootServer(addr, sharedadapters.NewInMemoryStorage())()

	cl, err := shared.New("")

	assert.Error(t, err)
	assert.Nil(t, cl)
}

func TestClient_CreateHash(t *testing.T) {
	const addr = ":9001"
	tearDown := bootServer(addr, sharedadapters.NewInMemoryStorage())
	defer tearDown()

	cl, _ := shared.New(addr)
	defer cl.CloseConn()

	hash, err := cl.CreateHash(context.Background(), "some-payload")

	assert.NoError(t, err)
	assert.NotNil(t, hash)
}

func TestClient_CreateHashForEmptyPayload(t *testing.T) {
	const addr = ":9001"
	tearDown := bootServer(addr, sharedadapters.NewInMemoryStorage())
	defer tearDown()

	cl, _ := shared.New(addr)
	defer cl.CloseConn()

	hash, err := cl.CreateHash(context.Background(), "")

	assert.Error(t, err)
	assert.Empty(t, hash)

	tearDown()
}

func TestClient_CheckHash(t *testing.T) {
	const addr = ":9001"
	storage := sharedadapters.NewInMemoryStorage()
	tearDown := bootServer(addr, storage)
	defer tearDown()

	cl, _ := shared.New(addr)
	defer cl.CloseConn()

	hash, _ := cl.CreateHash(context.Background(), "some-payload")
	exists, err := cl.CheckHash(context.Background(), hash)

	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestCheckHashThatDoesNotCreated(t *testing.T) {
	const addr = ":9001"
	storage := sharedadapters.NewInMemoryStorage()
	tearDown := bootServer(addr, storage)
	defer tearDown()

	cl, _ := shared.New(addr)
	defer cl.CloseConn()

	exists, err := cl.CheckHash(context.Background(), "some-hash")

	assert.NoError(t, err)
	assert.False(t, exists)
}

func TestClient_GetHash(t *testing.T) {
	const addr = ":9001"
	storage := sharedadapters.NewInMemoryStorage()
	tearDown := bootServer(addr, storage)
	defer tearDown()

	cl, _ := shared.New(addr)
	defer cl.CloseConn()
	created, _ := cl.CreateHash(context.Background(), "some-payload")
	hash, err := cl.GetHash(context.Background(), "some-payload")

	assert.NoError(t, err)
	assert.NotNil(t, hash)
	assert.Equal(t, created, hash)
}

func TestClient_GetHashReturnsErrorWhenNoHash(t *testing.T) {
	const addr = ":9001"
	tearDown := bootServer(addr, sharedadapters.NewInMemoryStorage())
	defer tearDown()

	cl, _ := shared.New(addr)
	defer cl.CloseConn()

	hash, err := cl.GetHash(context.Background(), "some-payload")

	assert.Error(t, err)
	assert.Empty(t, hash)
}

func bootServer(addr string, storage shared.HashStorage) func() {
	go func(addr string) {
		err := Serve(addr, storage)
		if err != nil {
			log.Fatalln(err)
		}
	}(addr)

	<-time.After(10 * time.Millisecond)

	return func() {
		grpcServer.Stop()
	}
}
