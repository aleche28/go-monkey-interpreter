package lexer

import "monkey/token"

type Lexer struct {
	input        string
	position     int  // curr pos in input, points to curr char
	readPosition int  // curr reading pos, after curr char (useful to peek next char)
	ch           byte // curr char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // kind of initialization
	return l
}

// read next char (stored in ch) and advance cursor,
// ch is set to 0 (nil) if eof
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ascii for nil
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
	// readPosition will always point to next position to read
	// position will always point to last position we read
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peek() == '=' {
			ch := l.ch
			l.readChar()
			tok.Literal = string(ch) + string(l.ch)
			tok.Type = token.EQ
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peek() == '=' {
			ch := l.ch
			l.readChar()
			tok.Literal = string(ch) + string(l.ch)
			tok.Type = token.NOT_EQ
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok // return cause readIdentifier last op is already a readChar (this could be handled better)
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar() // move to next pos
	return tok
}

func newToken(t token.TokenType, c byte) token.Token {
	return token.Token{Type: t, Literal: string(c)}
}

func isLetter(c byte) bool {
	return ('A' <= c && c <= 'Z') || ('a' <= c && c <= 'z') || c == '_'
}

func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) skipWhitespace() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func isWhitespace(c byte) bool {
	return c == ' ' || c == '\n' || c == '\t' || c == '\r'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}

func (l *Lexer) peek() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}
