// copy-uri is a command-line tool for cloning the body of a URL (on the web) to a file
// in a blob.Bucket endpoint.
package main

// ./bin/copy-uri -source-uri https://static.sfomuseum.org/media/189/736/720/3/1897367203_lt3HuJ5ALbY4SDoxTd4oi7abOF7gQZKM_c.jpg -target-uri cwd:// -filename test.jpg

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	_ "github.com/aaronland/gocloud-blob-s3"
	"github.com/aaronland/gocloud-blob/bucket"
	"github.com/aaronland/gocloud-blob/copy"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sfomuseum/go-flags/flagset"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
)

func main() {

	var source_uri string
	var filename string
	var target_uri string
	var show_progress bool
	var acl string
	var part_size int64
	var mode string

	fs := flagset.NewFlagSet("copy")

	fs.StringVar(&source_uri, "source-uri", "", "The URI of the file to copy.")
	fs.StringVar(&target_uri, "target-uri", "", "A valid gocloud.dev/blob.Bucket URI.")

	fs.StringVar(&filename, "filename", "", "The final filename of the file to copy. If empty the basename of the `-source-uri` flag value will be used.")

	fs.StringVar(&acl, "acl", "", "An optional AWS S3 ACL to assign to the file being copied.")
	fs.Int64Var(&part_size, "part-size", 0, "The buffer size (in bytes) to use when buffering data into chunks and sending them as parts to S3. If 0 the default value for the `aws/aws-sdk-go/service/s3/s3manager` package will be used.")

	fs.BoolVar(&show_progress, "show-progress", false, "Show copy progress.")
	fs.StringVar(&mode, "mode", "cli", "Valid options are: cli, lambda.")

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "BLOB")

	if err != nil {
		log.Fatalf("Failed to set flags from environment, %v", err)
	}

	ctx := context.Background()

	bucket, err := bucket.OpenBucket(ctx, target_uri)

	if err != nil {
		log.Fatalf("Failed to open bucket '%s', %v", target_uri, err)
	}

	defer bucket.Close()

	opts := &copy.CopyURLOptions{
		Bucket:       bucket,
		Filename:     filename,
		ShowProgress: show_progress,
	}

	if acl != "" || part_size != 0 {

		before := func(asFunc func(interface{}) bool) error {

			if acl != "" {

				input := &s3manager.UploadInput{}
				ok := asFunc(&input)

				if !ok {
					return fmt.Errorf("Not an S3 type")
				}

				input.ACL = aws.String(acl)
			}

			if part_size != 0 {

				uploader := &s3manager.Uploader{}
				ok := asFunc(&uploader)

				if !ok {
					return fmt.Errorf("Not an S3 type")
				}

				uploader.PartSize = part_size
			}

			return nil
		}

		opts.WriterOptions = &blob.WriterOptions{
			BeforeWrite: before,
		}
	}

	run := func(ctx context.Context, opts *copy.CopyURLOptions, source string) error {

		if opts.Filename == "" {
			opts.Filename = filepath.Base(source)
		}

		return copy.CopyURLStringToBucket(ctx, opts, source)
	}

	switch mode {
	case "cli":

		err := run(ctx, opts, source_uri)

		if err != nil {
			log.Fatalf("Failed to copy URL to bucket, %v", err)
		}

	case "lambda":

		handler := func(ctx context.Context, source string) error {
			return run(ctx, opts, source)
		}

		lambda.Start(handler)

	default:

		log.Fatalf("Invalid mode")
	}
}
