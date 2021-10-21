package grammar

import "strings"

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
	return !(isEpsilon(sym) || isNonTerminal(sym))
}

func isNonTerminal(sym Symbol) bool {
	c := sym[0]
	return (c >= 'A' && c <= 'Z')
}

type SymbolSet map[Symbol]struct{}

func newSymbolSet(syms ...Symbol) SymbolSet {
	s := make(SymbolSet)
	s.add(syms...)
	return s
}

func (set SymbolSet) String() string {
	sb := strings.Builder{}
	sb.WriteString("{")
	i := 0
	space := " "
	for s := range set {
		if i == len(set)-1 {
			space = ""
		}
		sb.WriteString(s + space)
		i++
	}
	sb.WriteString("}")
	return sb.String()
}

func (set SymbolSet) add(ss ...Symbol) SymbolSet {
	for _, s := range ss {
		set[s] = struct{}{}
	}
	return set
}
func (set SymbolSet) remove(s Symbol) SymbolSet {
	delete(set, s)
	return set
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

func (set SymbolSet) toList() []Symbol {
	ret := make([]Symbol, 0, len(set))
	for s := range set {
		ret = append(ret, s)
	}
	return ret
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
