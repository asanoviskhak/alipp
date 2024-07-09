package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"
	IDENT 	= "IDENT"
	INT 	= "INT"

	// Operators
	ASSIGN   	= "="
    PLUS     	= "+"
    MINUS    	= "-"
    EXCLAMATION = "!"
    ASTERISK 	= "*"
    SLASH   	= "/"
	LT		 	= "<"
	GT 		 	= ">"

    // Delimiters
    COMMA     = ","
    SEMICOLON = ";"

    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET      = "LET"
	TRUE     = "TRUE"
    FALSE    = "FALSE"
    IF       = "IF"
    ELSE     = "ELSE"
    RETURN   = "RETURN"

	// Excerpt From
	// Writing An Interpreter In Go
	// Thorsten Ball
	// This material may be protected by copyright.
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
	"true": TRUE,
	"false":  FALSE,
    "if":     IF,
    "else":   ELSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}