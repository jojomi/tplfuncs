package tplfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpaceDefaultHTML(t *testing.T) {
	assrt := assert.New(t)

	input := `<div class="a {{-space -}} b">ab {{- space -}} c</div>`

	emptyFuncMap := MakeHTMLFuncMap()
	emptyData := struct{}{}

	_, err := executeTemplateWithHTMLFuncMap(emptyFuncMap, input, emptyData)
	assrt.NotNil(err, "space function already defined in standard library")
}

func TestSpaceHTML(t *testing.T) {
	assrt := assert.New(t)

	input := `<div class="a {{- space -}} b">ab {{- space -}} c</div>`
	expected := `<div class="a b">ab c</div>`

	funcMap := MakeHTMLFuncMap(SpacingHelpersHTML())
	emptyData := struct{}{}

	out, err := executeTemplateWithHTMLFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "space function not loaded")
	assrt.Equal(expected, out)
}

func TestSpaceMultiHTML(t *testing.T) {
	assrt := assert.New(t)

	input := `<div class="a {{- space 0 -}} b">ab {{- space 2 -}} c</div>`
	expected := `<div class="ab">ab  c</div>`

	funcMap := MakeHTMLFuncMap(SpacingHelpersHTML())
	emptyData := struct{}{}

	out, err := executeTemplateWithHTMLFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "space function not loaded")
	assrt.Equal(expected, out)
}

func TestNewlineDefaultHTML(t *testing.T) {
	assrt := assert.New(t)

	input := `<div class="a {{-newline -}} b">ab {{- newline -}} c</div>`

	emptyFuncMap := MakeHTMLFuncMap()
	emptyData := struct{}{}

	_, err := executeTemplateWithHTMLFuncMap(emptyFuncMap, input, emptyData)
	assrt.NotNil(err, "newline function already defined in standard library")
}

func TestNewlineHTML(t *testing.T) {
	assrt := assert.New(t)

	input := `<div class="a {{- newline -}} b">ab {{- newline 2 -}} c</div>`
	expected := `<div class="a
b">ab

c</div>`

	funcMap := MakeHTMLFuncMap(SpacingHelpersHTML())
	emptyData := struct{}{}

	out, err := executeTemplateWithHTMLFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "newline function not loaded")
	assrt.Equal(expected, out)
}

func TestNewlineMultiHTML(t *testing.T) {
	assrt := assert.New(t)

	input := `<div class="a {{- newline 1 -}} b">{{ newline 0 }}ab {{- newline 2 -}} c</div>`
	expected := `<div class="a
b">ab

c</div>`

	funcMap := MakeHTMLFuncMap(SpacingHelpersHTML())
	emptyData := struct{}{}

	out, err := executeTemplateWithHTMLFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "newline function not loaded")
	assrt.Equal(expected, out)
}
