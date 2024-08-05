package main

import (
	"context"
	"log"

	"github.com/aaronland/gocloud-blob/app/copyuri"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

func main() {

	ctx := context.Background()
	err := copyuri.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
