package main

import (
	grpcHashing "hashing/internal/grpc"
	"log"
	adapters "shared/adapters"
	"shared/cfg"
)

func main() {
	vars := cfg.Load()
	if err := grpcHashing.Serve(vars.HashingServiceAddr, adapters.NewInMemoryStorage()); err != nil {
		log.Fatalln(err)
	}
}
