package commonprefix

import (
	"sort"
)

func CommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sort.Strings(strs)
	var first = strs[0]
	var last = strs[len(strs)-1]
	var count = 0
	for count < len(first) {
		if first[count] == last[count] {
			count++
		} else {
			break
		}
	}
	if count == 0 {
		return ""
	} else {
		return first[0:count]
	}
}
