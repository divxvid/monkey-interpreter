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
	Token           token.Token //this will store the token.LET token
	IdentifierName  *Identifier
	ExpressionValue Expression
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	//returns back the LET token literal in string
	return ls.Token.Literal
}

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
