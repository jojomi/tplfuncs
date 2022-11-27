package container

import "strings"

func (x *StringList) Map(mapper func(elem string) string) *StringList {
	for i := 0; i < len(x.strings); i++ {
		x.strings[i] = mapper(x.strings[i])
	}
	return x
}

// WrapAll wraps all elements with given strings.
func (x *StringList) WrapAll(before, after string) *StringList {
	return x.Map(func(elem string) string {
		return before + elem + after
	})
}

// ReplaceAll does rewrite all elements, use $0 for the current value.
func (x *StringList) ReplaceAll(newValue string) *StringList {
	return x.Map(func(elem string) string {
		return strings.ReplaceAll(newValue, "$0", elem)
	})
}

// TrimAll does rewrite all elements to strip leading and trailing whitespace.
func (x *StringList) TrimAll() *StringList {
	return x.Map(func(elem string) string {
		return strings.TrimSpace(elem)
	})
}

// IndentSpaceAll does rewrite all elements adding n leading spaces.
func (x *StringList) IndentSpaceAll(count int) *StringList {
	if count < 1 {
		return x
	}

	return x.Map(func(elem string) string {
		return strings.Repeat(" ", count) + elem
	})
}

// IndentTabAll does rewrite all elements adding n leading tabs.
func (x *StringList) IndentTabAll(count int) *StringList {
	if count < 1 {
		return x
	}

	return x.Map(func(elem string) string {
		return strings.Repeat("\t", count) + elem
	})
}

// UnindentAll does rewrite all elements trimming n leading characters.
func (x *StringList) UnindentAll(count int) *StringList {
	if count < 1 {
		return x
	}

	return x.Map(func(elem string) string {
		return elem[count:]
	})
}
