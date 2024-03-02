package ast

import (
	"bytes"
	"strings"

	"github.com/divxvid/monkey-interpreter/token"
)

type CallExpression struct {
	Token token.Token //the '(' keyword
	//this can either be an Identifier or FunctionLiteral
	//Identifier example: add(2, 3)
	//FunctionLiteral example: fn(x, y) { x+y; }(2, 3)
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}

func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}

func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
