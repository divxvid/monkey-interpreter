package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and literals
	IDENTIFIER = "IDENTIFIER" //general identifiers
	INT        = "INT"        //integers

	//Operators
	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	SLASH       = "/"
	ASTERISK    = "*"
	LESSTHAN    = "<"
	GREATERTHAN = ">"
	NOT         = "!"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "if"
	ELSE     = "else"
	TRUE     = "true"
	FALSE    = "false"
	RETURN   = "return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdentifier(identifier string) TokenType {
	//if the identifier is a valid keyword, we return the keyword's type
	if tok, ok := keywords[identifier]; ok {
		return tok
	}

	//otherwise we return the Identifier type
	return IDENTIFIER
}
