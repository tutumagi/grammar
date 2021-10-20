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

func newProduction(lhs Symbol) *Production {
	return &Production{
		lhs: lhs,
		rhs: make([]Symbol, 0),
	}
}

func (p *Production) RHS(sym ...Symbol) *Production {
	p.rhs = sym
	return p
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

				productions = append(productions, newProduction(lhs).RHS(symbols...))
			}
		}
	}
	return
}

func (g *G) makeFirstSet() {
	g.firstSet = make(FirstSet)

	for sym := range g.productions {
		g.firstSet[sym] = g.firstSetByProductions(sym)
	}
}

func (g *G) firstSetByProductions(sym Symbol) SymbolSet {
	set := make(SymbolSet)
	for _, production := range g.productions[sym] {
		set.union(g.rhsFirstSet(production.rhs...))
	}
	return set
}

func (g *G) rhsFirstSet(symbols ...Symbol) (set SymbolSet) {
	set = make(SymbolSet)
	head := symbols[0]
	rest := symbols[1:]
	if isNonTerminal(head) {
		headSet := g.firstSetByProductions(head)
		set.union(headSet)
		if len(rest) > 0 {
			if headSet.contain(epsilonS) {
				set.union(g.rhsFirstSet(rest...))
			}
		}
	} else {
		set.add(head)
	}
	return
}
