package ast

import "github.com/divxvid/monkey-interpreter/token"

// let statements are of type: let <identifier> = <expression>;
type LetStatement struct {
	Token           token.Token //this will store the token.LET token
	IdentifierName  *Identifier
	ExpressionValue Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	//returns back the LET token literal in string
	return ls.Token.Literal
}
