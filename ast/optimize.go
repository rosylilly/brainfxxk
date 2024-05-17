package ast

type MOVE struct {
	Count       int
	Expressions []Expression
}

func (e *MOVE) StartPos() int {
	return e.Expressions[0].StartPos()
}

func (e *MOVE) EndPos() int {
	return e.Expressions[len(e.Expressions)-1].EndPos()
}

func (e *MOVE) Bytes() []byte {
	b := []byte{}
	for _, expr := range e.Expressions {
		b = append(b, expr.Bytes()...)
	}
	return b
}

func (e *MOVE) String() string {
	return string(e.Bytes())
}

type CALC struct {
	Value       int
	Expressions []Expression
}

func (e *CALC) StartPos() int {
	return e.Expressions[0].StartPos()
}

func (e *CALC) EndPos() int {
	return e.Expressions[len(e.Expressions)-1].EndPos()
}

func (e *CALC) Bytes() []byte {
	b := []byte{}
	for _, expr := range e.Expressions {
		b = append(b, expr.Bytes()...)
	}
	return b
}

func (e *CALC) String() string {
	return string(e.Bytes())
}

type ZERORESET struct {
	Pos int
}

func (e *ZERORESET) StartPos() int {
	return e.Pos
}

func (e *ZERORESET) EndPos() int {
	return e.Pos
}

func (e *ZERORESET) Bytes() []byte {
	return []byte{'0'}
}

func (e *ZERORESET) String() string {
	return string(e.Bytes())
}

const (
	ForwardDirection  = 1
	BackwardDirection = -1
)

type ZEROSHIFT struct {
	Pos  int
	Leap int
}

func (e *ZEROSHIFT) StartPos() int {
	return e.Pos
}

func (e *ZEROSHIFT) EndPos() int {
	return e.Pos
}

func (e *ZEROSHIFT) Bytes() []byte {
	return []byte{'S'}
}

func (e *ZEROSHIFT) String() string {
	return string(e.Bytes())
}

type COPY struct {
	Pos       int
	CopyPlace []int
}

func (e *COPY) StartPos() int {
	return e.Pos
}

func (e *COPY) EndPos() int {
	return e.Pos
}

func (e *COPY) Bytes() []byte {
	return []byte{'C'}
}

func (e *COPY) String() string {
	return string(e.Bytes())
}
