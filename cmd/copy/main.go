package main

import (
	"context"
	"log"
	
	"github.com/aaronland/gocloud-blob/app/copy"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

func main() {

	ctx := context.Background()
	err := copy.Run(ctx)

	if err != nil {
		log.Fatal(err)
	}
}
