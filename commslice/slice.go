package commslice

import "strings"

// RemoveSliceStringItem
// delete item in source
func RemoveSliceStringItem(source []string, item string) []string {
	for k := 0; k < len(source); k++ {
		if source[k] == item {
			source = append(source[:k], source[k+1:]...)
			k = k - 1
		}
	}
	return source
}

// RemoveDuplicateStrElement
// delete repeat item
// contain " " string
func RemoveDuplicateStrElement(arr []string) []string {
	var list []string

	keys := make(map[string]bool)
	for _, entry := range arr {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// RemoveDuplicateStrElementWithOutNull
// delete repeat item
// not contain " " string
func RemoveDuplicateStrElementWithOutNull(arr []string) []string {
	var list []string

	keys := make(map[string]bool)
	for _, entry := range arr {
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

// SliceIntersectionString
// diff sinter union
func SliceIntersectionString(slice1 []string, slice2 []string) []string {
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

// SliceContrast
// x item in y slice
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

