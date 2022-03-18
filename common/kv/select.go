package kv

func getDigit(c rune) int {
	switch c {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	return -1
}

/*
遇到[，判断之前的值是否有效，是key还是index，然后执行TryGetProperty或者TryGetElement，然后inSquareBrackets=true，表示后续的值是整数，以便下一个边界elt.(Array).TryGetElement
遇到]，inSquareBrackets=true才能操作，且index>=0
遇到/，与[执行相同操作，但是不把后续的值标记为index，inSquareBrackets=false
*/
func doSelect(elt Element, path string) (Element, bool) {
	var chars []rune = make([]rune, 0) //记录key值，用于elt.GetProperty
	index := 0                         //记录索引，用于elt.GetElement
	inSquareBrackets := false          //表示在[]之间，此时index=index*10+ number(c)
	hasValue := false                  //用于屏蔽连续的//[]，如 [0/key 和 [0]/key 和[0]key等效
	curr := elt
	exists := false
	for _, c := range path {
		if c == '[' {
			if hasValue {
				if inSquareBrackets {
					curr, exists = curr.TryGetElement(index)
				} else {
					curr, exists = curr.TryGetProperty(string(chars))
				}
				if !exists {
					return Undefined, false
				}
			}
			index = 0
			chars = make([]rune, 0)
			inSquareBrackets = true
			hasValue = false
			continue
		}
		if c == ']' {
			if hasValue && inSquareBrackets {
				curr, exists = curr.TryGetElement(index)
				if !exists {
					return Undefined, false
				}
			}
			index = 0
			chars = make([]rune, 0)
			inSquareBrackets = false
			hasValue = false
			continue
		}
		if c == '/' {
			if hasValue {
				if inSquareBrackets {
					curr, exists = curr.TryGetElement(index)
				} else {
					curr, exists = curr.TryGetProperty(string(chars))
				}
				if !exists {
					return Undefined, false
				}
			}
			index = 0
			chars = make([]rune, 0)
			inSquareBrackets = false
			hasValue = false
			continue
		}

		//值有一个以上的字符，有效
		hasValue = true
		if inSquareBrackets {
			i := getDigit(c)
			if i < 0 {
				return Undefined, false
			}
			index = index*10 + i
			continue
		}
		chars = append(chars, c)
	}

	if hasValue {
		if inSquareBrackets {
			curr, exists = curr.TryGetElement(index)
		} else {
			curr, exists = curr.TryGetProperty(string(chars))
		}
		if !exists {
			return Undefined, false
		}
	}
	return curr, true
}

func TrySelect(elt Element, path string) (Element, bool) {
	return doSelect(elt, path)
}
func Select(elt Element, path string) Element {
	e, y := doSelect(elt, path)
	if y {
		return e
	}
	return Undefined
}
