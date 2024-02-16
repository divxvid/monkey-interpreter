package ast

import "bytes"

// this will be the root node of our AST
// this node contains all the statements of our program
type Program struct {
	Statements []Statement
}

// implementing TokenLiteral() makes Program implement the Node interface
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// concatenate all the string repr of all the Statements
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}
