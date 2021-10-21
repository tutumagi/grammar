package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPredictTable(t *testing.T) {
	terminals := []string{"a", "c", "b", "$", "e", "z", "f"}
	nonTerminals := []string{"S", "A", "B", "E"}

	predictTable := newPredictTable(terminals, nonTerminals)

	expectTerminals := []string{"a", "b", "c", "e", "f", "z", "$"}
	expectNonTerminals := []string{"S", "A", "B", "E"}
	assert.Equal(t, predictTable.terminals, expectTerminals)
	assert.Equal(t, predictTable.nonterminal, expectNonTerminals)
}
