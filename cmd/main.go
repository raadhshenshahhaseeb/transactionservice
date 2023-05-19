package cmd

import (
	"log"

	"github.com/hyperversalblocks/txservice/cmd/server"
)

func main() {
	if err := server.New(); err != nil {
		log.Fatal(err)
	}
}
