package lexer

import (
	"io"
)

type Lexer struct {
	pos    int
	reader io.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		pos:    0,
		reader: reader,
	}
}

func (l *Lexer) Next() (*Token, error) {
	b := make([]byte, 1)
	n, err := l.reader.Read(b)
	if err != nil {
		return nil, err
	}

	tokenType := CommentToken
	if t, ok := ByteToTokenType[b[0]]; ok {
		tokenType = t
	}

	token := &Token{
		Type: tokenType,
		Byte: b[0],
		Pos:  l.pos,
	}

	l.pos += n

	return token, nil
}
