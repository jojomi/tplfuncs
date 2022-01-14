package tplfuncs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExec(t *testing.T) {
	assrt := assert.New(t)

	input := `/bin/echo abc`
	expected := "abc\n"

	out, err := execFunc(input)
	assrt.Nil(err)
	assrt.Equal(expected, out)
}

func TestExecFail(t *testing.T) {
	assrt := assert.New(t)

	input := `/bin/invalid_binary_name`

	_, err := execFunc(input)
	assrt.NotNil(err)
}

func TestExecTemplate(t *testing.T) {
	assrt := assert.New(t)

	funcMap := MakeHTMLFuncMap(ExecHelpersHTML())
	emptyData := struct{}{}

	input := `a {{ exec "/bin/echo test" -}} b`
	expected := `a test
b`

	output, err := executeTemplateWithHTMLFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "exec function failed")
	assrt.Equal(expected, output)
}
