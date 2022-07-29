package helper

import "reflect"

func ToInterfaceSlice(slice interface{}) []interface{} {
	// Source: https://stackoverflow.com/questions/12753805/type-converting-slices-of-interfaces
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func InterfaceToStringSlice(arr interface{}) *[]string {
	if arr != nil {
		return InterfaceSliceToStringSlice(arr.([]interface{}))
	}
	return nil
}

func InterfaceSliceToStringSlice(arr []interface{}) *[]string {
	strArray := make([]string, len(arr))
	for i, v := range arr {
		strArray[i] = v.(string)
	}
	return &strArray
}

func InterfaceToIntSlice(arr interface{}) []int {
	if arr != nil {
		return InterfaceSliceToIntSlice(arr.([]interface{}))
	}
	return nil
}

func InterfaceSliceToIntSlice(arr []interface{}) []int {
	intArray := make([]int, len(arr))
	for i, v := range arr {
		intArray[i] = v.(int)
	}
	return intArray
}
