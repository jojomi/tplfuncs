package container

import (
	"strings"
)

type StringContainer struct {
	builder strings.Builder
}

func NewStringContainer() *StringContainer {
	return &StringContainer{
		builder: strings.Builder{},
	}
}

func (x *StringContainer) Add(input string) *StringContainer {
	x.AddAll(input)
	return x
}

func (x *StringContainer) AddAll(input ...string) *StringContainer {
	for _, i := range input {
		x.builder.WriteString(i)
	}
	return x
}

func (x *StringContainer) Clear() *StringContainer {
	x.builder.Reset()
	return x
}

func (x *StringContainer) String() string {
	return x.builder.String()
}
