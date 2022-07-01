package main

import (
	"github.com/jonathanch7/go-modern-api/cmd/bootstrap"
	"log"
)

func main() {
	bootstrapInitializer := new(bootstrap.Bootstrap)
	if err := bootstrapInitializer.Run(); err != nil {
		log.Fatal("error on deployment ", err)
	}
}
