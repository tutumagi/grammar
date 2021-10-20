package grammar

import "fmt"

type FirstSetDetail struct {
	// 相关的 nonterminal
	symbol Symbol
	// 所有的 production 产生的 FirstSet的集合
	set SymbolSet
	// 每一个 production 都会产生一个 FirstSet
	detail map[*Production]SymbolSet
}

func newFirstSetDetail(sym Symbol) *FirstSetDetail {
	return &FirstSetDetail{
		symbol: sym,
		set:    make(SymbolSet),
		detail: make(map[*Production]SymbolSet),
	}
}

func (a *FirstSetDetail) addDetail(production *Production, firstSet SymbolSet) {
	if CheckLL1Grammar {
		// 如果新计算出的 firstSet 和已有的不相交，则加入
		if a.set.disjoint(firstSet) {
			a.forceAddDetail(production, firstSet)
		} else {
			for prod, fs := range a.detail {
				if !fs.disjoint(firstSet) {
					// TODO: pretty error report
					fmt.Printf("`%s` and `%s` is not disjoint.", prod, production)
				}
			}
		}
	} else {
		a.forceAddDetail(production, firstSet)
	}
}

func (a *FirstSetDetail) forceAddDetail(production *Production, firstSet SymbolSet) {
	a.detail[production] = firstSet
	a.set.union(firstSet)
}
