package kv

import "strings"

type jsonArray struct {
	*jsonUndefined
	value []Element
}

func NewArray() Element {
	return &jsonArray{
		value: make([]Element, 0),
	}
}

func (elt *jsonArray) GetType() int {
	return ArrayType
}

func (elt *jsonArray) GetElement(idx int) Element {
	l := len(elt.value)
	if idx >= 0 && idx < l {
		return elt.value[idx]
	}
	return Undefined
}

func (elt *jsonArray) TryGetElement(idx int) (Element, bool) {
	l := len(elt.value)
	if idx >= 0 && idx < l {
		return elt.value[idx], true
	}
	return Undefined, false
}

func (elt *jsonArray) ArrayEnumerator() []Element {
	return elt.value
}

func (elt *jsonArray) Push(item Element) Element {
	elt.value = append(elt.value, item)
	return elt
}
func (elt *jsonArray) ToJson() string {
	childStrs := make([]string, 0)
	for _, i := range elt.ArrayEnumerator() {
		childStrs = append(childStrs, i.ToJson())
	}
	return "[" + strings.Join(childStrs, ",") + "]"
}

func (elt *jsonArray) StringValue() string {
	return elt.ToJson()
}

func (elt *jsonArray) Select(path string) Element {
	return Select(elt, path)
}
func (elt *jsonArray) TrySelect(path string) (Element, bool) {
	return TrySelect(elt, path)
}

func (elt *jsonArray) IsValid() bool {
	return true
}
