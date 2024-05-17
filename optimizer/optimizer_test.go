package optimizer_test

import (
	"strings"
	"testing"

	"github.com/rosylilly/brainfxxk/optimizer"
	"github.com/rosylilly/brainfxxk/parser"
)

func TestOptimizer(t *testing.T) {
	source := "++>>>++"

	p, err := parser.Parse(strings.NewReader(source))
	if err != nil {
		t.Fatal(err)
	}

	o := optimizer.NewOptimizer()
	prog, err := o.Optimize(p)

	t.Logf("%#v", prog)
}
