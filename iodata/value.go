package iodata

import "reflect"

type Value struct {
	data any
}

func (v *Value) ToString() (string, error) {
	rv := reflect.ValueOf(v.data)
	switch rv.Type().Kind() {
	case reflect.String:
		return rv.String(), nil
	}
	return "", ErrTypeMismatch
}
func (v *Value) ToInt64() (int64, error) {
	val, ok := v.data.(float64)
	if !ok {
		return 0, ErrTypeMismatch
	}
	return int64(val), nil
}
func (v *Value) ToInt() (int, error) {
	val, ok := v.data.(float64)
	if !ok {
		return 0, ErrTypeMismatch
	}
	return int(val), nil
}
