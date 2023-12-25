package iodata

type IOData struct {
	data map[string]any
}

func IODataFromJSON(json []byte) IOData {
	return IOData{
		data: map[string]any{},
	}
}
func (d *IOData) getValue(path string) (any, error) {
	return d.getValueUsingPath()
}

func (d *IOData) getValueUsingPath(
	path ...PathUnit,
) (any, error) {
	var value any = d.data
	var currentType = Struct
	for _, val := range path {
		if currentType == Array {
			arr, ok := value.([]any)
			if !ok {
				return nil, ErrTypeMismatch
			}
			arrSize := len(arr)
			if val.Index >= arrSize || arrSize < 0 {
				return nil, ErrIndexOutOfRange
			}
			value = arr[val.Index]

		} else if currentType == Struct {
			m, ok := value.(map[string]any)
			if !ok {
				return nil, ErrTypeMismatch
			}

			if value, ok = m[val.Name]; !ok {
				return nil, ErrKeyIsNotInMap
			}

		} else {

			return nil, ErrTypeNotFound

		}
		currentType = val.Type
	}
	return value, nil
}
