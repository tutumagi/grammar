package grammar

type Symbol = string

// nonterminal is UpperCase
// terminal is lowercase
const (
	epsilonS        = "ε" // ε U+03B5
	rightEndMarkerS = "$"
	productionS     = "->"
	alternateS      = "|"
)

func isEpsilon(sym Symbol) bool {
	return sym == epsilonS
}

func isTerminal(sym Symbol) bool {
	c := sym[0]
	return !(c >= 'A' && c <= 'Z')
}

func isNonTerminal(sym Symbol) bool {
	return !(isEpsilon(sym) || isTerminal(sym))
}

type SymbolSet map[Symbol]struct{}

func newSymbolSet(syms ...Symbol) SymbolSet {
	s := make(SymbolSet)
	s.add(syms...)
	return s
}

func (set SymbolSet) add(ss ...Symbol) {
	for _, s := range ss {
		set[s] = struct{}{}
	}
}
func (set SymbolSet) remove(s Symbol) {
	delete(set, s)
}
func (set SymbolSet) contain(s Symbol) bool {
	_, ok := set[s]
	return ok
}
func (set SymbolSet) union(other SymbolSet) SymbolSet {
	for m := range other {
		set.add(m)
	}
	return set
}
func (set SymbolSet) intersect(other SymbolSet) SymbolSet {
	result := make(SymbolSet)
	for m := range other {
		if set.contain(m) {
			result.add(m)
		}
	}
	return result
}

func (set SymbolSet) disjoint(other SymbolSet) bool {
	return len(set.intersect(other)) == 0
}

const (
	NotFound = -1
)

func indexOfSymbolList(sym Symbol, list []Symbol) int {
	ret := NotFound
	for i, s := range list {
		if s == sym {
			ret = i
			break
		}
	}
	return ret
}
