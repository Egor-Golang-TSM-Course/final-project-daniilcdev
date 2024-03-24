package main

import (
	"hash-system/hashing/cmd/adapters"
	"hash-system/hashing/internal/grpc"
	"log"
)

func main() {
	if err := grpc.Serve(":9000", adapters.NewInMemoryStorage()); err != nil {
		log.Fatalln(err)
	}
}
