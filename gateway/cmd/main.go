package main

import (
	"hash-system/gateway/internal"
	"log"
)

func main() {
	if err := internal.StartApiServer(); err != nil {
		log.Fatalln(err)
	}
}
