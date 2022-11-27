package container

import (
	"regexp"
	"strings"
)

func (x *StringList) Filter(keep func(elem string) (keep bool)) *StringList {
	var newStrings []string

	for _, elem := range x.All() {
		if !keep(elem) {
			continue
		}
		newStrings = append(newStrings, elem)
	}

	x.strings = newStrings
	return x
}

func (x *StringList) FilterRegexp(regExp string) *StringList {
	keepFunc := func(elem string) bool {
		r := regexp.MustCompile(regExp)
		return r.MatchString(elem)
	}
	return x.Filter(keepFunc)
}

func (x *StringList) FilterRegexpNot(regExp string) *StringList {
	keepFunc := func(elem string) bool {
		r := regexp.MustCompile(regExp)
		return r.MatchString(elem)
	}
	return x.Filter(x.invert(keepFunc))
}

func (x *StringList) FilterContains(filter string) *StringList {
	keepFunc := func(elem string) bool {
		return strings.Contains(elem, filter)
	}
	return x.Filter(keepFunc)
}

func (x *StringList) FilterContainsRegexp(regExp string) *StringList {
	keepFunc := func(elem string) bool {
		r := regexp.MustCompile(regExp)
		return len(r.FindStringIndex(elem)) > 0
	}
	return x.Filter(keepFunc)
}

// WithoutEmptyStartEnd removes leading and trailing empty elements.
func (x *StringList) WithoutEmptyStartEnd() *StringList {
	var (
		firstContent = 0
		lastContent  = len(x.strings)
	)

	// leading
	for i, l := range x.strings {
		if strings.TrimSpace(l) != "" {
			break
		}
		firstContent = i + 1
	}

	// trailing
	for i := len(x.strings) - 1; i >= 0; i-- {
		if strings.TrimSpace(x.strings[i]) != "" {
			break
		}
		lastContent = i
	}

	x.strings = x.strings[firstContent:lastContent]
	return x
}

func (x *StringList) WithoutEmpty() *StringList {
	return x.FilterRegexpNot(`^\s*$`)
}

func (x *StringList) WithoutLineComments() *StringList {
	return x.FilterRegexpNot(`^\s*(//|#)`)
}

func (x *StringList) invert(drop func(elem string) (keep bool)) (keep func(elem string) (keep bool)) {
	return func(elem string) (keep bool) {
		return !drop(elem)
	}
}
