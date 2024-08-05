package copy

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var source_bucket_uri string
var source_path string

var target_bucket_uri string
var target_path string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("copy")

	fs.StringVar(&source_bucket_uri, "source-uri", "", "...")
	fs.StringVar(&source_path, "source-path", "", "...")

	fs.StringVar(&target_bucket_uri, "target-uri", "", "...")
	fs.StringVar(&target_path, "target-path", "", "...")

	return fs
}
