package tplfuncs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToText(t *testing.T) {
	assert := assert.New(t)

	input := `a {{- space -}} b {{- newline -}} c`
	expected := `a b
c`

	funcMap := ToTextFuncMap(MakeHTMLFuncMap(SpacingHelpersHTML()))
	emptyData := struct{}{}

	out, err := executeTemplateWithFuncmap(funcMap, input, emptyData)
	assert.Nil(err, "spacing functions not loaded")
	assert.Equal(expected, out)
}

func TestConvertToHTML(t *testing.T) {
	assert := assert.New(t)

	input := `<div class="a {{- space -}} b">ab {{- newline -}} c</div>`
	expected := `<div class="a b">ab
c</div>`

	funcMap := ToHTMLFuncMap(MakeFuncMap(SpacingHelpers()))
	emptyData := struct{}{}

	out, err := executeTemplateWithHTMLFuncmap(funcMap, input, emptyData)
	assert.Nil(err, "spacing functions not loaded")
	assert.Equal(expected, out)
}
