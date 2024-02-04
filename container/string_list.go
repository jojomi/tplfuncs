package container

import (
	"strings"

	"github.com/jojomi/tplfuncs/text"
)

// StringList is a slice of strings with functions for manipulating it.
type StringList struct {
	strings []string
}

// NewStringList returns an empty StringList.
func NewStringList() *StringList {
	return &StringList{
		strings: make([]string, 0),
	}
}

// NewStringListFromMultilineString returns a StringList created by splitting the given string in lines.
func NewStringListFromMultilineString(input string) *StringList {
	lines := strings.Split(input, "\n")
	return NewStringListFromList(lines)
}

// NewStringListFromList returns a new StringList with given content.
func NewStringListFromList(input []string) *StringList {
	s := NewStringList()
	s.AddAll(input...)
	return s
}

// Add adds a string to the StringList and returns it for chaining.
func (x *StringList) Add(input string) *StringList {
	x.AddAll(input)
	return x
}

// AddAll adds an arbitrary number of strings to the StringList and returns it for chaining.
func (x *StringList) AddAll(input ...string) *StringList {
	x.strings = append(x.strings, input...)
	return x
}

func (x *StringList) AddList(input []string) *StringList {
	x.AddAll(input...)
	return x
}

func (x *StringList) Remove(input string) *StringList {
	x.RemoveAll(input)
	return x
}

func (x *StringList) RemoveAll(input ...string) *StringList {
	newStrings := make([]string, 0, len(x.strings)-len(input))

outer:
	for _, s := range x.strings {
		for _, in := range input {
			if in == s {
				continue outer
			}
		}
		newStrings = append(newStrings, s)
	}

	x.strings = newStrings
	return x
}

func (x *StringList) RemoveList(input *StringList) *StringList {
	return x.RemoveAll(input.All()...)
}

func (x *StringList) Len() int {
	return len(x.strings)
}

func (x *StringList) Empty() bool {
	return x.Len() == 0
}

func (x *StringList) Clear() *StringList {
	x.strings = make([]string, 0)
	return x
}

func (x *StringList) Has(query string) bool {
	for _, s := range x.strings {
		if s == query {
			return true
		}
	}
	return false
}

func (x *StringList) AsText(delim, twoDelim, lastDelim string) string {
	return text.Join(delim, twoDelim, lastDelim, x.strings)
}

func (x *StringList) Joined(delim string) string {
	return strings.Join(x.strings, delim)
}

func (x *StringList) AsString() string {
	return x.Joined("\n")
}

// All returns a slice of all contained strings.
func (x *StringList) All() []string {
	return x.strings
}

// manipulation functions

/*func (x *StringList) Sorted() *StringList {
	return x
}

func (x *StringList) Inverted() *StringList {
	return x
}*/

// internal helpers

// matching
