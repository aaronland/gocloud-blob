package bucket

import (
	"context"
	"fmt"
	gc_blob "gocloud.dev/blob"
	"net/url"
)

// OpenBucket is a local helper function to open a gocloud.dev/blob Bucket URI and ensuring
// that files will not be written with their corresponding metdata (`.attrs`) files.
func OpenBucket(ctx context.Context, bucket_uri string) (*gc_blob.Bucket, error) {

	u, err := url.Parse(bucket_uri)

	if err != nil {
		return nil, fmt.Errorf("Failed to parse bucket URI, %w", err)
	}

	q := u.Query()

	if q.Get("metadata") != "skip" {
		q.Set("metadata", "skip")
		u.RawQuery = q.Encode()
		bucket_uri = u.String()
	}

	return gc_blob.OpenBucket(ctx, bucket_uri)
}
