package grammar

type PredictTable struct {
	terminals   []Symbol
	nonterminal []Symbol

	contents []*Production
}

func newPredictTable(nonterminals []Symbol, terminals []Symbol) *PredictTable {
	terminals = SortSymbol(terminals)
	return &PredictTable{
		terminals:   terminals,
		nonterminal: nonterminals,
		contents:    make([]*Production, len(terminals)*len(nonterminals)),
	}
}

func (a *PredictTable) add(nonterminal Symbol, terminal Symbol, content *Production) {
	row := a.addRow(nonterminal)
	col := a.addCol(terminal)
	a.addContent(row, col, content)
}

func (a *PredictTable) addRow(nonterminal Symbol) (index int) {
	if index = indexOfSymbolList(nonterminal, a.nonterminal); index == NotFound {
		a.nonterminal = append(a.nonterminal, nonterminal)
		index = len(a.nonterminal) - 1
	}
	return
}

func (a *PredictTable) addCol(terminal Symbol) (index int) {
	if index = indexOfSymbolList(terminal, a.terminals); index == NotFound {
		a.terminals = append(a.terminals, terminal)
		index = len(a.terminals) - 1
	}
	return
}

func (a *PredictTable) addContent(row int, col int, content *Production) {
	// rowCount := len(a.nonterminal)
	colCount := len(a.terminals)
	a.contents[row*colCount+col] = content
}
