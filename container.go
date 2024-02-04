package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/jojomi/tplfuncs/container"
)

// ContainerHelpers returns a text template FuncMap with functions related to container processing.
// The idea is to allow more sophisticated processing inside the template when adding custom functions is not possible.
// Make sure to catch the return value in a (possible unused) variable when you just call these functions for their side effects.
func ContainerHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"makeStringList":                    makeStringListFunc,
		"makeStringListFromList":            container.NewStringListFromList,
		"makeAnyList":                       makeAnyListFunc,
		"makeStringAnyMap":                  makeStringAnyMapFunc,
		"makeStringAnyMapFromMap":           makeStringAnyMapFromMapFunc,
		"makeStringContainer":               makeStringContainerFunc,
		"makeStringListFromMultilineString": container.NewStringListFromMultilineString,
	}
}

// ContainerHelpersHTML returns an HTML template FuncMap with functions related to container processing
func ContainerHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(ContainerHelpers())
}

func makeStringListFunc(input ...string) *container.StringList {
	s := container.NewStringList()
	s.AddAll(input...)
	return s
}

func makeAnyListFunc(input ...interface{}) *container.AnyList {
	s := container.NewAnyList()
	s.Add(input...)
	return s
}

func makeStringAnyMapFunc(input ...interface{}) *container.StringAnyMap {
	s := container.NewStringAnyMap()
	s.AddAll(input...)
	return s
}

func makeStringAnyMapFromMapFunc(data map[string]interface{}) *container.StringAnyMap {
	return container.NewStringAnyMapFromMap(data)
}

func makeStringContainerFunc(input ...string) *container.StringContainer {
	s := container.NewStringContainer()
	s.AddAll(input...)
	return s
}
