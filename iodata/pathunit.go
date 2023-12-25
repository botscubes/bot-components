package iodata

type PathUnitType = int

const (
	Array  = 0
	Struct = 1
	Value  = 2
)

type PathUnit struct {
	Type    PathUnitType
	Name    string
	Index   int
	Subpath []PathUnit
}

func getPathFromString(path string) ([]PathUnit, error) {
	for char := range []rune(path) {
		println(char)
	}

	return nil, nil
}
