package kv

import "fmt"

type jsonNumber struct {
	*jsonUndefined
	value float64
}

func NewNumber(v float64) Element {
	return &jsonNumber{
		value: v,
	}
}
func (elt *jsonNumber) GetType() int {
	return NumberType
}
func (elt *jsonNumber) GetNumber() float64 {
	return elt.value
}

func (elt *jsonNumber) TryGetNumber() (float64, bool) {
	return elt.value, true
}

func (elt *jsonNumber) RealValue() interface{} {
	return elt.value
}
func (elt *jsonNumber) IsSimple() bool {
	return true
}
func (elt *jsonNumber) ToJson() string {
	return fmt.Sprintf("%g", elt.value)
}

func (elt *jsonNumber) StringValue() string {
	return fmt.Sprintf("%g", elt.value)
}

func (elt *jsonNumber) IsValid() bool {
	return true
}
func (elt *jsonNumber) IsEqual(e Element) bool {
	if e.GetType() != NumberType {
		return false
	}
	return elt.value == e.GetNumber()
}
