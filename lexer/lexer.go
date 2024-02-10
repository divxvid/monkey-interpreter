package lexer

import "github.com/divxvid/monkey-interpreter/token"

// position points to the character in the input that corresponds to the ch byte
// i.e. ch = input[position]
// readPosition is basically position+1 i.e. the next character to be read
// TODO: update the ch to take in a rune instead of a byte to support unicode
type Lexer struct {
	input        string
	position     int  //current position in input (points to current character)
	readPosition int  //current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           0,
	}

	l.readChar() //load in the first character of the input
	return l
}

// read the next character from the input and set it's value into the ch variable
// if we have reached the end of input, we store a zero(0) value instead
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 //when we have consumed all the input
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)

			//we do an early return because readIdentifier stops when we get a
			//non letter character, so we don't need to call l.readChar() again at the bottom
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT

			//same reason for early return as above
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// TODO: [if possible] Make a generic function that takes in a boolean function and returns and
// returns a slice of string. We can replace readIdentifier and readNumber with this function
func (l *Lexer) readIdentifier() string {
	previousPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[previousPosition:l.position]
}

// TODO: add functionality to read hex, octal and floats for future
func (l *Lexer) readNumber() string {
	previousPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[previousPosition:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
