package context

import (
	"encoding/json"
	"reflect"
)

type Context struct {
	data map[string]any
}

func NewContextFromJSON(jsonData []byte) (*Context, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return &Context{
		data: data,
	}, nil
}
func (d *Context) GetValue(path string) (*Value, error) {
	return d.getValueUsingPath(NewPathUnitIterator(path))
}

func (d *Context) getValueUsingPath(
	iter *PathUnitIterator,
) (*Value, error) {
	var value any = d.data
	objectOrArrayName := "IOData"

	for iter.HasNext() {
		val, err := iter.Next()
		if err != nil {
			return nil, err
		}

		if val.Type == Array {
			arr, ok := value.([]any)
			if !ok {
				return nil, NewErrTypeAssertion(reflect.TypeOf(value).String(), "array")
			}
			if val.Subpath != nil {
				val, err := d.getValueUsingPath(val.Subpath)
				if err != nil {
					return nil, err
				}
				idx, err := val.ToInt()
				if err != nil {
					return nil, err
				}
				arrSize := len(arr)
				if idx >= arrSize || arrSize < 0 {
					return nil, NewErrIndexOutOfRange(objectOrArrayName, idx, 0, arrSize)
				}
				value = arr[idx]

			} else {
				arrSize := len(arr)
				if val.Index >= arrSize || arrSize < 0 {
					return nil, NewErrIndexOutOfRange(objectOrArrayName, val.Index, 0, arrSize)
				}
				value = arr[val.Index]
			}

		} else if val.Type == Object {
			m, ok := value.(map[string]any)
			if !ok {
				return nil, NewErrTypeAssertion(reflect.TypeOf(value).String(), "map")
			}
			if val.Subpath != nil {
				val, err := d.getValueUsingPath(val.Subpath)
				if err != nil {
					return nil, err
				}
				property, err := val.ToString()
				if err != nil {
					return nil, err
				}
				if value, ok = m[property]; !ok {
					return nil, NewErrNoPropertyInObject(objectOrArrayName, property)
				}
				objectOrArrayName = property
			} else {

				if value, ok = m[val.Propery]; !ok {
					return nil, NewErrNoPropertyInObject(objectOrArrayName, val.Propery)
				}
				objectOrArrayName = val.Propery
			}
		} else {

			return nil, ErrTypeNotFound

		}
	}
	return &Value{data: value}, nil
}
