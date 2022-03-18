package kv

type jsonBoolean struct {
	*jsonUndefined
	value bool
}

func NewBoolean(v bool) Element {
	return &jsonBoolean{
		value: v,
	}
}
func (elt *jsonBoolean) GetType() int {
	return BooleanType
}

func (elt *jsonBoolean) GetBoolean() bool {
	return elt.value
}

func (elt *jsonBoolean) TryGetBoolean() (bool, bool) {
	return elt.value, true
}

func (elt *jsonBoolean) RealValue() interface{} {
	return elt.value
}
func (elt *jsonBoolean) IsSimple() bool {
	return true
}

func (elt *jsonBoolean) ToJson() string {
	if elt.value {
		return "true"
	}
	return "false"
}

func (elt *jsonBoolean) StringValue() string {
	if elt.value {
		return "true"
	}
	return "false"
}

func (elt *jsonBoolean) IsValid() bool {
	return true
}
func (elt *jsonBoolean) IsEqual(e Element) bool {
	if e.GetType() != BooleanType {
		return false
	}
	return elt.value == e.GetBoolean()
}
