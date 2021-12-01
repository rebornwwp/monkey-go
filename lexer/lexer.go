package lexer

import (
	"example.com/m/token"
)

func newToken(tokenType token.TokenType, l string) token.Token {
	return token.Token{Type: tokenType, Literal: l}
}

type Lexer struct {
	input        string
	position     int
	readPosition int //current reading position in input (after current char); next char
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tk token.Token

	// skip
	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		c := l.peekChar()
		if c == '=' {
			l.readChar()
			tk = newToken(token.EQ, "==")
		} else {
			tk = newToken(token.ASSIGN, string(l.ch))
		}
	case '+':
		tk = newToken(token.PLUS, string(l.ch))
	case '-':
		tk = newToken(token.MINUS, string(l.ch))
	case '!':
		c := l.peekChar()
		if c == '=' {
			l.readChar()
			tk = newToken(token.NOTEQ, "!=")
		} else {
			tk = newToken(token.BANG, string(l.ch))
		}
	case '*':
		tk = newToken(token.ASTERISK, string(l.ch))
	case '/':
		tk = newToken(token.SLASH, string(l.ch))
	case '<':
		tk = newToken(token.LT, string(l.ch))
	case '>':
		tk = newToken(token.GT, string(l.ch))
	case ',':
		tk = newToken(token.COMMA, string(l.ch))
	case ';':
		tk = newToken(token.SEMICOLON, string(l.ch))
	case '(':
		tk = newToken(token.LPAREN, string(l.ch))
	case ')':
		tk = newToken(token.RPAREN, string(l.ch))
	case '{':
		tk = newToken(token.LBRACE, string(l.ch))
	case '}':
		tk = newToken(token.RBRACE, string(l.ch))
	case '"':
		tk.Type = token.STRING
		tk.Literal = l.readString()
	case '[':
		tk = newToken(token.LBRACKET, string(l.ch))
	case ']':
		tk = newToken(token.RBRACKET, string(l.ch))
	case ':':
		tk = newToken(token.COLON, string(l.ch))
	case 0:
		tk.Type = token.EOF
		tk.Literal = ""
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			tkType := token.LookupIdent(ident)
			tk = newToken(tkType, ident)
			return tk
		} else if isDigit(l.ch) {
			num := l.readNumber()
			tk = newToken(token.INT, num)
			return tk
		}
		tk = newToken(token.ILLEGAL, "")
		return tk
	}

	l.readChar()
	return tk
}

func isLetter(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // the ASCII code for the "NUL" character to indicate there is no character to consume
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}
