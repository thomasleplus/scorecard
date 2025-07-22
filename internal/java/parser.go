package java

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ossf/scorecard/v5/internal/java/java20"
)

func ParseFile(content []byte) {
	is := antlr.NewInputStream(string(content))
	lexer := java20.NewJava20Lexer(is.)
	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}
		fmt.Printf("%s (%q)\n",
			lexer.SymbolicNames[t.GetTokenType()], t.GetText())
	}
}