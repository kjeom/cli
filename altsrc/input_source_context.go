package altsrc

import (
	"time"

	"github.com/kjeom/cli/v2"
)

// InputSourceContext is an interface used to allow
// other input sources to be implemented as needed.
//
// Source returns an identifier for the input source. In case of file source
// it should return path to the file.
type InputSourceContext interface {
	Source() string

	Int(name string) (int, error)
	Uint64(name string) (uint64, error)
	Duration(name string) (time.Duration, error)
	Float64(name string) (float64, error)
	String(name string) (string, error)
	StringSlice(name string) ([]string, error)
	IntSlice(name string) ([]int, error)
	Int64Slice(name string) ([]int64, error)
	Generic(name string) (cli.Generic, error)
	Bool(name string) (bool, error)

	isSet(name string) bool
}
