package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const payload = "test-payload"
const hash = "abc-def-am"

func TestAdd(t *testing.T) {
	s := NewInMemoryStorage()
	s.Add(payload, hash)

	assert.True(t, s.Contains(hash))
}

func TestGetReturnsErrorWhenNoHashForKey(t *testing.T) {
	s := NewInMemoryStorage()

	hash, err := s.Get(payload)

	assert.Error(t, err)
	assert.Empty(t, hash)
}

func TestGetReturnsStoredHash(t *testing.T) {
	s := NewInMemoryStorage()

	s.Add(payload, hash)

	stored, err := s.Get(payload)

	assert.NoError(t, err)
	assert.Equal(t, hash, stored)
}
