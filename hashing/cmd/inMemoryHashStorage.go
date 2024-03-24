package main

import (
	"errors"
	client "hash-system/hashing/pkg"
	"sync"
)

type InMemoryStorage struct {
	cache  map[string]string
	hashes map[string]struct{}
	mu     *sync.RWMutex
}

func NewInMemoryStorage() client.HashStorage {
	return &InMemoryStorage{
		cache:  make(map[string]string, 64),
		hashes: make(map[string]struct{}, 64),
		mu:     &sync.RWMutex{},
	}
}

func (ims *InMemoryStorage) Contains(hash string) bool {
	ims.mu.RLock()
	defer ims.mu.Unlock()

	_, ok := ims.hashes[hash]
	return ok
}

func (ims *InMemoryStorage) Get(payload string) (string, error) {
	ims.mu.RLock()
	defer ims.mu.Unlock()
	hash, ok := ims.cache[payload]

	if !ok {
		return "", errors.New("hash not created")
	}

	return hash, nil
}

func (ims *InMemoryStorage) Add(payload, hash string) {
	ims.mu.Lock()
	defer ims.mu.Unlock()

	ims.cache[payload] = hash
	ims.hashes[hash] = struct{}{}
}
