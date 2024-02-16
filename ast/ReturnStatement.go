package ast

import (
	"bytes"

	"github.com/divxvid/monkey-interpreter/token"
)

// return statements are of the type: return <expression>;
type ReturnStatement struct {
	Token       token.Token //this will store the token.RETURN token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	//returns back the RETURN token literal in string
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
