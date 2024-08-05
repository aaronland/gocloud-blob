package copy

import (
	"flag"

	"gocloud.dev/blob"
	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	SourceBucketURI string
	SourcePath string
	TargetBucketURI string
	TargetPath string
	WriterOptions *blob.WriterOptions
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)
	
	opts := &RunOptions{
		SourceBucketURI: source_bucket_uri,
		SourcePath: source_path,
		TargetBucketURI: target_bucket_uri,
		TargetPath: target_path,
	}

	return opts, nil
}
