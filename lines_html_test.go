package tplfuncs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchHTML(t *testing.T) {
	assrt := assert.New(t)

	input := `Line 1
Line 2
Line 3`

	template := `<div>{{ .in | match "Line 2" -}}</div>`
	expected := `<div>Line 2</div>`

	funcMap := MakeHTMLFuncMap(LineHelpersHTML())
	data := map[string]string{
		"in": input,
	}

	out, err := executeTemplateWithHTMLFuncMap(funcMap, template, data)
	assrt.Nil(err, "line functions not loaded")
	assrt.Equal(expected, out)
}
