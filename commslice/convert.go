package commslice

import (
	"github.com/reaperhero/instrument/mathutil"
	"github.com/reaperhero/instrument/strutil"
	"reflect"
	"strconv"
)

// ConvertStringsToInts string slice to int slice
func ConvertStringsToInts(ss []string) ([]int, error) {
	var ints []int
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		ints = append(ints, iVal)
	}
	return ints,nil
}

// ConvertInterfaceToInt64s convert interface{}(allow: array,slice) to []int64
func ConvertInterfaceToInt64s(arr interface{}) (ret []int64, err error) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, ErrInvalidType
	}

	for i := 0; i < rv.Len(); i++ {
		i64, err := mathutil.Int64(rv.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		ret = append(ret, i64)
	}
	return
}

// ConvertInterfaceToStrings convert interface{}(allow: array,slice) to []string
func ConvertInterfaceToStrings(arr interface{}) (ret []string, err error) {
	rv := reflect.ValueOf(arr)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, ErrInvalidType
	}

	for i := 0; i < rv.Len(); i++ {
		str, err := strutil.ToString(rv.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		ret = append(ret, str)
	}
	return ret, nil
}

// ConvertSliceToStrings convert []interface{} to []string
func ConvertSliceToStrings(arr []interface{}) []string {
	ss := make([]string, len(arr))
	for i, v := range arr {
		ss[i] = strutil.MustString(v)
	}
	return ss
}
