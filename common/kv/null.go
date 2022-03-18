package kv

type jsonNull struct {
	*jsonUndefined
}

var Null Element = (*jsonNull)(nil)

func NewNull() Element {
	return Null
}
func (elt *jsonNull) GetType() int {
	return NullType
}
func (elt *jsonNull) GetNull() interface{} {
	return nil
}

func (elt *jsonNull) TryGetNull() (interface{}, bool) {
	return nil, true
}

func (elt *jsonNull) RealValue() interface{} {
	return nil
}
func (elt *jsonNull) IsSimple() bool {
	return true
}
func (elt *jsonNull) ToJson() string {
	return "null"
}

func (elt *jsonNull) StringValue() string {
	return "null"
}

func (elt *jsonNull) IsValid() bool {
	return true
}
func (elt *jsonNull) IsEqual(e Element) bool {
	return e.GetType() == NullType
}
