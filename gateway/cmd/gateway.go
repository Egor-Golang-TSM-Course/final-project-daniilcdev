package main

import "hash-system/gateway/internal"

func startGateway() error {
	return internal.StartApiServer()
}
