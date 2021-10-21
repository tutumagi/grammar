package grammar

import (
	"fmt"
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func PrintFirstSet(writer io.Writer, firstSet map[Symbol]SymbolSet) {
	printMapSymbolSet(writer, firstSet, func(lhs Symbol) string {
		return "FIRST(" + lhs + ")"
	})
}

func PrintFollowSet(writer io.Writer, followSet map[Symbol]SymbolSet) {
	printMapSymbolSet(writer, followSet, func(lhs Symbol) string {
		return "FOLLOW(" + lhs + ")"
	})
}

func PrintPredictTable(writer io.Writer, predictTable *PredictTable) {
	t := table.NewWriter()
	t.SetOutputMirror(writer)
	// avoid render header in upper case
	t.SetStyle(table.StyleDefault)
	t.Style().Format.Header = text.FormatDefault

	rowCount := len(predictTable.nonterminal)
	colCount := len(predictTable.terminals)
	// draw title
	header := table.Row{"#"}
	for _, terminal := range predictTable.terminals {
		header = append(header, terminal)
	}
	t.AppendHeader(header)

	for r := 0; r < rowCount; r++ {
		row := table.Row{predictTable.nonterminal[r]}
		for l := 0; l < colCount; l++ {
			row = append(row, predictTable.contents[r*colCount+l].String())
		}
		t.AppendRow(row)
		t.AppendSeparator()
	}
	t.Render()
}

func printMapSymbolSet(writer io.Writer, symbolSet map[Symbol]SymbolSet, wrapperLHS func(lhs Symbol) string) {
	if wrapperLHS == nil {
		wrapperLHS = func(lhs Symbol) string {
			return lhs
		}
	}
	for lhs, set := range symbolSet {
		fmt.Fprintf(writer, "%s = %s\n", wrapperLHS(lhs), set.String())
	}
	fmt.Fprintln(writer)
}
