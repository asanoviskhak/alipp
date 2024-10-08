package parser

import (
	"fmt"
	"testing"

	"github.com/asanoviskhak/alipp/src/ast"
	"github.com/asanoviskhak/alipp/src/lexer"
)

func checkParserErrors(t *testing.T, parser *Parser) {
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

func TestLetStatement(t *testing.T) {
	input := `сакта x = 3;
сакта y = 12;
сакта bishkek = 312;`

	lexerInstance := lexer.New(input)
	parserInstance := NewParser(lexerInstance)

	program := parserInstance.ParseProgram()
	checkParserErrors(t, parserInstance)

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

func testLetStatement(t *testing.T, statement ast.Statement, name string) bool {
	if statement.TokenLiteral() != "сакта" {
		t.Errorf("statement.TokenLiteral not 'сакта'. got=%q", statement.TokenLiteral())
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
	input := `кайтар 3;
кайтар 7;
кайтар 891011;`

	lexerInstance := lexer.New(input)
	parser := NewParser(lexerInstance)

	program := parser.ParseProgram()
	checkParserErrors(t, parser)

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

func TestIdentifierExpression(t *testing.T) {
	input := "салам;"

	lexer := lexer.New(input)
	parser := NewParser(lexer)
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if statementsLength := len(program.Statements); statementsLength != 1 {
		t.Fatalf("program doesn't have enough statements. got=%d", statementsLength)
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	identifier, ok := statement.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("expression is not a *ast.Identifier. got=%T", statement.Expression)
	}

	testValue := "салам"
	if currentValue := identifier.Value; currentValue != testValue {
		t.Errorf("identifier.Value not %s. got=%s", testValue, currentValue)
	}
	if currentLiteral := identifier.TokenLiteral(); currentLiteral != testValue {
		t.Errorf("identifier.TokenLiteral() is not %s. got=%s", testValue, currentLiteral)
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	lexer := lexer.New(input)
	parser := NewParser(lexer)
	program := parser.ParseProgram()
	checkParserErrors(t, parser)

	if statementsLength := len(program.Statements); statementsLength != 1 {
		t.Fatalf("program doesn't have enough statements. got=%d", statementsLength)
	}

	statement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("statement is not ast.ExpressionStatement. got=%T", program.Statements[0])
	}

	integerLiteral, ok := statement.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("expression is not a *ast.IntegerLiteral. got=%T", statement.Expression)
	}

	if currentValue := integerLiteral.Value; currentValue != 5 {
		t.Errorf("integerLiteral.Value not %d. got=%d", 5, currentValue)
	}

	testValue := "5"
	if currentLiteral := integerLiteral.TokenLiteral(); currentLiteral != testValue {
		t.Errorf("integerLiteral.TokenLiteral() is not %s. got=%s", testValue, currentLiteral)
	}
}

func TestParsingPrefix(test *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, currentTest := range prefixTests {
		lexer := lexer.New(currentTest.input)
		parser := NewParser(lexer)
		program := parser.ParseProgram()
		checkParserErrors(test, parser)

		if len(program.Statements) != 1 {
			test.Fatalf("program.Statements does not contain %d statements, got=%T", 1, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			test.Fatalf("program.Statements[0] is not ast.ExpressionStatement, got=%T", program.Statements[0])
		}

		expression, ok := statement.Expression.(*ast.PrefixExpression)
		if !ok {
			test.Fatalf("statement is not ast.PrefixExpression, got=%T", statement.Expression)
		}

		if expression.Operator != currentTest.operator {
			test.Fatalf("expression.Operator is not '%s'. got=%s", currentTest.operator, expression.Operator)
		}

		if !testIntegerLiteral(test, expression.Right, currentTest.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, integerLiteral ast.Expression, value int64) bool {
	integer, ok := integerLiteral.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("integerLiteral is not *ast.IntegerLiteral. got=%T", integerLiteral)
		return false
	}

	if integer.Value != value {
		t.Errorf("integer.Value is not %d. got=%d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral() is not %d. got=%s", value, integer.TokenLiteral())
		return false
	}

	return true
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		lexer := lexer.New(tt.input)
		parser := NewParser(lexer)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		statement, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		infixExpression, ok := statement.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("infixExpression is not ast.InfixExpression. got=%T", statement.Expression)
		}

		if !testIntegerLiteral(t, infixExpression.Left, tt.leftValue) {
			return
		}

		if infixExpression.Operator != tt.operator {
			t.Fatalf("infixExpression.Operator is not '%s'. got=%s",
				tt.operator, infixExpression.Operator)
		}

		if !testIntegerLiteral(t, infixExpression.Right, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-a * b",
			"((-a) * b)",
		},
		{
			"!-a",
			"(!(-a))",
		},
		{
			"a + b + c",
			"((a + b) + c)",
		},
		{
			"a + b - c",
			"((a + b) - c)",
		},
		{
			"a * b * c",
			"((a * b) * c)",
		},
		{
			"a * b / c",
			"((a * b) / c)",
		},
		{
			"a + b / c",
			"(a + (b / c))",
		},
		{
			"a + b * c + d / e - f",
			"(((a + (b * c)) + (d / e)) - f)",
		},
		{
			"3 + 4; -5 * 5",
			"(3 + 4)((-5) * 5)",
		},
		{
			"5 > 4 == 3 < 4",
			"((5 > 4) == (3 < 4))",
		},
		{
			"5 < 4 != 3 > 4",
			"((5 < 4) != (3 > 4))",
		},
		{
			"3 + 4 * 5 == 3 * 1 + 4 * 5",
			"((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
	}
	for _, tt := range tests {
		lexer := lexer.New(tt.input)
		parser := NewParser(lexer)
		program := parser.ParseProgram()
		checkParserErrors(t, parser)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
