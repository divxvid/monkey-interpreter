package ast

import "github.com/divxvid/monkey-interpreter/token"

type Node interface {
	TokenLiteral() string
}

// In crude terms, a statement is a block of code that doesn't return any value
type Statement interface {
	Node
	statementNode()
}

// Expression on the other hand is something that returns back some value
type Expression interface {
	Node
	expressionNode()
}

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

// let statements are of type: let <identifier> = <expression>;
type LetStatement struct {
	token token.Token //this will store the token.LET token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	//returns back the LET token literal in string
	return ls.token.Literal
}

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
