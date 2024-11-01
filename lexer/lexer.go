package lexer

import (
	"log"

	"github.com/apex-woot/monkey-v0/misc"
	"github.com/apex-woot/monkey-v0/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	misc.ParseAndApplyFlags()

	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
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
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.RT, l.ch)
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
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readLiteral(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			log.Printf("position - %d readPosition - %d ch — %s tokenLit - %s tokenType - %s \n", l.position, l.readPosition, string(l.ch), tok.Literal, tok.Type)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readLiteral(isDigit)
			log.Printf("position - %d readPosition - %d ch — %s tokenLit - %s tokenType - %s \n", l.position, l.readPosition, string(l.ch), tok.Literal, tok.Type)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
			log.Printf("position - %d readPosition - %d ch — %s tokenLit - %s tokenType - %s \n", l.position, l.readPosition, string(l.ch), tok.Literal, tok.Type)
		}
	}

	log.Printf("position - %d readPosition - %d ch — %s tokenLit - %s tokenType - %s \n", l.position, l.readPosition, string(l.ch), tok.Literal, tok.Type)
	l.readChar()
	return tok
}

func (l *Lexer) readLiteral(isCorrectLiteral func(byte) bool) string {
	log.Printf("READING LITERAL — %s\n", string(l.ch))
	position := l.position
	for isCorrectLiteral(l.ch) {
		l.readChar()
		log.Printf("position - %d readPosition - %d ch — %s\n", l.position, l.readPosition, string(l.ch))
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
