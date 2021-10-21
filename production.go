package grammar

import (
	"fmt"
	"strings"
)

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

func (p *Production) String() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("%s -> %v", p.lhs, p.rhs)
}

func makeProductions(source string) (start Symbol, productions map[Symbol][]*Production) {
	lines := strings.Split(source, "\n")
	productions = make(map[Symbol][]*Production, len(lines))
	for _, line := range lines {
		sym, lineRules := makeLineProduction(line)
		if start == "" {
			start = sym // start rule
		}
		if len(lineRules) > 0 {
			productions[sym] = append(productions[sym], lineRules...)
		}
	}
	return
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
