package remove

import (
	"context"
	_ "fmt"
	"io"
	"log/slog"
	"strings"

	"gocloud.dev/blob"
)

// RemoveTree will remove 'uri' and all its contents from bucket 'b'.
func RemoveTree(ctx context.Context, b *blob.Bucket, uri string) error {

	var removeTree func(context.Context, *blob.Bucket, string) error

	removeTree = func(ctx context.Context, b *blob.Bucket, prefix string) error {

		iter := b.List(&blob.ListOptions{
			Delimiter: "/",
			Prefix:    prefix,
		})

		for {
			obj, err := iter.Next(ctx)

			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			if obj.IsDir {

				err = removeTree(ctx, b, obj.Key)

				if err != nil {
					return err
				}

			}

			// trailing slashes confuse Go Cloud...

			path := strings.TrimRight(obj.Key, "/")
			slog.Info("Delete key", "path", path)
			continue

			err = b.Delete(ctx, path)

			if err != nil {
				slog.Error("Failed to delete key", "path", path, "error", err)
			}
		}

		return nil
	}

	return removeTree(ctx, b, uri)
}
