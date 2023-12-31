package context

import (
	"reflect"
)

type Value struct {
	data any
}

func (v *Value) ToString() (string, error) {
	rv := reflect.ValueOf(v.data)
	switch rv.Type().Kind() {
	case reflect.String:
		return rv.String(), nil
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
