package main

import (
	grpcHashing "hashing/internal/grpc"
	"log"
	adapters "shared/adapters"
)

func main() {
	if err := grpcHashing.Serve("hashing:9000", adapters.NewInMemoryStorage()); err != nil {
		log.Fatalln(err)
	}
}
