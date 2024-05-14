package lexer_test

import (
	"strings"
	"testing"

	"github.com/rosylilly/brainfxxk/lexer"
)

func TestLexer(t *testing.T) {
	testCases := []struct {
		input    string
		expected []*lexer.Token
	}{
		{
			input: "+-><.,[]",
			expected: []*lexer.Token{
				{Type: lexer.ValueIncrementToken, Byte: '+', Pos: 0},
				{Type: lexer.ValueDecrementToken, Byte: '-', Pos: 1},
				{Type: lexer.PointerIncrementToken, Byte: '>', Pos: 2},
				{Type: lexer.PointerDecrementToken, Byte: '<', Pos: 3},
				{Type: lexer.OutputToken, Byte: '.', Pos: 4},
				{Type: lexer.InputToken, Byte: ',', Pos: 5},
				{Type: lexer.WhileStartToken, Byte: '[', Pos: 6},
				{Type: lexer.WhileEndToken, Byte: ']', Pos: 7},
			},
		},
		{
			input: "+++++",
			expected: []*lexer.Token{
				{Type: lexer.ValueIncrementToken, Byte: '+', Pos: 0},
				{Type: lexer.ValueIncrementToken, Byte: '+', Pos: 1},
				{Type: lexer.ValueIncrementToken, Byte: '+', Pos: 2},
				{Type: lexer.ValueIncrementToken, Byte: '+', Pos: 3},
				{Type: lexer.ValueIncrementToken, Byte: '+', Pos: 4},
			},
		},
		{
			input: "!@#$%^&*()", // invalid characters
			expected: []*lexer.Token{
				{Type: lexer.CommentToken, Byte: '!', Pos: 0},
				{Type: lexer.CommentToken, Byte: '@', Pos: 1},
				{Type: lexer.CommentToken, Byte: '#', Pos: 2},
				{Type: lexer.CommentToken, Byte: '$', Pos: 3},
				{Type: lexer.CommentToken, Byte: '%', Pos: 4},
				{Type: lexer.CommentToken, Byte: '^', Pos: 5},
				{Type: lexer.CommentToken, Byte: '&', Pos: 6},
				{Type: lexer.CommentToken, Byte: '*', Pos: 7},
				{Type: lexer.CommentToken, Byte: '(', Pos: 8},
				{Type: lexer.CommentToken, Byte: ')', Pos: 9},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			l := lexer.NewLexer(strings.NewReader(tc.input))

			for _, expected := range tc.expected {
				token, err := l.Next()
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}

				if token.Type != expected.Type {
					t.Errorf("expected type %v, but got %v", expected.Type, token.Type)
				}

				if token.Byte != expected.Byte {
					t.Errorf("expected byte %v, but got %v", expected.Byte, token.Byte)
				}

				if token.Pos != expected.Pos {
					t.Errorf("expected pos %v, but got %v", expected.Pos, token.Pos)
				}
			}
		})
	}
}
