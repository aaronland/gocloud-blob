package main

import (
	"context"
	"log/slog"
	"os"

	_ "github.com/aaronland/gocloud-blob-s3"
	"github.com/aaronland/gocloud-blob/app/rmdir"
	_ "gocloud.dev/blob/fileblob"
)

func main() {

	ctx := context.Background()
	logger := slog.Default()

	err := rmdir.Run(ctx, logger)

	if err != nil {
		logger.Error("Failed to run rmdir application", "error", err)
		os.Exit(1)
	}

}
