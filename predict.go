package grammar

import (
	"os"
)

type PredictTable struct {
	terminals   []Symbol
	nonterminal []Symbol

	contents []*Production
}

func newPredictTable(terminals []Symbol, nonterminals []Symbol) *PredictTable {
	terminals = SortSymbol(terminals)
	if terminals[len(terminals)-1] != rightEndMarkerS {
		terminals = append(terminals, rightEndMarkerS)
	}
	return &PredictTable{
		terminals:   terminals,
		nonterminal: nonterminals,
		contents:    make([]*Production, len(terminals)*len(nonterminals)),
	}
}

func (a *PredictTable) add(nonterminal Symbol, terminal Symbol, content *Production) {
	row := indexOfSymbolList(nonterminal, a.nonterminal)
	col := indexOfSymbolList(terminal, a.terminals)
	a.addContent(row, col, content)
}

func (a *PredictTable) addContent(row int, col int, content *Production) {
	// rowCount := len(a.nonterminal)
	colCount := len(a.terminals)
	a.contents[row*colCount+col] = content
}

func (a *PredictTable) dump() {
	PrintPredictTable(os.Stdout, a)
}
