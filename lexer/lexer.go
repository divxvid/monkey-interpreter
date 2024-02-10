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
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}
