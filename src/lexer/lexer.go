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


func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexerInstance *Lexer) readIdentifier() string {
    position := lexerInstance.position
    for isLetter(lexerInstance.ch) {
        lexerInstance.readChar()
    }
    return lexerInstance.input[position:lexerInstance.position]
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

func (lexerInstance *Lexer) consumeWhitespace() {
	for lexerInstance.ch == ' ' || lexerInstance.ch == '\t' || lexerInstance.ch == '\n' || lexerInstance.ch == '\r' {
		lexerInstance.readChar()
	}
}

func (lexerInstance *Lexer) readNumber() string {
	position := lexerInstance.position
	for isDigit(lexerInstance.ch) {
		lexerInstance.readChar()
	}

	return lexerInstance.input[position:lexerInstance.position]
}

func (lexerInstance *Lexer) peekChar() byte {
	if lexerInstance.readPosition >= len(lexerInstance.input) {
		return 0
	} else {
		return lexerInstance.input[lexerInstance.readPosition]
	}
}

func (lexerInstance *Lexer) NextToken() token.Token {
	var tok token.Token

	lexerInstance.consumeWhitespace()

	switch lexerInstance.ch {
	case '=':
		if lexerInstance.peekChar() == '=' {
			ch := lexerInstance.ch
			lexerInstance.readChar()
			literal := string(ch) + string(lexerInstance.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, lexerInstance.ch)
		}
	case '+':
		tok = newToken(token.PLUS, lexerInstance.ch)
	case '-':
		tok = newToken(token.MINUS, lexerInstance.ch)
	case '!':
		if lexerInstance.peekChar() == '=' {
			ch := lexerInstance.ch
			lexerInstance.readChar()
			literal := string(ch) + string(lexerInstance.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.EXCLAMATION, lexerInstance.ch)
		}
	case '/':
		tok = newToken(token.SLASH, lexerInstance.ch)
	case '*':
		tok = newToken(token.ASTERISK, lexerInstance.ch)
	case '<':
		tok = newToken(token.LT, lexerInstance.ch)
	case '>':
		tok = newToken(token.GT, lexerInstance.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexerInstance.ch)
	case '(':
		tok = newToken(token.LPAREN, lexerInstance.ch)
    case ')':
        tok = newToken(token.RPAREN, lexerInstance.ch)
    case ',':
        tok = newToken(token.COMMA, lexerInstance.ch)
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
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(lexerInstance.ch) {
			tok.Type = token.INT
			tok.Literal = lexerInstance.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexerInstance.ch)
		}
	}
	// Before returning the token we advance our pointers into the 
	// input so when we call NextToken() again the lexerInstance.ch field is already updated.
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



