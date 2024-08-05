package main

import (
	"context"
	"log"

	"github.com/aaronland/gocloud-blob/app/read"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

func main() {

	ctx := context.Background()
	err := read.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
