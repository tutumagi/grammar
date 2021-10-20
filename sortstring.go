package grammar

import "sort"

type sortSymbol []Symbol

func (s sortSymbol) Less(i, j int) bool {
	if s[i] == rightEndMarkerS {
		return false
	}
	if s[j] == rightEndMarkerS {
		return true
	}
	return s[i] < s[j]
}

func (s sortSymbol) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortSymbol) Len() int {
	return len(s)
}

func SortSymbol(s []Symbol) []Symbol {
	dst := make([]Symbol, len(s))
	copy(dst, s)
	sort.Sort(sortSymbol(dst))
	return dst
}
