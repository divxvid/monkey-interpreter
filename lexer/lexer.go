package lexer

// position points to the character in the input that corresponds to the ch byte
// i.e. ch = input[position]
// readPosition is basically position+1 i.e. the next character to be read
type Lexer struct {
	input        string
	position     int  //current position in input (points to current character)
	readPosition int  //current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	return &Lexer{
		input:        input,
		position:     0,
		readPosition: 0,
		ch:           0,
	}
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
