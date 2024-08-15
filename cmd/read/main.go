package main

import (
	"context"
	"log"

	"github.com/aaronland/gocloud-blob/app/read"
	_ "github.com/aaronland/gocloud-blob/s3"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	_ "gocloud.dev/blob/s3blob"
)

func main() {

	ctx := context.Background()
	err := read.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
