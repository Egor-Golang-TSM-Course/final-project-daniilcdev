package main

import "log"

func main() {
	if err := startGateway(); err != nil {
		log.Fatalln(err)
	}
}
