package ast

type Node interface {
	TokenLiteral() string
	String() string
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
