package iodata

import (
	"fmt"
	"strconv"
)

type PathUnitType = int

const (
	Array  = 0
	Struct = 1
)

type PathUnit struct {
	Type    PathUnitType
	Propery string
	Index   int
	Subpath []*PathUnit
}

type PathUnitIterator struct {
	path       string
	curr_index int
}

func NewPathUnitIterator(path string) *PathUnitIterator {
	return &PathUnitIterator{
		path:       path,
		curr_index: 0,
	}
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func isLetter(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

func (it *PathUnitIterator) getExplicitArrayIndex() (*PathUnit, error) {
	str := it.path
	l := it.curr_index
	for {
		it.curr_index++
		if it.curr_index < len(str) {
			ch := str[it.curr_index]
			if ch == ']' {

				idx, err := strconv.Atoi(str[l:it.curr_index])
				if err != nil {
					return nil, err
				}
				it.curr_index++
				return &PathUnit{
					Type:    Array,
					Propery: "",
					Index:   idx,
					Subpath: nil,
				}, nil

			} else if !isDigit(ch) {
				return nil, ErrWrongIndex
			}
		} else {
			return nil, ErrNoClosingSquareBracket
		}
	}
}

func (it *PathUnitIterator) getArrayIndex() (*PathUnit, error) {
	str := it.path
	if it.curr_index < len(str) {
		ch := str[it.curr_index]
		if isDigit(ch) {
			return it.getExplicitArrayIndex()
			//						} else if isLetter(ch) || ch == '_' {
			//							l := it.curr_index
			//							r := it.curr_index
			//							openBracketCount := 0
			//							for {
			//								r++
			//								if r < len(str) {
			//									ch = str[r]
			//									if ch == ']' {

			//									}
			//									return &PathUnit{
			//										Type:    Array,
			//										Name:    str[li:it.curr_index],
			//										Index:   0,
			//										Subpath: nil,
			//									}, nil
			//								} else {
			//									return nil, ErrNoClosingSquareBracket
			//								}

			//							}
		} else {
			return nil, ErrVariableName
		}
	} else {
		return nil, ErrNoClosingSquareBracket
	}

}

func (it *PathUnitIterator) getStructProperty() (*PathUnit, error) {
	if it.HasNext() {
		str := it.path
		ch := str[it.curr_index]
		li := it.curr_index
		if isLetter(ch) || ch == '_' {
			for {
				it.curr_index++
				if it.curr_index < len(str) {
					ch = str[it.curr_index]
					if ch == '.' || ch == '[' {
						return &PathUnit{
							Type:    Struct,
							Propery: str[li:it.curr_index],
							Index:   0,
							Subpath: nil,
						}, nil
					} else if !isDigit(ch) && !isLetter(ch) && !(ch == '_') {
						return nil, ErrVariableName
					}
				} else {
					return &PathUnit{
						Type:    Struct,
						Propery: str[li:it.curr_index],
						Index:   0,
						Subpath: nil,
					}, nil
				}
			}
		} else {
			return nil, ErrVariableNameBeginning
		}
	} else {

		return nil, ErrVariableNameNotSpecified
	}
}

func (it *PathUnitIterator) Next() (*PathUnit, error) {
	if it.HasNext() {
		str := it.path
		li := it.curr_index
		ch := str[li]
		if (li == 0) && (isLetter(ch) || ch == '_') {
			return it.getStructProperty()
		} else if ch == '.' {
			it.curr_index++
			return it.getStructProperty()
		} else if ch == '[' {
			it.curr_index++
			return it.getArrayIndex()
		} else {

			return nil, fmt.Errorf("%c", ch) //ErrUnknownCharacter
		}
	}
	return nil, nil
}
func (it *PathUnitIterator) HasNext() bool {
	if it.curr_index < len(it.path) {
		return true
	}
	return false
}
