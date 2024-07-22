package parserunicode

import (
	"testing"

	ast "github.com/asanoviskhak/alipp/src/ast-unicode"
	lexer "github.com/asanoviskhak/alipp/src/lexer-unicode"
)

func TestLetStatement(t *testing.T) {
	input := `
		бер x = 3;
		бер y = 12;

		бер bishkek = 312;
	`

	lexerInstance := lexer.New(input)
	parserInstance := NewParser(lexerInstance)

	program := parserInstance.ParseProgram()
	checkParseErrors(t, parserInstance)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"bishkek"},
	}

	for i, test := range tests {
		statement := program.Statements[i]

		if !testLetStatement(t, statement, test.expectedIdentifier) {
			return
		}
	}
}

func checkParseErrors(t *testing.T, parser *Parser) {
	errors := parser.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "бер" {
		t.Errorf("statement.TokenLiteral not 'бер'. got=%q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)
	if !ok {
		t.Errorf("statement not *ast.LetStatement. got=%T", statement)
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name.TokenLiteral() not '%s'. got=%s", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
		кайтар 3;
		кайтар 7;
		кайтар 891011;
	`

	lexerInstance := lexer.New(input)
	parser := NewParser(lexerInstance)

	program := parser.ParseProgram()
	checkParseErrors(t, parser)

	if statementsLength := len(program.Statements); statementsLength != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", statementsLength)
	}

	for _, statement := range program.Statements {
		returnStatement, ok := statement.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("statement not *ast.ReturnStatement. got=%T", statement)
			continue
		}

		if tokenLiteral := returnStatement.TokenLiteral(); tokenLiteral != "кайтар" {
			t.Errorf("returnStmt.TokenLiteral not 'кайтар', got %q",
				tokenLiteral)
		}

	}
}
