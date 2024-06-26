package context

import (
	"reflect"
	"strconv"
)

type Value struct {
	data any
}

func (v *Value) GetRawValue() any {
	return v.data
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
	case reflect.Int, reflect.Int64, reflect.Int32:
		s := strconv.FormatInt(rv.Int(), 10)
		return s, nil
	}

	if b, ok := v.data.([]byte); ok {
		return string(b), nil
	}

	return "", NewErrTypeAssertion(reflect.TypeOf(v.data).String(), "string")
}
func (v *Value) ToInt64() (int64, error) {
	rv := reflect.ValueOf(v.data)
	switch rv.Type().Kind() {
	case reflect.String:
		i, err := strconv.ParseInt(rv.String(), 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	case reflect.Float64:
		return int64(rv.Float()), nil
	}

	return 0, NewErrTypeAssertion(reflect.TypeOf(v.data).String(), "int64")
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

func (v *Value) ToBytes() ([]byte, error) {
	val, ok := v.data.([]byte)
	if !ok {
		return nil, NewErrTypeAssertion(reflect.TypeOf(v.data).String(), "[]byte")
	}
	return val, nil
}
