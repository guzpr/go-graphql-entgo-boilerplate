package db

import (
	"fmt"
	"log"
	"os"

	"github.com/sekalahita/epirus/internal/ent/gen"
)

func CreateEntClient() *gen.Client {
	var options []gen.Option
	if os.Getenv("DEBUG_ENTCLIENT") == "1" {
		options = []gen.Option{
			gen.Debug(),
			gen.Log(func(nit ...interface{}) {
				fmt.Println(nit)
			}),
		}
	}

	client, err := gen.Open(os.Getenv("DATABASE_TYPE"), os.Getenv("DATABASE_SOURCE"), options...)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}

	return client
}
