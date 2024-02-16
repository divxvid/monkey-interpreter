package ast

import "github.com/divxvid/monkey-interpreter/token"

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
