package cmd

import (
	"github.com/hyperversalblocks/txservice/cmd/server"
	"log"
)

func main() {
	if err := server.New(); err != nil {
		log.Fatal(err)
	}
}
