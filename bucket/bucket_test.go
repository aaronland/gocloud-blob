package bucket

import (
	"context"
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	_ "gocloud.dev/blob/s3blob"
	"testing"
)

func TestOpenBucket(t *testing.T) {

	uris := []string{
		"file:///tmp",
		"mem://",
		"s3://example?region=us-east-1",
	}

	for _, uri := range uris {

		ctx := context.Background()

		_, err := OpenBucket(ctx, uri)

		if err != nil {
			t.Fatalf("Failed to create bucket for '%s', %v", uri, err)
		}
	}
}
