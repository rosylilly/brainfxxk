package interpreter

import (
	"io"
)

type Config struct {
	Writer               io.Writer
	Reader               io.Reader
	MemorySize           int
	RaiseErrorOnOverflow bool
	RaiseErrorOnEOF      bool
}
