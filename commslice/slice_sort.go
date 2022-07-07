package commslice

import "sort"

type LessFun func(x, y interface{}) bool
type sortStruct struct {
	item []interface{}
	less func(x, y interface{}) bool
}

func (x sortStruct) Len() int           { return len(x.item) }
func (x sortStruct) Less(i, j int) bool { return x.less(x.item[i], x.item[j]) }
func (x sortStruct) Swap(i, j int)      { x.item[i], x.item[j] = x.item[j], x.item[i] }

func NewCustomSort(item []interface{}, fn LessFun) []interface{} {
	sortSource := sortStruct{
		item: item,
		less: fn,
	}
	sort.Sort(sortSource)
	return sortSource.item
}

var (
	OptionSortInt LessFun = func(x, y interface{}) bool {
		xInt, ok := x.(int)
		if !ok {
			panic("x not int")
		}
		yInt, ok := y.(int)
		if !ok {
			panic("y not int")
		}
		return xInt > yInt
	}
)
