package rmdir

import (
	"flag"
)

type RunOptions struct {
	BucketURI string
	Path      string
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	opts := &RunOptions{
		BucketURI: bucket_uri,
		Path:      path,
	}

	return opts, nil
}
