package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"

	"github.com/rosylilly/brainfxxk/interpreter"
)

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

	if err := interpreter.Run(ctx, source, os.Stdout, os.Stdin); err != nil {
		log.Fatal(err)
	}
}
