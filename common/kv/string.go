package kv

import "fmt"

type jsonString struct {
	*jsonUndefined
	value string
}

func NewString(v string) Element {
	return &jsonString{
		value: v,
	}
}
func (elt *jsonString) GetType() int {
	return StringType
}

func (elt *jsonString) GetString() string {
	return elt.value
}
func (elt *jsonString) TryGetString() (string, bool) {
	return elt.value, true
}

func (elt *jsonString) RealValue() interface{} {
	return elt.value
}
func (elt *jsonString) IsSimple() bool {
	return true
}

func (elt *jsonString) ToJson() string {
	return fmt.Sprintf("%q", elt.value)
}

func (elt *jsonString) StringValue() string {
	return elt.value
}

func (elt *jsonString) IsValid() bool {
	return true
}
func (elt *jsonString) IsEqual(e Element) bool {
	if e.GetType() != StringType {
		return false
	}
	return elt.value == e.GetString()
}
