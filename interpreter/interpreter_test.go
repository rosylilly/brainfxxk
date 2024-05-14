package interpreter_test

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/rosylilly/brainfxxk/interpreter"
)

func TestInterpreter(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	testCases := []struct {
		source   string
		input    string
		expected string
	}{
		{
			source:   "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.",
			input:    "",
			expected: "Hello World!\n",
		},
		{
			source: `
>++++[<++++++++>-]>++++++++[<++++++>-]<++.<.>+.<.>++.<.>++.<.>------..<.>
.++.<.>--.++++++.<.>------.>+++[<+++>-]<-.<.>-------.+.<.> -.+++++++.<.>
------.--.<.>++.++++.<.>---.---.<.> +++.-.<.>+.+++.<.>--.--.<.> ++.++++.<.>
---.-----.<.>+++++.+.<.>.------.<.> ++++++.----.<.> ++++.++.<.> -.-----.<.>
+++++.+.<.>.--.`,
			input:    "",
			expected: "2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97",
		},
		{
			source:   "+[>,.<]",
			input:    "Hello",
			expected: "Hello",
		},
		{
			source: `
++++++[->++++>>+>+>-<<<<<]>[<++++>>+++>++++>>+++>+++++>+++++>>>>>>++>>++<
<<<<<<<<<<<<<-]<++++>+++>-->+++>->>--->++>>>+++++[->++>++<<]<<<<<<<<<<[->
-[>>>>>>>]>[<+++>.>.>>>>..>>>+<]<<<<<-[>>>>]>[<+++++>.>.>..>>>+<]>>>>+<-[
<<<]<[[-<<+>>]>>>+>+<<<<<<[->>+>+>-<<<<]<]>>[[-]<]>[>>>[>.<<.<<<]<[.<<<<]
>]>.<<<<<<<<<<<]`,
			input:    "",
			expected: makeFizzBuzz(100),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.source, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			w := &bytes.Buffer{}
			if err := interpreter.Run(ctx, strings.NewReader(tc.source), w, r); err != nil {
				t.Fatal(err)
			}

			if w.String() != tc.expected {
				t.Errorf("got: %v, expected: %v", w.String(), tc.expected)
			}
		})
	}
}

type infinityReader struct{}

func (ir *infinityReader) Read(p []byte) (n int, err error) {
	p[0] = 'I'
	return 1, nil
}

func TestInterpreterInfinityRead(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	source := "+[>,.<]"
	ir := &infinityReader{}
	w := &bytes.Buffer{}

	err := interpreter.Run(ctx, strings.NewReader(source), w, ir)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("got: %v, expected: %v", err, context.DeadlineExceeded)
	}
}

func makeFizzBuzz(count int) string {
	var b strings.Builder
	for i := 1; i <= count; i++ {
		b.WriteString(fizzBuzz(i))
		b.WriteString("\n")
	}
	return b.String()
}

func fizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return fmt.Sprint(n)
	}
}
