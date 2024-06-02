package rmdir

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var bucket_uri string
var path string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("rmdir")
	fs.StringVar(&bucket_uri, "bucket-uri", "", "...")
	fs.StringVar(&path, "path", ".", "...")

	return fs
}
