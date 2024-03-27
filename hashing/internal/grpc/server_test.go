package grpcHashing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type mockStorage struct {
}

func (ms *mockStorage) Add(payload, hash string)           {}
func (ms *mockStorage) Get(payload string) (string, error) { return "", nil }
func (ms *mockStorage) Contains(hash string) bool          { return false }

func TestServe(t *testing.T) {
	var err error
	go func() {
		err = Serve(":9000", &mockStorage{})
	}()

	<-time.After(10 * time.Millisecond)

	assert.NoError(t, err)
}

func TestServeFailed(t *testing.T) {
	var err error
	go func() {
		err = Serve("bad-addr", &mockStorage{})
	}()

	<-time.After(10 * time.Millisecond)

	assert.Error(t, err)
}
