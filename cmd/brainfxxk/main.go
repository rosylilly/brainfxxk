package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/rosylilly/brainfxxk/interpreter"
)

var (
	config = &interpreter.Config{
		Writer:               os.Stdout,
		Reader:               os.Stdin,
		MemorySize:           30000,
		RaiseErrorOnOverflow: false,
		RaiseErrorOnEOF:      false,
	}
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options...] [file]:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.IntVar(&config.MemorySize, "memory-size", config.MemorySize, "memory size")
	flag.BoolVar(&config.RaiseErrorOnOverflow, "raise-error-on-overflow", config.RaiseErrorOnOverflow, "raise error on overflow")
	flag.BoolVar(&config.RaiseErrorOnEOF, "raise-error-on-eof", config.RaiseErrorOnEOF, "raise error on eof")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	flag.Parse()

	var source io.ReadCloser = os.Stdin
	if flag.NArg() > 0 {
		fp, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}

		source = fp
	}
	defer source.Close()

	if err := interpreter.Run(ctx, source, config); err != nil {
		log.Fatal(err)
	}
}
