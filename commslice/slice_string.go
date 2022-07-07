package commslice

import (
	"github.com/reaperhero/instrument/mathutil"
	"strings"
)

// StringsRemove a value form a string slice
func StringsRemove(source []string, s string) []string {
	ns := make([]string, 0, len(source))
	for _, v := range source {
		if v != s {
			ns = append(ns, v)
		}
	}

	return ns
}

// StringsGetRandom get random element from an array/slice
func StringsGetRandom(source []string) string {
	i := mathutil.RandomInt(0, len(source)-1)
	return source[i]
}

// StringsTrim trim string slice item.
// ss item is trim with cut slice in all item
func StringsTrim(ss []string, cutSet ...string) (ns []string) {
	cutSetLn := len(cutSet)
	hasCutSet := cutSetLn > 0 && cutSet[0] != ""

	var trimSet string
	if hasCutSet {
		trimSet = cutSet[0]
	}
	if cutSetLn > 1 {
		trimSet = strings.Join(cutSet, "")
	}

	for _, str := range ss {
		if hasCutSet {
			ns = append(ns, strings.Trim(str, trimSet))
		} else {
			ns = append(ns, strings.TrimSpace(str))
		}
	}
	return
}

// RemoveDuplicateItem
// delete repeat item and result maybe contain " " string
func RemoveDuplicateItem(source []string) []string {
	var list []string

	keys := make(map[string]bool)
	for _, entry := range source {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// RemoveDuplicateItemWithOutNull
// delete repeat item and result not contain " " string
func RemoveDuplicateItemWithOutNull(source []string) []string {
	var list []string

	keys := make(map[string]bool)
	for _, entry := range source {
		entry = strings.Replace(entry, " ", "", -1)
		if len(entry) == 0 {
			continue
		}
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// StringSliceIntersection
// union collection
func StringSliceIntersection(slice1 []string, slice2 []string) []string {
	m := make(map[string]struct{})
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v] = struct{}{}
	}
	for _, v := range slice2 {
		if _, ok := m[v]; ok {
			nn = append(nn, v)
		}
	}
	return nn
}

// StringSliceDifference
// slice1 - slice2
func StringSliceDifference(slice1 []string, slice2 []string) []string {
	keyMap := make(map[string]struct{})
	for _, s := range slice2 {
		keyMap[s] = struct{}{}
	}
	nn := make([]string, 0)
	for searchKey, _ := range keyMap {
		for k := 0; k < len(slice1); k++ {
			if slice1[k] != searchKey {
				slice1 = append(slice1[:k], slice1[k+1:]...)
				k--
				continue
			}
			nn = append(nn, slice1[k])
		}
	}
	return nn
}

// DiffNoOrderStringSlice
// diff no ordered slice
func DiffNoOrderStringSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	diff := make(map[string]int, len(x))
	for _, _x := range x {
		diff[_x]++
	}
	for _, _y := range y {
		if _, ok := diff[_y]; !ok {
			return false
		}
		diff[_y] -= 1
		if diff[_y] == 0 {
			delete(diff, _y)
		}
	}
	if len(diff) == 0 {
		return true
	}
	return false
}

// SliceContrast  x slice in y slice
func SliceContrast(x, y []string) bool {
	if len(y) == 0 {
		return false
	}
	if len(x) == 0 {
		return true
	}
	diff := make(map[string]int)
	for _, _y := range y {
		diff[_y] = 1
	}
	for _, v := range x {
		if _, ok := diff[v]; !ok {
			return false
		}
	}
	return true
}

// StringsReverse string slice [first second three] -> [three second first]
func StringsReverse(source []string) {
	ln := len(source)

	for i := 0; i < ln/2; i++ {
		li := ln - i - 1
		source[i], source[li] = source[li], source[i]
	}
}

// StringsHas check the []string contains the given element
func StringsHas(ss []string, val string) bool {
	for _, ele := range ss {
		if ele == val {
			return true
		}
	}
	return false
}
