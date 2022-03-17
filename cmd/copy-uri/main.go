// copy-uri is a command-line tool for cloning the body of a URL (on the web) to a file
// in a blob.Bucket endpoint.
package main

import (
	"context"
	"flag"
	"github.com/aaronland/gocloud-blob/copy"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	"log"
)

func main() {

	uri := flag.String("uri", "", "...")
	fname := flag.String("filename", "", "...")
	bucket_uri := flag.String("bucket-uri", "", "...")
	show_progress := flag.Bool("show-progress", false, "...")

	flag.Parse()

	ctx := context.Background()

	bucket, err := blob.OpenBucket(ctx, *bucket_uri)

	if err != nil {
		log.Fatalf("Failed to open bucket, %v", err)
	}

	opts := &copy.CopyURLOptions{
		Bucket:       bucket,
		Filename:     *fname,
		ShowProgress: *show_progress,
	}

	err = copy.CopyURLStringToBucket(ctx, opts, *uri)

	if err != nil {
		log.Fatalf("Failed to copy URL to bucket, %v", err)
	}
}
