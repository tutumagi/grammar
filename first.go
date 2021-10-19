package grammar

import (
	"strings"
)

type (
	FirstSet  map[Symbol]SymbolSet
	FollowSet map[Symbol]SymbolSet
)

type G struct {
	source string

	productions map[Symbol][]*Production

	firstSet FirstSet
	// followSet FollowSet
}

type Production struct {
	lhs Symbol
	rhs []Symbol
}

func newProduction(lhs Symbol, rhs ...Symbol) *Production {
	return &Production{
		lhs: lhs,
		rhs: rhs,
	}
}

func NewGrammar(source string) *G {
	return &G{
		source: source,
	}
}

func (g *G) makeProductions() map[Symbol][]*Production {
	lines := strings.Split(g.source, "\n")
	g.productions = make(map[Symbol][]*Production, len(lines))
	for _, line := range lines {
		sym, lineRules := g.makeLineProduction(line)
		if len(lineRules) > 0 {
			g.productions[sym] = append(g.productions[sym], lineRules...)
		}
	}
	return g.productions
}

func (g *G) makeLineProduction(line string) (lhs Symbol, productions []*Production) {
	items := strings.Split(line, productionS)
	if len(items) <= 1 {
		return
	}
	prod := &Production{}
	// 拿到这一行的 LHS 和 RHS
	for i := 0; i < len(items); i++ {
		item := strings.TrimSpace(items[i])
		if i == 0 {
			// 拿到 LHS
			lhs = item
			prod.lhs = item
		} else {
			// 拿到 RHS，可能会有多个 production，通过 `alternateS` 符号分割
			mulipleRules := strings.Split(item, alternateS)
			for _, rule := range mulipleRules {
				// 每个 production 的符号集合
				// 每个符号之间用空格隔开，比如 S -> A B C d
				symbols := strings.Split(strings.TrimSpace(rule), " ")

				productions = append(productions, newProduction(lhs, symbols...))
			}
		}
	}
	return
}

func (g *G) makeFirstSet() {
	g.firstSet = make(FirstSet)

	for sym := range g.productions {
		g.firstSet[sym] = g.firstSetBySymbol(sym)
	}
}

func (g *G) firstSetBySymbol(sym Symbol) (set SymbolSet) {
	set = make(SymbolSet)
	defer func() {
		if ss, ok := g.firstSet[sym]; ok {
			ss.union(set)
		} else {
			g.firstSet[sym] = set
		}
	}()
	if isTerminal(sym) {
		set.add(sym)
	} else if isEpsilon(sym) {
		set.add(sym)
	} else {
		var ok bool
		if set, ok = g.firstSet[sym]; ok {
			return
		}
		set = g.firstSetByProductions(sym, g.productions[sym])
		g.firstSet[sym] = set
	}
	return
}

func (g *G) firstSetByProductions(lhs Symbol, productions []*Production) (set SymbolSet) {
	set = make(SymbolSet)
	for _, production := range productions {
		set.union(g.firstSetByProduction(production))
	}

	return set
}

func (g *G) firstSetByProduction(production *Production) (set SymbolSet) {
	set = make(SymbolSet)
	var preSet SymbolSet
	for _, sym := range production.rhs {
		if preSet != nil {
			if preSet.contain(epsilonS) {
				preSet.remove(epsilonS)
			} else {
				break
			}
		}
		curSet := g.firstSetBySymbol(sym)
		set.union(curSet)

		preSet = curSet
	}
	return
}
