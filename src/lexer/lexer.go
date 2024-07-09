package lexer

import (
	"github.com/asanoviskhak/alipp/src/token"
)

type Lexer struct {
	input 			string
	position 		int
	readPosition 	int
	ch 				byte
	chUnicode		rune
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[position:l.position]
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexerInstance *Lexer) NextToken() token.Token {
	var tok token.Token
	switch lexerInstance.ch {
	case '=':
		tok = newToken(token.ASSIGN, lexerInstance.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexerInstance.ch)
	case '(':
		tok = newToken(token.LPAREN, lexerInstance.ch)
    case ')':
        tok = newToken(token.RPAREN, lexerInstance.ch)
    case ',':
        tok = newToken(token.COMMA, lexerInstance.ch)
    case '+':
        tok = newToken(token.PLUS, lexerInstance.ch)
    case '{':
        tok = newToken(token.LBRACE, lexerInstance.ch)
    case '}':
        tok = newToken(token.RBRACE, lexerInstance.ch)
    case 0:
        tok.Literal = ""
        tok.Type = token.EOF
	default:
		if isLetter(lexerInstance.ch) {
			tok.Literal = lexerInstance.readIdentifier()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexerInstance.ch)
		}
	}

	lexerInstance.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
    return token.Token{Type: tokenType, Literal: string(ch)}
}

func New(input string) *Lexer {
	lexerInstance := &Lexer{input: input}
	lexerInstance.readChar()
	return lexerInstance
}

func NewUnicode(input string) *Lexer {
	lexerInstance := &Lexer{input: input}
	lexerInstance.readCharUnicode()
	return lexerInstance
}

func (lexerInstance *Lexer) readChar() {
	if lexerInstance.readPosition >= len(lexerInstance.input) {
		lexerInstance.ch = 0
	} else {
		lexerInstance.ch = lexerInstance.input[lexerInstance.readPosition]
	}

	lexerInstance.position = lexerInstance.readPosition
	lexerInstance.readPosition += 1
}

func (lexerInstance *Lexer) readCharUnicode() {
	if lexerInstance.readPosition >= len(lexerInstance.input) {
		lexerInstance.chUnicode = 0
	} else {
		lexerInstance.chUnicode = rune(lexerInstance.input[lexerInstance.readPosition])
	}

	lexerInstance.position = lexerInstance.readPosition
	lexerInstance.readPosition += 1
}

