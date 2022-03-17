// copy-uri is a command-line tool for cloning the body of a URL (on the web) to a file
// in a blob.Bucket endpoint.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/aaronland/gocloud-blob/copy"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	_ "github.com/aaronland/gocloud-blob-s3"
	"log"
)

func main() {

	uri := flag.String("uri", "", "...")
	fname := flag.String("filename", "", "...")
	bucket_uri := flag.String("bucket-uri", "", "...")
	show_progress := flag.Bool("show-progress", false, "...")
	acl := flag.String("acl", "", "...")

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

	if *acl != "" {

		before := func(asFunc func(interface{}) bool) error {

			req := &s3manager.UploadInput{}
			ok := asFunc(&req)

			if !ok {
				return fmt.Errorf("Not an S3 type")
			}

			req.ACL = aws.String(*acl)
			return nil
		}

		opts.WriterOptions = &blob.WriterOptions{
			BeforeWrite: before,
		}
	}

	err = copy.CopyURLStringToBucket(ctx, opts, *uri)

	if err != nil {
		log.Fatalf("Failed to copy URL to bucket, %v", err)
	}
}
