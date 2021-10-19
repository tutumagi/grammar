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
	expectRules := map[Symbol][]*Production{
		"E": {
			newProduction("E", "T", "E'"),
		},
		"E'": {
			newProduction("E'", "+", "T", "E'"),
			newProduction("E'", "e"),
		},
		"T": {
			newProduction("T", "F", "T'"),
		},
		"T'": {
			newProduction("T'", "*", "F", "T'"),
			newProduction("T'", "e"),
		},
		"F": {
			newProduction("F", "(", "E", ")"),
			newProduction("F", "id"),
		},
	}

	g := NewGrammar(testData)
	productions := g.makeProductions()
	assert.Equal(t, expectRules, productions)

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
