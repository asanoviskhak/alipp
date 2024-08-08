package ast

import (
	"bytes"

	"github.com/asanoviskhak/alipp/src/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (program *Program) String() string {
	var out bytes.Buffer

	for _, s := range program.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// let statement
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (statement *LetStatement) statementNode() {}
func (statement *LetStatement) TokenLiteral() string {
	return statement.Token.Literal
}

func (statement *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(statement.TokenLiteral() + " ")
	out.WriteString(statement.Name.String())
	out.WriteString(" = ")

	if statement.Value != nil {
		out.WriteString(statement.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (identifier *Identifier) expressionNode() {}
func (identifier *Identifier) TokenLiteral() string {
	return identifier.Token.Literal
}
func (identifier *Identifier) String() string { return identifier.Value }

// return statement
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (returnStatement *ReturnStatement) statementNode() {}
func (returnStatement *ReturnStatement) TokenLiteral() string {
	return returnStatement.Token.Literal
}

func (statement *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(statement.TokenLiteral() + " ")

	if statement.ReturnValue != nil {
		out.WriteString(statement.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// expression statement
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (expressionStatement *ExpressionStatement) statementNode() {}
func (expressionStatement *ExpressionStatement) TokenLiteral() string {
	return expressionStatement.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// Integer literal
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (integerLiteral *IntegerLiteral) expressionNode() {}
func (integerLiteral *IntegerLiteral) TokenLiteral() string {
	return integerLiteral.Token.Literal
}
func (integerLiteral *IntegerLiteral) String() string {
	return integerLiteral.Token.Literal
}
