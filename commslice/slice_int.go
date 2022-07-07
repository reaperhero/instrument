package commslice

// IntsHas check the []int contains the given value
func IntsHas(ints []int, val int) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

// Int64sHas check the []int64 contains the given value
func Int64sHas(ints []int64, val int64) bool {
	for _, ele := range ints {
		if ele == val {
			return true
		}
	}
	return false
}

