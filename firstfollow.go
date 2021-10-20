package grammar

type G struct {
	source string

	start Symbol // start rule

	productions map[Symbol][]*Production

	firstSet     map[Symbol]SymbolSet
	followSet    map[Symbol]SymbolSet
	predictTable *PredictTable

	firstSetDetail map[Symbol]*FirstSetDetail
}

func NewGrammar(source string) *G {
	g := &G{
		source: source,
	}
	g.start, g.productions = makeProductions(source)

	g.firstSetDetail = make(map[string]*FirstSetDetail)

	return g
}

func (g *G) makeFirstSet() {
	g.firstSet = make(map[Symbol]SymbolSet)

	for sym := range g.productions {
		g.firstSet[sym] = g.nonterminalFirstSet(sym)
	}
}

func (g *G) nonterminalFirstSet(sym Symbol) SymbolSet {
	if set, ok := g.firstSet[sym]; ok {
		return set
	}
	set := make(SymbolSet)
	for _, production := range g.productions[sym] {
		theFirstSet := g.rhsFirstSet(production.rhs...)
		if _, ok := g.firstSetDetail[sym]; !ok {
			g.firstSetDetail[sym] = newFirstSetDetail(sym)
		}
		g.firstSetDetail[sym].addDetail(production, theFirstSet)

		set.union(theFirstSet)
	}
	return set
}

func (g *G) rhsFirstSet(symbols ...Symbol) SymbolSet {
	set := make(SymbolSet)
	head := symbols[0]
	rest := symbols[1:]
	if isNonTerminal(head) {
		// 如果是 nonterminal
		headSet := g.nonterminalFirstSet(head)
		set.union(headSet)
		if len(rest) > 0 {
			// FirstSet(head) 包含 ε，则结果为 FirstSet(head) - ε + FirstSet(rest)
			if headSet.contain(epsilonS) {
				set.remove(epsilonS)
				set.union(g.rhsFirstSet(rest...))
			}
			// FirstSet(head) 不包含 ε，则结果就是 FirstSet(head)
		}
	} else {
		// 如果是 ε 或者 terminal，则直接加入到 FirstSet
		set.add(head)
	}
	return set
}

func (g *G) makeFollowSet() {
	g.followSet = make(map[Symbol]SymbolSet)

	for sym := range g.productions {
		if isNonTerminal(sym) {
			g.followSet[sym] = g.symbolFollowSet(sym)
		}
	}
}

func (g *G) symbolFollowSet(sym Symbol) SymbolSet {
	if set, ok := g.followSet[sym]; ok {
		return set
	}
	// followB 方便注释
	followB := make(SymbolSet)
	// 规则1: 如果B是产生式开始的地方，则将input right end marker(这里是$) 放入 Follow(B)
	if sym == g.start {
		followB.add(rightEndMarkerS)
	}
	for _, productions := range g.productions {
		for _, production := range productions {
			ret := indexOfSymbolList(sym, production.rhs)
			if ret == NotFound {
				continue
			}
			beta := production.rhs[ret+1:]
			if len(beta) > 0 {
				// 规则2
				// 如果产生式如下形式：X -> ɑBβ，则先计算 First(β)
				betaFirst := g.rhsFirstSet(beta...)
				// 如果First(β)不包含ε 则Follow(B) = First(β)
				followB.union(betaFirst)
				if betaFirst.contain(epsilonS) {
					// 如果First(β)包含ε 则Follow(B) = First(β) - ε + Follow(X)
					followB.remove(epsilonS)
					followB.union(g.symbolFollowSet(production.lhs))
				}
			} else {
				// fix infinite recursion
				if production.lhs != sym {
					// 规则3
					// 如果产生式如下形式：X -> ɑB
					// 则 Follow(X) 包含在 Follow(B)当中
					followB.union(g.symbolFollowSet(production.lhs))
				}
			}
		}
	}

	return followB
}

func (g *G) makePredict() {
	g.predictTable = newPredictTable(g.terminalsAndNonterminals())
	for nonterminal, fsd := range g.firstSetDetail {
		for production, firstSet := range fsd.detail {
			for terminal := range firstSet {
				g.predictTable.add(nonterminal, terminal, production)
			}
		}
	}
}

func (g *G) terminalsAndNonterminals() (terminals []Symbol, nonterminals []Symbol) {
	terminalSet := newSymbolSet()
	nonterminalSet := newSymbolSet()

	add := func(sym ...Symbol) {
		for _, s := range sym {
			if isTerminal(s) {
				terminalSet.add(s)
			} else if isNonTerminal(s) {
				nonterminalSet.add(s)
			}
		}
	}

	for sym, productions := range g.productions {
		add(sym)
		for _, production := range productions {
			add(sym, production.lhs)
			add(production.rhs...)
		}
	}
	return terminalSet.toList(), nonterminalSet.toList()
}
