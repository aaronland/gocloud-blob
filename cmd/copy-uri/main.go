// copy-uri is a command-line tool for cloning the body of a URL (on the web) to a file
// in a blob.Bucket endpoint.
package main

import (
	"context"
	"flag"
	"fmt"
	_ "github.com/aaronland/gocloud-blob-s3"
	"github.com/aaronland/gocloud-blob/copy"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	"log"
)

func main() {

	source_uri := flag.String("source-uri", "", "...")
	fname := flag.String("filename", "", "...")
	target_uri := flag.String("target-uri", "", "...")
	show_progress := flag.Bool("show-progress", false, "...")
	acl := flag.String("acl", "", "...")
	part_size := flag.Int64("part-size", 0, "...")

	flag.Parse()

	ctx := context.Background()

	bucket, err := blob.OpenBucket(ctx, *target_uri)

	if err != nil {
		log.Fatalf("Failed to open bucket '%s', %v", *target_uri, err)
	}

	opts := &copy.CopyURLOptions{
		Bucket:       bucket,
		Filename:     *fname,
		ShowProgress: *show_progress,
	}

	if *acl != "" || *part_size != 0 {

		before := func(asFunc func(interface{}) bool) error {

			if *acl != "" {

				input := &s3manager.UploadInput{}
				ok := asFunc(&input)

				if !ok {
					return fmt.Errorf("Not an S3 type")
				}

				input.ACL = aws.String(*acl)
			}

			if *part_size != 0 {

				uploader := &s3manager.Uploader{}
				ok := asFunc(&uploader)

				if !ok {
					return fmt.Errorf("Not an S3 type")
				}

				uploader.PartSize = *part_size
			}

			return nil
		}

		opts.WriterOptions = &blob.WriterOptions{
			BeforeWrite: before,
		}
	}

	err = copy.CopyURLStringToBucket(ctx, opts, *source_uri)

	if err != nil {
		log.Fatalf("Failed to copy URL to bucket, %v", err)
	}
}
