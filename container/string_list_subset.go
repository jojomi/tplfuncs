package container

import (
	"fmt"
	"regexp"
	"strings"
)

func (x *StringList) First(count int) *StringList {
	end := count
	if len(x.strings) < end {
		end = len(x.strings)
	}
	x.strings = x.strings[0:end]
	return x
}

func (x *StringList) Last(count int) *StringList {
	end := count
	if len(x.strings) < end {
		end = len(x.strings)
	}
	x.strings = x.strings[0:end]
	return x
}

func (x *StringList) WithoutFirst(count int) *StringList {
	end := count
	if len(x.strings) < end {
		end = len(x.strings)
	}
	x.strings = x.strings[0:end]
	return x
}

func (x *StringList) WithoutLast(count int) *StringList {
	end := count
	if len(x.strings) < end {
		end = len(x.strings)
	}
	x.strings = x.strings[0:end]
	return x
}

// TODO: continue

func (x *StringList) SubsetRegexp(startRegexp, endRegexp string) (*StringList, error) {
	var (
		start  = -1
		end    = -1
		rStart = regexp.MustCompile(startRegexp)
		rEnd   = regexp.MustCompile(endRegexp)
	)

	if startRegexp == "" {
		start = 0
	}
	if endRegexp == "" {
		end = len(x.strings)
	}

	for i, elem := range x.strings {
		if start < 0 && rStart.MatchString(elem) {
			start = i
			continue
		}
		if start > -1 && end < 0 && rEnd.MatchString(elem) {
			end = i + 1
			break
		}
	}

	if start == -1 {
		return x, fmt.Errorf("could not find start line matching regexp %s", startRegexp)
	}
	if end == -1 {
		return x, fmt.Errorf("could not find end line matching regexp %s", endRegexp)
	}

	x.strings = x.strings[start:end]
	return x, nil
}

// SubsetContainsAfter returns the first line with the given subset string, expanded to _count_ lines in total.
func (x *StringList) SubsetContainsAfter(substring string, count int) (*StringList, error) {
	var (
		start = -1
		end   = -1
	)

	for i, elem := range x.strings {
		if start < 0 && strings.Contains(elem, substring) {
			start = i
			end = i + count
			continue
		}
	}

	if start == -1 {
		return x, fmt.Errorf("could not find start line matching substring %s", substring)
	}

	x.strings = x.strings[start:end]
	return x, nil
}
