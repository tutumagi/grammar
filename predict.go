package grammar

import (
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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
	// sb := &strings.Builder{}
	// rowCount := len(a.nonterminal)
	// colCount := len(a.terminals)
	// // draw title
	// sb.WriteString("     ")
	// for _, terminal := range a.terminals {
	// 	sb.WriteString(fmt.Sprintf("%20s", terminal))
	// }
	// sb.WriteString("\n")
	// for r := 0; r < rowCount; r++ {
	// 	sb.WriteString(fmt.Sprintf("%s    ", a.nonterminal[r]))
	// 	for l := 0; l < colCount; l++ {
	// 		sb.WriteString(fmt.Sprintf("%20s", a.contents[r*colCount+l].String()))
	// 	}
	// 	sb.WriteString("\n")
	// }

	// fmt.Print(sb.String())

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	rowCount := len(a.nonterminal)
	colCount := len(a.terminals)
	// draw title
	header := table.Row{"#"}
	for _, terminal := range a.terminals {
		header = append(header, terminal)
	}
	t.AppendHeader(header)

	for r := 0; r < rowCount; r++ {
		row := table.Row{a.nonterminal[r]}
		for l := 0; l < colCount; l++ {
			row = append(row, a.contents[r*colCount+l].String())
		}
		t.AppendRow(row)
		t.AppendSeparator()
	}
	t.Render()

	// fmt.Print(sb.String())

}
