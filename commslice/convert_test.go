package commslice

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToInt64s(t *testing.T) {
	is := assert.New(t)

	ints, err := ConvertInterfaceToInt64s([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int64{1, 2}", fmt.Sprintf("%#v", ints))
}

func TestConvertStringsToInts(t *testing.T) {
	ints, err := ConvertStringsToInts([]string{"1", "2", "3"})
	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, []int{1, 2, 3}, ints)
}

func TestToStrings(t *testing.T) {
	is := assert.New(t)

	ss, err := ConvertInterfaceToStrings([]int{1, 2})
	is.Nil(err)
	is.Equal(`[]string{"1", "2"}`, fmt.Sprintf("%#v", ss))
}

func TestConvertSliceToStrings(t *testing.T) {
	item := []interface{}{5, "2", "3"}
	result := ConvertSliceToStrings(item)
	assert.Equal(t, []string{"5", "2", "3"}, result)
}
