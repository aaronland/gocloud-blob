package main

import (
	"context"
	"flag"
	"log"

	"github.com/aaronland/gocloud-blob/bucket"
	"github.com/aaronland/gocloud-blob/remove"
)

func main() {

	var bucket_uri string
	var path string

	flag.StringVar(&bucket_uri, "bucket-uri", "", "...")
	flag.StringVar(&path, "path", ".", "...")

	flag.Parse()

	ctx := context.Background()

	b, err := bucket.OpenBucket(ctx, bucket_uri)

	if err != nil {
		log.Fatalf("Failed to open bucket, %v", err)
	}

	defer b.Close()

	err = remove.RemoveTree(ctx, b, path)

	if err != nil {
		log.Fatalf("Failed to remove tree for %s, %v", path, err)
	}

}
