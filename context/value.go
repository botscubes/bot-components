package context

import (
	"reflect"
	"strconv"
)

type Value struct {
	data any
}

func (v *Value) ToString() (string, error) {
	rv := reflect.ValueOf(v.data)
	switch rv.Type().Kind() {
	case reflect.String:
		return rv.String(), nil
	case reflect.Float64:
		s := strconv.FormatFloat(rv.Float(), 'g', -1, 64)
		return s, nil
	case reflect.Bool:
		s := strconv.FormatBool(rv.Bool())
		return s, nil
	}
	return "", NewErrTypeAssertion(reflect.TypeOf(v.data).String(), "string")
}
func (v *Value) ToInt64() (int64, error) {
	val, ok := v.data.(float64)
	if !ok {
		return 0, NewErrTypeAssertion(reflect.TypeOf(v.data).String(), "int")
	}
	return int64(val), nil
}
func (v *Value) ToInt() (int, error) {
	val, ok := v.data.(float64)
	if !ok {
		return 0, NewErrTypeAssertion(reflect.TypeOf(v.data).String(), "float")
	}
	return int(val), nil
}

func (v *Value) ToBool() (bool, error) {
	val, ok := v.data.(bool)
	if !ok {
		return false, NewErrTypeAssertion(reflect.TypeOf(v.data).String(), "bool")
	}
	return val, nil
}
