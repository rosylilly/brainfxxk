package ast

type Program struct {
	Expressions []Expression
}

func (p *Program) String() string {
	s := ""
	for _, e := range p.Expressions {
		s += e.String()
	}
	return s
}
