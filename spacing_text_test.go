package tplfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpaceDefault(t *testing.T) {
	assrt := assert.New(t)

	input := "ab {{- space -}} c"

	emptyFuncMap := MakeFuncMap()
	emptyData := struct{}{}

	_, err := executeTemplateWithFuncMap(emptyFuncMap, input, emptyData)
	assrt.NotNil(err, "space function already defined in standard library")
}

func TestSpace(t *testing.T) {
	assrt := assert.New(t)

	input := "ab {{- space -}} c"
	expected := "ab c"

	funcMap := MakeFuncMap(SpacingHelpers())
	emptyData := struct{}{}

	out, err := executeTemplateWithFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "space function not loaded")
	assrt.Equal(expected, out)
}

func TestSpaceMulti(t *testing.T) {
	assrt := assert.New(t)

	input := "a {{- space 0 -}} b {{- space 3 -}} c"
	expected := "ab   c"

	funcMap := MakeFuncMap(SpacingHelpers())
	emptyData := struct{}{}

	out, err := executeTemplateWithFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "space function not loaded")
	assrt.Equal(expected, out)
}

func TestNewlineDefault(t *testing.T) {
	assrt := assert.New(t)

	input := "ab {{- newline -}} c"

	emptyFuncMap := MakeFuncMap()
	emptyData := struct{}{}

	_, err := executeTemplateWithFuncMap(emptyFuncMap, input, emptyData)
	assrt.NotNil(err, "newline function already defined in standard library")
}

func TestNewline(t *testing.T) {
	assrt := assert.New(t)

	input := "ab {{- newline -}} c"
	expected := `ab
c`

	funcMap := MakeFuncMap(SpacingHelpers())
	emptyData := struct{}{}

	out, err := executeTemplateWithFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "newline function not loaded")
	assrt.Equal(expected, out)
}

func TestNewlineMulti(t *testing.T) {
	assrt := assert.New(t)

	input := "a {{- newline 0 -}} b {{- newline 2 -}} c"
	expected := `ab

c`

	funcMap := MakeFuncMap(SpacingHelpers())
	emptyData := struct{}{}

	out, err := executeTemplateWithFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "newline function not loaded")
	assrt.Equal(expected, out)
}
