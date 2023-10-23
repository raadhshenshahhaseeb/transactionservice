package main

import (
	"log"

	"github.com/hyperversalblocks/txservice/cmd/bootstrapper"
)

func main() {
	if err := bootstrapper.New(); err != nil {
		log.Fatal(err)
	}
}
