package lexer

type TokenType byte

const (
	CommentToken TokenType = iota
	PointerIncrementToken
	PointerDecrementToken
	ValueIncrementToken
	ValueDecrementToken
	OutputToken
	InputToken
	WhileStartToken
	WhileEndToken
)

var (
	ByteToTokenType = map[byte]TokenType{
		'>': PointerIncrementToken,
		'<': PointerDecrementToken,
		'+': ValueIncrementToken,
		'-': ValueDecrementToken,
		'.': OutputToken,
		',': InputToken,
		'[': WhileStartToken,
		']': WhileEndToken,
	}
)

type Token struct {
	Type TokenType
	Byte byte
	Pos  int
}
