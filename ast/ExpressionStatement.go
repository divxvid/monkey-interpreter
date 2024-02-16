package ast

import "github.com/divxvid/monkey-interpreter/token"

// this statement represents any expression that is freely written
// for example: x+2; <-- we can try this in a repl and this should work
// we need an expression statement to create a wrapper around such cases
// and allow it to be appended to the statements list inside Program struct
type ExpressionStatement struct {
	Token      token.Token //this will store the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
