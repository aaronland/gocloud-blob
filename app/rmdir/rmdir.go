package rmdir

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/aaronland/gocloud-blob/bucket"
	"github.com/aaronland/gocloud-blob/remove"
	"github.com/sfomuseum/go-flags/flagset"
)

func Run(ctx context.Context, logger *slog.Logger) error {
	fs := DefaultFlagSet()

	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) error {

	flagset.Parse(fs)

	opts, err := RunOptionsFromFlagSet(fs)

	if err != nil {
		return fmt.Errorf("Failed to derive run options, %w", err)
	}
	return RunWithOptions(ctx, opts, logger)
}

func RunWithOptions(ctx context.Context, opts *RunOptions, logger *slog.Logger) error {

	b, err := bucket.OpenBucket(ctx, opts.BucketURI)

	if err != nil {
		return fmt.Errorf("Failed to open bucket, %v", err)
	}

	defer b.Close()

	err = remove.RemoveTree(ctx, b, opts.Path)

	if err != nil {
		return fmt.Errorf("Failed to remove tree for %s, %v", opts.Path, err)
	}

	return nil
}
