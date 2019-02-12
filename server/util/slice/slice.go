package slice

import (
	"fmt"
	"reflect"
)

func Contains(slice interface{}, v interface{}) bool {
	interfaces, err := sliceToInterfaces(slice)
	if err != nil {
		return false
	}

	for _, elm := range interfaces {
		if elm == v {
			return true
		}
	}
	return false
}

func AnyMatch(slice interface{}, matcher interface{}) bool {
	values, err := sliceToValues(slice)
	if err != nil {
		return false
	}

	m := reflect.ValueOf(matcher)
	if m.Kind() != reflect.Func {
		return false
	}

	for _, value := range values {
		ret := m.Call([]reflect.Value{value})
		if len(ret) == 1 && ret[0].Kind() == reflect.Bool && ret[0].Bool() {
			return true
		}
	}

	return false
}

func sliceToInterfaces(slice interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("slice.sliceToInterfeces does not have slice.: %v", v.Kind())
	}

	length := v.Len()
	s := make([]interface{}, length)
	for i := 0; i < length; i++ {
		s[i] = v.Index(i).Interface()
	}

	return s, nil
}

func sliceToValues(slice interface{}) ([]reflect.Value, error) {
	value := reflect.ValueOf(slice)
	if value.Kind() != reflect.Slice {
		return nil, fmt.Errorf("slice.sliceToValues does not have slice.: %v", value.Kind())
	}

	length := value.Len()
	values := make([]reflect.Value, length)
	for i := 0; i < length; i++ {
		values[i] = value.Index(i)
	}
	return values, nil
}
