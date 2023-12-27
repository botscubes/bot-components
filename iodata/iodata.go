package iodata

import (
	"encoding/json"
)

type IOData struct {
	data map[string]any
}

func NewIODataFromJSON(jsonData []byte) (*IOData, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}
	return &IOData{
		data: data,
	}, nil
}
func (d *IOData) getValue(path string) (any, error) {
	return d.getValueUsingPath(NewPathUnitIterator(path))
}

func (d *IOData) getValueUsingPath(
	iter *PathUnitIterator,
) (any, error) {
	var value any = d.data

	for iter.HasNext() {
		val, err := iter.Next()
		if err != nil {
			return nil, err
		}

		if val.Type == Array {
			arr, ok := value.([]any)
			if !ok {
				return nil, ErrTypeMismatch
			}
			arrSize := len(arr)
			if val.Index >= arrSize || arrSize < 0 {
				return nil, ErrIndexOutOfRange
			}
			value = arr[val.Index]

		} else if val.Type == Struct {
			m, ok := value.(map[string]any)
			if !ok {
				return nil, ErrTypeMismatch
			}

			if value, ok = m[val.Propery]; !ok {
				return nil, ErrKeyIsNotInMap
			}

		} else {

			return nil, ErrTypeNotFound

		}
	}
	return value, nil
}
