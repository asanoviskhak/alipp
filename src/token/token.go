package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ТУУРА_ЭМЕС"
	EOF     = "БҮТТҮ"
	IDENT   = "ИДЕНТИФИКАТОР"
	INT     = "БҮТҮН_САН"

	// Operators
	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	EXCLAMATION = "!"
	ASTERISK    = "*"
	SLASH       = "/"
	LT          = "<"
	GT          = ">"
	EQ          = "=="
	NOT_EQ      = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "ФУНКЦИЯ"
	LET      = "САКТА"
	TRUE     = "ТУУРА"
	FALSE    = "КАТА"
	IF       = "ЭГЕР"
	ELSE     = "ЖЕ"
	RETURN   = "КАЙТАР"

	// Excerpt From
	// Writing An Interpreter In Go
	// Thorsten Ball
	// This material may be protected by copyright.
)

var keywords = map[string]TokenType{
	"функция": FUNCTION,
	"функ":    FUNCTION,
	"сакта":   LET,
	"туура":   TRUE,
	"ката":    FALSE,
	"эгер":    IF,
	"же":      ELSE,
	"кайтар":  RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
