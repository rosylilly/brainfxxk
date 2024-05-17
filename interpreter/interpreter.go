package interpreter

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/rosylilly/brainfxxk/ast"
	"github.com/rosylilly/brainfxxk/optimizer"
	"github.com/rosylilly/brainfxxk/parser"
)

var (
	ErrInputFinished  = fmt.Errorf("input finished")
	ErrMemoryOverflow = fmt.Errorf("memory overflow")
)

type Interpreter struct {
	Program *ast.Program
	Config  *Config
	Memory  []byte
	Pointer int
}

func Run(ctx context.Context, s io.Reader, c *Config) error {
	p, err := parser.Parse(s)
	if err != nil {
		return err
	}

	return NewInterpreter(p, c).Run(ctx)
}

func NewInterpreter(p *ast.Program, c *Config) *Interpreter {
	return &Interpreter{
		Program: p,
		Config:  c,
		Memory:  make([]byte, c.MemorySize),
		Pointer: 0,
	}
}

func (i *Interpreter) Run(ctx context.Context) error {
	p, err := optimizer.NewOptimizer().Optimize(i.Program)
	if err != nil {
		return err
	}

	err = i.runExpressions(ctx, p.Expressions)
	if errors.Is(err, ErrInputFinished) && !i.Config.RaiseErrorOnEOF {
		return nil
	}
	return err

}

func (i *Interpreter) runExpressions(ctx context.Context, exprs []ast.Expression) error {
	for _, expr := range exprs {
		if err := i.runExpression(ctx, expr); err != nil {
			return err
		}
	}
	return nil
}

func (i *Interpreter) runExpression(ctx context.Context, expr ast.Expression) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	switch e := expr.(type) {
	case *ast.PointerIncrementExpression:
		if i.Pointer == len(i.Memory)-1 && i.Config.RaiseErrorOnOverflow {
			return fmt.Errorf("%w: %d to pointer overflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Pointer += 1
	case *ast.PointerDecrementExpression:
		if i.Pointer == 0 && i.Config.RaiseErrorOnOverflow {
			return fmt.Errorf("%w: %d to pointer underflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Pointer -= 1
	case *ast.ValueIncrementExpression:
		if i.Memory[i.Pointer] == 255 && i.Config.RaiseErrorOnOverflow {
			return fmt.Errorf("%w: %d to memory overflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Memory[i.Pointer] += 1
	case *ast.ValueDecrementExpression:
		if i.Memory[i.Pointer] == 0 && i.Config.RaiseErrorOnOverflow {
			return fmt.Errorf("%w: %d to memory underflow, on %d:%d", ErrMemoryOverflow, i.Pointer, e.StartPos(), e.EndPos())
		}
		i.Memory[i.Pointer] -= 1
	case *ast.OutputExpression:
		if _, err := i.Config.Writer.Write([]byte{i.Memory[i.Pointer]}); err != nil {
			return err
		}
	case *ast.InputExpression:
		b := make([]byte, 1)
		if _, err := i.Config.Reader.Read(b); err != nil {
			if errors.Is(err, io.EOF) {
				return ErrInputFinished
			}
			return err
		}
		i.Memory[i.Pointer] = b[0]
	case *ast.WhileExpression:
		for i.Memory[i.Pointer] != 0 {
			if err := i.runExpressions(ctx, e.Body); err != nil {
				return err
			}
		}
	}
	return nil
}
