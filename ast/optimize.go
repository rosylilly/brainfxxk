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
