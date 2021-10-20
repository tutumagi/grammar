package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = `
E   -> T E'
E'  -> + T E' | e
T   -> F T'
T'  -> * F T' | e
F   -> ( E ) | id
`

func TestMakeRules(t *testing.T) {
	expectProductions := map[Symbol][]*Production{
		"E": {
			newProduction("E").RHS("T", "E'"),
		},
		"E'": {
			newProduction("E'").RHS("+", "T", "E'"),
			newProduction("E'").RHS("e"),
		},
		"T": {
			newProduction("T").RHS("F", "T'"),
		},
		"T'": {
			newProduction("T'").RHS("*", "F", "T'"),
			newProduction("T'").RHS("e"),
		},
		"F": {
			newProduction("F").RHS("(", "E", ")"),
			newProduction("F").RHS("id"),
		},
	}

	g := NewGrammar(testData)
	productions := g.makeProductions()
	assert.Equal(t, expectProductions, productions)

	g.makeFirstSet()
	expectFirstSet := map[Symbol]SymbolSet{
		"E":  newSymbolSet("(", "id"),
		"E'": newSymbolSet("+", epsilonS),
		"T":  newSymbolSet("(", "id"),
		"T'": newSymbolSet("*", epsilonS),
		"F":  newSymbolSet("(", "id"),
	}

	for sym, expectFS := range expectFirstSet {
		assert.Equal(t, g.firstSet[sym], expectFS)
	}

	// spew.Dump(g.firstSet)
}

func TestMakeFirst(t *testing.T) {

}
