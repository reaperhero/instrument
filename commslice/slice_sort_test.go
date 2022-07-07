package commslice

import (
	"fmt"
	"testing"
)

func TestCustSort(t *testing.T) {
	item := []interface{}{1, 2, 4, 3, 10, 4}
	result := NewCustomSort(item, OptionSortInt)
	fmt.Println(result)
}
