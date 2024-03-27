package main

import (
	internal "gateway/internal"
	"log"
)

func main() {
	if err := internal.StartApiServer(); err != nil {
		log.Fatalln(err)
	}
}
