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
	ASSIGN 	= "ASSIGN"
	PLUS 	= "+"

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

	// Excerpt From
	// Writing An Interpreter In Go
	// Thorsten Ball
	// This material may be protected by copyright.
)