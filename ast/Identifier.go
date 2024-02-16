package ast

import "github.com/divxvid/monkey-interpreter/token"

// Identifier struct is an expression because it generates a value when used anywhere
// other than assignment
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}
