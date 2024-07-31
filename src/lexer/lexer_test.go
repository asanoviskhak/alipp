package lexer

import (
	"testing"

	"github.com/asanoviskhak/alipp/src/token"
)

func TestNextToken(testing *testing.T) {
	input := `сакта five = 5;
сакта ten = 10;
сакта add = функ(x, y) {x + y;};
сакта result = add(five, ten);
!-/*5;
5 < 10 > 5;
эгер (5 < 10) {
кайтар туура;
} же {
кайтар ката;
}
10 == 10;
10 != 9;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "сакта"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "сакта"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "сакта"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "функ"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "сакта"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EXCLAMATION, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "эгер"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "кайтар"},
		{token.TRUE, "туура"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "же"},
		{token.LBRACE, "{"},
		{token.RETURN, "кайтар"},
		{token.FALSE, "ката"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexerInstance := New(input)

	for index, tokenTest := range tests {
		tokenNext := lexerInstance.NextToken()
		if tokenNext.Type != tokenTest.expectedType {
			testing.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				index, tokenTest.expectedType, tokenNext.Type)
		}

		if tokenNext.Literal != tokenTest.expectedLiteral {
			testing.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				index, tokenTest.expectedLiteral, tokenNext.Literal)
		}
	}
}
