package lexer

import (
	"testing"
)

func TestValidMedley(t *testing.T) {
	input := "machin123    def   defextern <= 123  extern 456hello  #comment def"
	targetResults := []KaleidoTokenContext{
		{KTokenIdentifier, "machin123"},
		{KTokenDef, ""},
		{KTokenIdentifier, "defextern"},
		{KTokenSymbol, "<"},
		{KTokenSymbol, "="},
		{KTokenNumber, "123"},
		{KTokenExtern, ""},
		{KTokenNumber, "456"},
		{KTokenIdentifier, "hello"},
		{KTokenEOF, ""},
	}
	lexer := NewKaleidoLexer(input)
	for i := 0; i < len(targetResults); i++ {
		result := lexer.NextToken()
		if result.Token != targetResults[i].Token || result.Value != targetResults[i].Value {
			t.Fatalf("Was waiting for: %v but creceived: %v", result, &targetResults[i])
		}
	}

}