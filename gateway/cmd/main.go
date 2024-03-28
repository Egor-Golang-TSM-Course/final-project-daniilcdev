package main

import (
	internal "gateway/internal"
	"log"
	"shared/cfg"
)

func main() {
	vars := cfg.Load()
	if err := internal.StartApiServer(vars); err != nil {
		log.Fatalln(err)
	}
}
