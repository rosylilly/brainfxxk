package interpreter

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/rosylilly/brainfxxk/ast"
	"github.com/rosylilly/brainfxxk/parser"
)

var (
	ErrInputFinished = fmt.Errorf("input finished")
)

type Interpreter struct {
	Program *ast.Program
	Memory  []byte
	Pointer int
}

func Run(ctx context.Context, s io.Reader, w io.Writer, r io.Reader) error {
	p, err := parser.Parse(s)
	if err != nil {
		return err
	}

	return NewInterpreter(p).Run(ctx, w, r)
}

func NewInterpreter(p *ast.Program) *Interpreter {
	return &Interpreter{
		Program: p,
		Memory:  make([]byte, 128),
		Pointer: 0,
	}
}

func (i *Interpreter) Run(ctx context.Context, w io.Writer, r io.Reader) error {
	err := i.runExpressions(ctx, i.Program.Expressions, w, r)

	if errors.Is(err, ErrInputFinished) {
		return nil
	}
	return err

}

func (i *Interpreter) runExpressions(ctx context.Context, exprs []ast.Expression, w io.Writer, r io.Reader) error {
	for _, expr := range exprs {
		if err := i.runExpression(ctx, expr, w, r); err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) runExpression(ctx context.Context, expr ast.Expression, w io.Writer, r io.Reader) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	switch e := expr.(type) {
	case *ast.PointerIncrementExpression:
		i.Pointer += 1
	case *ast.PointerDecrementExpression:
		i.Pointer -= 1
	case *ast.ValueIncrementExpression:
		i.Memory[i.Pointer] += 1
	case *ast.ValueDecrementExpression:
		i.Memory[i.Pointer] -= 1
	case *ast.OutputExpression:
		if _, err := w.Write([]byte{i.Memory[i.Pointer]}); err != nil {
			return err
		}
	case *ast.InputExpression:
		b := make([]byte, 1)
		if _, err := r.Read(b); err != nil {
			if errors.Is(err, io.EOF) {
				return ErrInputFinished
			}
			return err
		}
		i.Memory[i.Pointer] = b[0]
	case *ast.WhileExpression:
		for i.Memory[i.Pointer] != 0 {
			if err := i.runExpressions(ctx, e.Body, w, r); err != nil {
				return err
			}
		}
	}
	return nil
}
