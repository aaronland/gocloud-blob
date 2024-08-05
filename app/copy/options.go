package copy

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
	"gocloud.dev/blob"
)

type RunOptions struct {
	SourceBucketURI string
	SourcePath      string
	TargetBucketURI string
	TargetPath      string
	WriterOptions   *blob.WriterOptions
	ACL             string
}

func RunOptionsFromFlagSet(fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	opts := &RunOptions{
		SourceBucketURI: source_bucket_uri,
		SourcePath:      source_path,
		TargetBucketURI: target_bucket_uri,
		TargetPath:      target_path,
		WriterOptions:   new(blob.WriterOptions),
		ACL:             str_acl,
	}

	return opts, nil
}
