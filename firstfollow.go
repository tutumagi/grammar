package grammar

import (
	"strings"
)

type (
	FirstSet  = map[Symbol]SymbolSet
	FollowSet = map[Symbol]SymbolSet
)

type G struct {
	source string

	start Symbol // start rule

	productions map[Symbol][]*Production

	firstSet  FirstSet
	followSet FollowSet
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
		sym, lineRules := makeLineProduction(line)
		if g.start == "" {
			g.start = sym // start rule
		}
		if len(lineRules) > 0 {
			g.productions[sym] = append(g.productions[sym], lineRules...)
		}
	}
	return g.productions
}

func makeLineProduction(line string) (lhs Symbol, productions []*Production) {
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
		g.firstSet[sym] = g.nonterminalFirstSet(sym)
	}
}

func (g *G) nonterminalFirstSet(sym Symbol) SymbolSet {
	if set, ok := g.firstSet[sym]; ok {
		return set
	}
	set := make(SymbolSet)
	for _, production := range g.productions[sym] {
		set.union(g.rhsFirstSet(production.rhs...))
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
	g.followSet = make(FollowSet)

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
