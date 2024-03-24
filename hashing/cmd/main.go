package main

import (
	"hash-system/hashing/internal/grpc"
	"log"
)

func main() {
	if err := grpc.Serve(":9000", NewInMemoryStorage()); err != nil {
		log.Fatalln(err)
	}
}
