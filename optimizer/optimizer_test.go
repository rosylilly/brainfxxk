package optimizer_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/rosylilly/brainfxxk/optimizer"
	"github.com/rosylilly/brainfxxk/parser"
)

func TestOptimizer(t *testing.T) {
	testCases := []struct {
		source   string
		expected string
	}{
		{
			source:   "++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.",
			expected: "Hello World!\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.source, func(t *testing.T) {
			p, err := parser.Parse(strings.NewReader(tc.source))
			if err != nil {
				t.Fatal(err)
			}
			o := optimizer.NewOptimizer()
			prog, _ := o.Optimize(p)
			fmt.Printf("program Length: [%v]", len(p.Expressions))
			fmt.Printf("optimize program Length:[%v], \nprogram:%#v", len(prog.Expressions), prog)
		})
	}
}
