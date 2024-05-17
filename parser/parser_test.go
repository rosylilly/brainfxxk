package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/rosylilly/brainfxxk/ast"
	"github.com/rosylilly/brainfxxk/lexer"
	"github.com/rosylilly/brainfxxk/parser"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		input    string
		expected *ast.Program
	}{
		{
			input: "+-><.,[]",
			expected: &ast.Program{
				Expressions: []ast.Expression{
					&ast.ValueIncrementExpression{Pos: 0},
					&ast.ValueDecrementExpression{Pos: 1},
					&ast.PointerIncrementExpression{Pos: 2},
					&ast.PointerDecrementExpression{Pos: 3},
					&ast.OutputExpression{Pos: 4},
					&ast.InputExpression{Pos: 5},
					&ast.WhileExpression{
						StartPosition: 6,
						EndPosition:   7,
						Body:          []ast.Expression{},
					},
				},
			},
		},
		{
			input: "+[->[+-<]>]",
			expected: &ast.Program{
				Expressions: []ast.Expression{
					&ast.ValueIncrementExpression{Pos: 0},
					&ast.WhileExpression{
						StartPosition: 1,
						EndPosition:   10,
						Body: []ast.Expression{
							&ast.ValueDecrementExpression{Pos: 2},
							&ast.PointerIncrementExpression{Pos: 3},
							&ast.WhileExpression{
								StartPosition: 4,
								EndPosition:   8,
								Body: []ast.Expression{
									&ast.ValueIncrementExpression{Pos: 5},
									&ast.ValueDecrementExpression{Pos: 6},
									&ast.PointerDecrementExpression{Pos: 7},
								},
							},
							&ast.PointerIncrementExpression{Pos: 9},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			program, err := parser.Parse(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			if program.String() != tc.expected.String() {
				t.Errorf("got: %v, expected: %v", program.String(), tc.expected.String())
			}

			if !reflect.DeepEqual(program, tc.expected) {
				t.Errorf("got: %v, expected: %v", program, tc.expected)
			}
		})
	}
}

func TestParserError(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "+[",
			expected: "syntax error: unclosed while block",
		},
		{
			input:    "+]",
			expected: "syntax error: unexpected token ] at 1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			l := lexer.NewLexer(strings.NewReader(tc.input))
			p := parser.NewParser(l)

			_, err := p.Parse()
			if err == nil {
				t.Fatal("expected error, but got nil")
			}

			if err.Error() != tc.expected {
				t.Errorf("got: %v, expected: %v", err.Error(), tc.expected)
			}
		})
	}
}

func TestParserPrinter(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "+-><.,[]",
			expected: `+-><.,[]`,
		},
		{
			input:    "+[->[+-<]>]",
			expected: `+[->[+-<]>]`,
		},
		{
			input:    "+!!!![->[+-<]>]",
			expected: `+!!!![->[+-<]>]`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			program, err := parser.Parse(strings.NewReader(tc.input))
			if err != nil {
				t.Fatal(err)
			}

			if program.String() != tc.expected {
				t.Errorf("got: %v, expected: %v", program.String(), tc.expected)
			}
		})
	}
}
