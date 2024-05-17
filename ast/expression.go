package ast

type Expression interface {
	StartPos() int
	EndPos() int
	Bytes() []byte
	String() string
}

type PointerIncrementExpression struct {
	Pos int
}

func (e *PointerIncrementExpression) StartPos() int {
	return e.Pos
}

func (e *PointerIncrementExpression) EndPos() int {
	return e.Pos
}

func (e *PointerIncrementExpression) Bytes() []byte {
	return []byte{'>'}
}

func (e *PointerIncrementExpression) String() string {
	return string(e.Bytes())
}

type PointerDecrementExpression struct {
	Pos int
}

func (e *PointerDecrementExpression) StartPos() int {
	return e.Pos
}

func (e *PointerDecrementExpression) EndPos() int {
	return e.Pos
}

func (e *PointerDecrementExpression) Bytes() []byte {
	return []byte{'<'}
}

func (e *PointerDecrementExpression) String() string {
	return string(e.Bytes())
}

type ValueIncrementExpression struct {
	Pos int
}

func (e *ValueIncrementExpression) StartPos() int {
	return e.Pos
}

func (e *ValueIncrementExpression) EndPos() int {
	return e.Pos
}

func (e *ValueIncrementExpression) Bytes() []byte {
	return []byte{'+'}
}

func (e *ValueIncrementExpression) String() string {
	return string(e.Bytes())
}

type ValueDecrementExpression struct {
	Pos int
}

func (e *ValueDecrementExpression) StartPos() int {
	return e.Pos
}

func (e *ValueDecrementExpression) EndPos() int {
	return e.Pos
}

func (e *ValueDecrementExpression) Bytes() []byte {
	return []byte{'-'}
}

func (e *ValueDecrementExpression) String() string {
	return string(e.Bytes())
}

type OutputExpression struct {
	Pos int
}

func (e *OutputExpression) StartPos() int {
	return e.Pos
}

func (e *OutputExpression) EndPos() int {
	return e.Pos
}

func (e *OutputExpression) Bytes() []byte {
	return []byte{'.'}
}

func (e *OutputExpression) String() string {
	return string(e.Bytes())
}

type InputExpression struct {
	Pos int
}

func (e *InputExpression) StartPos() int {
	return e.Pos
}

func (e *InputExpression) EndPos() int {
	return e.Pos
}

func (e *InputExpression) Bytes() []byte {
	return []byte{','}
}

func (e *InputExpression) String() string {
	return string(e.Bytes())
}

type WhileExpression struct {
	StartPosition int
	EndPosition   int
	Body          []Expression
}

func (e *WhileExpression) StartPos() int {
	return e.StartPosition
}

func (e *WhileExpression) EndPos() int {
	return e.EndPosition
}

func (e *WhileExpression) Bytes() []byte {
	b := []byte{'['}
	for _, expr := range e.Body {
		b = append(b, expr.Bytes()...)
	}
	b = append(b, ']')
	return b
}

func (e *WhileExpression) String() string {
	return string(e.Bytes())
}

type Comment struct {
	Start int
	End   int
	Body  []byte
}

func (c *Comment) StartPos() int {
	return c.Start
}

func (c *Comment) EndPos() int {
	return c.End
}

func (c *Comment) Bytes() []byte {
	return c.Body
}

func (c *Comment) String() string {
	return string(c.Body)
}
