package paginate

import (
	"fmt"
	"reflect"
)

func Paginate(item interface{}, page, size int) (int, error) {
	if page <= 0 || size < 0 {
		return 0, fmt.Errorf("page size not legal")
	}
	offset := (page - 1) * size
	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Slice {
		return 0, fmt.Errorf("item not slice")
	}

	if v.Len() == 0 {
		return 0, fmt.Errorf("item len is zero")
	}
	count := v.Len()

	switch v.Index(0).Kind() {
	case reflect.String:
		var list []string
		for i := 0; i < count; i++ {
			list = append(list, v.Index(i).Interface().(string))
		}
		if count < offset {
			return count, nil
		}
		if count-offset <= size {
			list = list[offset:count]
		}
		if count-offset > size {
			list = list[offset : offset+size]
		}
		reflect.ValueOf(item).Elem().Set(reflect.ValueOf(list))
	case reflect.Interface:
		var list []interface{}
		for i := 0; i < count; i++ {
			list = append(list, v.Index(i).Interface())
		}
		if count < offset {
			return count, nil
		}
		if count-offset <= size {
			list = list[offset:count]
		}
		if count-offset > size {
			list = list[offset : offset+size]
		}
		reflect.ValueOf(item).Elem().Set(reflect.ValueOf(list))
	case reflect.Int:
		var list []int
		for i := 0; i < count; i++ {
			list = append(list, v.Index(i).Interface().(int))
		}
		if count < offset {
			return count, nil
		}
		if count-offset <= size {
			list = list[offset:count]
		}
		if count-offset > size {
			list = list[offset : offset+size]
		}
		reflect.ValueOf(item).Elem().Set(reflect.ValueOf(list))

	}

	return count, nil
}
