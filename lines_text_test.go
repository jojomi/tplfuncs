package tplfuncs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLine(t *testing.T) {
	assrt := assert.New(t)

	input := `Line 1
Line 2
Line 3`

	assrt.Equal("", lineFunc(0, input))
	assrt.Equal("Line 1", lineFunc(1, input))
	assrt.Equal("Line 2", lineFunc(2, input))
	assrt.Equal("Line 3", lineFunc(3, input))
	assrt.Equal("", lineFunc(4, input))
}

func TestLineOrErr(t *testing.T) {
	assrt := assert.New(t)

	input := `Line 1
Line 2
Line 3`

	v, err := lineOrErrFunc(1, input)
	assrt.Nil(err)
	assrt.Equal("Line 1", v)

	_, err = lineOrErrFunc(0, input)
	assrt.NotNil(err)

	_, err = lineOrErrFunc(4, input)
	assrt.NotNil(err)
}

func TestTrim(t *testing.T) {
	assrt := assert.New(t)

	input := `
  Line 1
Line 2
Line 3  

`
	expected := `  Line 1
Line 2
Line 3  `

	assrt.Equal(expected, trimFunc(input))
}

func TestTrimAll(t *testing.T) {
	assrt := assert.New(t)

	input := `  Line 1
Line 2  
  Line 3  
`
	expected := `Line 1
Line 2
Line 3`

	assrt.Equal(expected, trimAllFunc(input))
}

func TestMatch(t *testing.T) {
	assrt := assert.New(t)

	input := `Invalid
Valid
Invalid
Invalid`
	expected := `Valid`

	assrt.Equal(expected, matchFunc("Valid", input))
}

func TestNotMatch(t *testing.T) {
	assrt := assert.New(t)

	input := `Invalid
Valid
Invalid
Invalid`
	expected := `Valid`

	assrt.Equal(expected, notMatchFunc("Invalid", input))
}

func TestWithoutEmptyLines(t *testing.T) {
	assrt := assert.New(t)

	input := `
 A

B

`
	expected := ` A
B`

	assrt.Equal(expected, withoutEmptyLinesFunc(input))
}

func TestWithoutLineComments(t *testing.T) {
	assrt := assert.New(t)

	input := `# Line 1
  # Comment 1
Line 2
// Line 3
	// Comment 3
Line 4`
	expected := `Line 2
Line 4`

	assrt.Equal(expected, withoutLineCommentsFunc(input))
}

func TestHeadDefault(t *testing.T) {
	assrt := assert.New(t)

	input := "ab {{- head -}} c"

	emptyFuncMap := MakeFuncMap()
	data := map[string]string{
		"input": `Line 1
Line 2
Line 3`,
	}

	_, err := executeTemplateWithFuncMap(emptyFuncMap, input, data)
	assrt.NotNil(err, "head function already defined in standard library")
}

func TestHead(t *testing.T) {
	assrt := assert.New(t)

	input := "{{- .input | head 2 -}}"
	expected := `Line 1
Line 2`

	funcMap := MakeFuncMap(LineHelpers())
	data := map[string]string{
		"input": `Line 1
Line 2
Line 3`,
	}

	out, err := executeTemplateWithFuncMap(funcMap, input, data)
	assrt.Nil(err, "head function not loaded")
	assrt.Equal(expected, out)
}

func TestSkipHead(t *testing.T) {
	assrt := assert.New(t)

	input := "{{- .input | skipHead 2 -}}"
	expected := `Line 3`

	funcMap := MakeFuncMap(LineHelpers())
	data := map[string]string{
		"input": `Line 1
Line 2
Line 3`,
	}

	out, err := executeTemplateWithFuncMap(funcMap, input, data)
	assrt.Nil(err, "skipHead function not loaded")
	assrt.Equal(expected, out)
}

func TestTail(t *testing.T) {
	assrt := assert.New(t)

	input := "{{- .input | tail 2 -}}"
	expected := `Line 2
Line 3`

	funcMap := MakeFuncMap(LineHelpers())
	data := map[string]string{
		"input": `Line 1
Line 2
Line 3`,
	}

	out, err := executeTemplateWithFuncMap(funcMap, input, data)
	assrt.Nil(err, "tail function not loaded")
	assrt.Equal(expected, out)
}

func TestSkipTail(t *testing.T) {
	assrt := assert.New(t)

	input := "{{- .input | skipTail 2 -}}"
	expected := `Line 1`

	funcMap := MakeFuncMap(LineHelpers())
	data := map[string]string{
		"input": `Line 1
Line 2
Line 3`,
	}

	out, err := executeTemplateWithFuncMap(funcMap, input, data)
	assrt.Nil(err, "skipTail function not loaded")
	assrt.Equal(expected, out)
}

func TestRegexpReplaceLine(t *testing.T) {
	assrt := assert.New(t)

	input := `{{- .input | regexpReplaceLine "text (\\d+)" "$1" -}}`
	expected := `13
2
text`

	funcMap := MakeFuncMap(LineHelpers())
	data := map[string]string{
		"input": `text 13
text 2
text`,
	}

	out, err := executeTemplateWithFuncMap(funcMap, input, data)
	assrt.Nil(err, "regexpReplaceLine function not loaded")
	assrt.Equal(expected, out)
}

func TestWrapLines(t *testing.T) {
	assrt := assert.New(t)

	input := `{{- .input | wrapLines "(" ")" -}}`
	expected := `(Line 1)
(Line 2)
(Line 3)`

	funcMap := MakeFuncMap(LineHelpers())
	data := map[string]string{
		"input": `Line 1
Line 2
Line 3`,
	}

	out, err := executeTemplateWithFuncMap(funcMap, input, data)
	assrt.Nil(err, "wrapLines function not loaded")
	assrt.Equal(expected, out)
}

