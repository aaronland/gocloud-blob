package main

import (
	"context"
	"log"

	"github.com/aaronland/gocloud-blob/app/copy"
	_ "github.com/aaronland/gocloud-blob/s3"	
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	_ "gocloud.dev/blob/s3blob"
)

func main() {

	ctx := context.Background()
	err := copy.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
