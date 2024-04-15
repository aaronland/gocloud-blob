package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	_ "github.com/aaronland/gocloud-blob-s3"
	"github.com/aaronland/gocloud-blob/bucket"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

func main() {

	source_uri := flag.String("source-uri", "", "A valid gocloud.dev/blob.Bucket URI.")
	target_uri := flag.String("target-uri", "", "A valid gocloud.dev/blob.Bucket URI.")
	
	source_path := flag.String("source-path", "", "The path (relative to `-source-uri`) for the file to copy.")
	target_path := flag.String("target-path", "", "The path (relative to `-target-uri`) where the target file should be copied.")

	acl := flag.String("acl", "", "An optional AWS S3 ACL to assign to the file being copied.")
	part_size := flag.Int64("part-size", 0, "The buffer size (in bytes) to use when buffering data into chunks and sending them as parts to S3. If 0 the default value for the `aws/aws-sdk-go/service/s3/s3manager` package will be used.")

	flag.Parse()

	ctx := context.Background()

	source_bucket, err := bucket.OpenBucket(ctx, *source_uri)

	if err != nil {
		log.Fatalf("Failed to open source bucket '%s', %v", *source_uri, err)
	}

	defer source_bucket.Close()

	target_bucket, err := bucket.OpenBucket(ctx, *target_uri)

	if err != nil {
		log.Fatalf("Failed to open target bucket '%s', %v", *target_uri, err)
	}

	defer target_bucket.Close()

	var wr_opts *blob.WriterOptions

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

		wr_opts.BeforeWrite = before
	}

	source_r, err := source_bucket.NewReader(ctx, *source_path, nil)

	if err != nil {
		log.Fatalf("Failed to create source reader, %v", err)
	}

	defer source_r.Close()

	target_wr, err := target_bucket.NewWriter(ctx, *target_path, wr_opts)

	if err != nil {
		log.Fatalf("Failed to create target writer, %v", err)
	}

	_, err = io.Copy(target_wr, source_r)

	if err != nil {
		log.Fatalf("Failed to copy source file to target, %v", err)
	}

	err = target_wr.Close()

	if err != nil {
		log.Fatalf("Failed to close target writer, %v", err)
	}
}
