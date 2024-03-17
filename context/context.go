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
func NewContext() *Context {
	return &Context{
		data: make(map[string]any),
	}
}

func (ctx *Context) ToJSON() ([]byte, error) {
	return json.Marshal(ctx.data)
}
func (ctx *Context) GetValue(path string) (*Value, error) {
	return ctx.getValueUsingPath(NewPathUnitIterator(path))
}

func (ctx *Context) getValueUsingPath(
	iter *PathUnitIterator,
) (*Value, error) {
	var value any = ctx.data
	objectOrArrayName := "Context"

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
				val, err := ctx.getValueUsingPath(val.Subpath)
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
				val, err := ctx.getValueUsingPath(val.Subpath)
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

func (ctx *Context) SetValue(path string, value *any) error {
	iter := NewPathUnitIterator(path)
	var data any = ctx.data
	objectOrArrayName := "Context"

	for iter.HasNext() {
		val, err := iter.Next()
		if err != nil {
			return err
		}

		if val.Type == Array {
			arr, ok := data.([]any)
			if !ok {
				return NewErrTypeAssertion(reflect.TypeOf(data).String(), "array")
			}
			if val.Subpath != nil {
				val, err := ctx.getValueUsingPath(val.Subpath)
				if err != nil {
					return err
				}
				idx, err := val.ToInt()
				if err != nil {
					return err
				}
				arrSize := len(arr)
				if idx >= arrSize || arrSize < 0 {
					return NewErrIndexOutOfRange(objectOrArrayName, idx, 0, arrSize)
				}
				data = arr[idx]

			} else {
				arrSize := len(arr)
				if val.Index >= arrSize || arrSize < 0 {
					return NewErrIndexOutOfRange(objectOrArrayName, val.Index, 0, arrSize)
				}
				if !iter.HasNext() {
					arr[val.Index] = *value
				} else {
					data = arr[val.Index]
				}
			}

		} else if val.Type == Object {
			m, ok := data.(map[string]any)
			if !ok {
				return NewErrTypeAssertion(reflect.TypeOf(data).String(), "map")
			}
			if val.Subpath != nil {
				val, err := ctx.getValueUsingPath(val.Subpath)
				if err != nil {
					return err
				}
				property, err := val.ToString()
				if err != nil {
					return err
				}
				if data, ok = m[property]; !ok {
					return NewErrNoPropertyInObject(objectOrArrayName, property)
				}
				objectOrArrayName = property
			} else {
				if !iter.HasNext() {
					m[val.Propery] = *value
				} else {
					if data, ok = m[val.Propery]; !ok {
						return NewErrNoPropertyInObject(objectOrArrayName, val.Propery)
					}
					objectOrArrayName = val.Propery
				}
			}
		} else {

			return ErrTypeNotFound

		}
	}
	return nil

}
