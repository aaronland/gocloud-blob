package bucket

import (
	"context"
	"testing"
	"fmt"
	
	_ "gocloud.dev/blob/fileblob"
	_ "gocloud.dev/blob/memblob"
	_ "github.com/aaronland/gocloud-blob/s3"
)

func TestOpenBucket(t *testing.T) {

	uris := []string{
		"file:///tmp",
		"mem://",
		// "s3://example?region=us-east-1",
		"cwd://",
	}

	for _, uri := range uris {

		ctx := context.Background()

		b, err := OpenBucket(ctx, uri)

		if err != nil {
			t.Fatalf("Failed to create bucket for '%s', %v", uri, err)
		}

		defer b.Close()
		
		fmt.Printf("Opened %s\n", uri)
	}
}
