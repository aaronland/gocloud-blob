package copy

import (
	"context"
	"flag"
	"fmt"
	"io"

	"github.com/aaronland/gocloud-blob/bucket"
	"github.com/aaronland/gocloud-blob/s3"
	aws_s3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

func Run(ctx context.Context) error {

	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	opts, err := RunOptionsFromFlagSet(fs)

	if err != nil {
		return err
	}

	return RunWithOptions(ctx, opts)
}

func RunWithOptions(ctx context.Context, opts *RunOptions) error {

	source_bucket, err := bucket.OpenBucket(ctx, opts.SourceBucketURI)

	if err != nil {
		return fmt.Errorf("Failed to open source bucket, %w", err)
	}

	defer source_bucket.Close()

	target_bucket, err := bucket.OpenBucket(ctx, opts.TargetBucketURI)

	if err != nil {
		return fmt.Errorf("Failed to open target bucket, %w", err)
	}

	defer target_bucket.Close()

	source_r, err := source_bucket.NewReader(ctx, opts.SourcePath, nil)

	if err != nil {
		return fmt.Errorf("Failed to create new reader, %w", err)
	}

	defer source_r.Close()

	if opts.ACL != "" {

		acl, err := s3.StringACLToObjectCannedACL(opts.ACL)

		if err != nil {
			return fmt.Errorf("Failed to derive ACL object, %w", err)
		}

		before := func(asFunc func(interface{}) bool) error {

			req := &aws_s3.PutObjectInput{}
			ok := asFunc(&req)

			if ok {
				req.ACL = acl
			}

			return nil
		}

		opts.WriterOptions.BeforeWrite = before
	}

	target_wr, err := target_bucket.NewWriter(ctx, opts.TargetPath, opts.WriterOptions)

	if err != nil {
		return fmt.Errorf("Failed to create new writer, %w", err)
	}

	_, err = io.Copy(target_wr, source_r)

	if err != nil {
		return fmt.Errorf("Failed to copy file, %w", err)
	}

	err = target_wr.Close()

	if err != nil {
		return fmt.Errorf("Failed to close writer, %w", err)
	}

	return nil
}
