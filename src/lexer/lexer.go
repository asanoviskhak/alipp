package lexer

import (
	"testing"

	"github.com/asanoviskhak/alipp/src/token/token.go"
)

func TestNextToken(testing *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
	}
}