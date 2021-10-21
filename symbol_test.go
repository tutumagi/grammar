package grammar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSymbolSet(t *testing.T) {
	lowerSet := newSymbolSet("a", "b", "c")
	upperSet := newSymbolSet("A", "B", "C")
	mixSet := newSymbolSet("A", "B", "C", "a", "b", "c")

	assert.Equal(t, true, lowerSet.disjoint(upperSet))
	assert.Equal(t, false, lowerSet.disjoint(mixSet))
	assert.ElementsMatch(t, []Symbol{"a", "b", "c"}, lowerSet.toList())
	assert.Equal(t, mixSet, newSymbolSet().union(lowerSet).union(upperSet))
	assert.Equal(t, true, lowerSet.contain("a"))
	assert.Equal(t, false, lowerSet.contain("A"))

	assert.Equal(t, newSymbolSet("a", "b", "c", "d"), lowerSet.add("d"))
	assert.Equal(t, newSymbolSet("a", "b", "d"), lowerSet.remove("c"))
}
