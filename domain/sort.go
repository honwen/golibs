package domain

import (
	"index/suffixarray"
	"sort"
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func Sort(domains ...[]string) (sorted []string) {
	all := []string{}
	for i := range domains {
		all = append(all, domains[i]...)
	}

	coredomains := []string{}
	idx := suffixarray.New([]byte(strings.Join(all, "")))
	for i := range all {
		offsets := idx.Lookup([]byte("."+all[i]), 1)
		if len(offsets) > 0 {
			coredomains = append(coredomains, all[i])
		}
	}
	for i := range all {
		cnt := 0
		idx = suffixarray.New([]byte(all[i]))
		for ii := range coredomains {
			offsets := idx.Lookup([]byte("."+coredomains[ii]), 1)
			cnt += len(offsets)
		}
		if cnt == 0 {
			sorted = append(sorted, all[i])
		}
	}
	for i := range sorted {
		sorted[i] = reverseString(sorted[i])
	}
	sort.Strings(sorted)
	for i := range sorted {
		sorted[i] = reverseString(sorted[i])
	}
	return
}
