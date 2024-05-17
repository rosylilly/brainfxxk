package parser

import (
	"errors"
	"fmt"
	"io"

	"github.com/rosylilly/brainfxxk/ast"
	"github.com/rosylilly/brainfxxk/lexer"
)

var (
	ErrInvalidSyntax = errors.New("syntax error")
)

type Parser struct {
	lexer *lexer.Lexer
}

func Parse(r io.Reader) (*ast.Program, error) {
	l := lexer.NewLexer(r)
	p := NewParser(l)
	return p.Parse()
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		lexer: l,
	}
}

func (p *Parser) Parse() (*ast.Program, error) {
	exprs := []ast.Expression{}
	stack := [][]ast.Expression{exprs}

	for {
		token, err := p.lexer.Next()
		if err != nil && !errors.Is(err, io.EOF) {
			return nil, err
		}
		if token == nil {
			break
		}

		switch token.Type {
		case lexer.PointerIncrementToken:
			exprs = append(exprs, &ast.PointerIncrementExpression{Pos: token.Pos})
		case lexer.PointerDecrementToken:
			exprs = append(exprs, &ast.PointerDecrementExpression{Pos: token.Pos})
		case lexer.ValueIncrementToken:
			exprs = append(exprs, &ast.ValueIncrementExpression{Pos: token.Pos})
		case lexer.ValueDecrementToken:
			exprs = append(exprs, &ast.ValueDecrementExpression{Pos: token.Pos})
		case lexer.OutputToken:
			exprs = append(exprs, &ast.OutputExpression{Pos: token.Pos})
		case lexer.InputToken:
			exprs = append(exprs, &ast.InputExpression{Pos: token.Pos})
		case lexer.WhileStartToken:
			expr := &ast.WhileExpression{
				StartPosition: token.Pos,
				EndPosition:   token.Pos,
				Body:          []ast.Expression{},
			}
			exprs = append(exprs, expr)
			stack = append(stack, exprs)
			exprs = expr.Body
		case lexer.WhileEndToken:
			if len(stack) == 1 {
				return nil, fmt.Errorf("%w: unexpected token %c at %d", ErrInvalidSyntax, token.Byte, token.Pos)
			}
			body := exprs
			exprs = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if len(exprs) == 0 {
				return nil, fmt.Errorf("%w: empty while block at %d", ErrInvalidSyntax, token.Pos)
			}

			expr := exprs[len(exprs)-1]
			if we, ok := expr.(*ast.WhileExpression); !ok {
				return nil, fmt.Errorf("%w: unexpected token %c at %d", ErrInvalidSyntax, token.Byte, token.Pos)
			} else {
				we.EndPosition = token.Pos
				we.Body = body
			}
		case lexer.CommentToken:
			var expr ast.Expression
			if len(exprs) > 0 {
				expr = exprs[len(exprs)-1]
			}

			if cm, ok := expr.(*ast.Comment); !ok {
				exprs = append(exprs, &ast.Comment{Start: token.Pos, End: token.Pos, Body: []byte{token.Byte}})
			} else {
				cm.Body = append(cm.Body, token.Byte)
				cm.End = token.Pos
			}
		}
	}

	if len(stack) != 1 {
		return nil, fmt.Errorf("%w: unclosed while block", ErrInvalidSyntax)
	}

	return &ast.Program{
		Expressions: exprs,
	}, nil
}
