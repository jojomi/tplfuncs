package tplfuncs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
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

func TestExecHomeTemplate(t *testing.T) {
	assrt := assert.New(t)

	funcMap := MakeHTMLFuncMap(ExecHelpersHTML())
	emptyData := struct{}{}

	input := `a {{ execHome "/bin/echo test" -}} b`
	expected := `a test
b`

	output, err := executeTemplateWithHTMLFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "execHome function failed")
	assrt.Equal(expected, output)
}

func TestExecTempTemplate(t *testing.T) {
	assrt := assert.New(t)

	funcMap := MakeHTMLFuncMap(ExecHelpersHTML())
	emptyData := struct{}{}

	input := `a {{ execTemp "/bin/echo test" -}} b`
	expected := `a test
b`

	output, err := executeTemplateWithHTMLFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "execTemp function failed")
	assrt.Equal(expected, output)
}

func TestExecWdTemplate(t *testing.T) {
	assrt := assert.New(t)

	funcMap := MakeHTMLFuncMap(ExecHelpersHTML())
	emptyData := struct{}{}

	input := fmt.Sprintf(`a {{ execWd "/bin/echo test" "%s" -}} b`, os.TempDir())
	expected := `a test
b`

	output, err := executeTemplateWithHTMLFuncMap(funcMap, input, emptyData)
	assrt.Nil(err, "execWd function failed")
	assrt.Equal(expected, output)
}

func TestExecWd(t *testing.T) {
	assrt := assert.New(t)

	homeDir, err := os.UserHomeDir()
	assrt.Nil(err)
	tempDir := os.TempDir()

	a, err := execWdFunc("/bin/ls", homeDir)
	assrt.Nil(err)
	b, err := execWdFunc("/bin/ls", tempDir)
	assrt.Nil(err)
	assrt.NotEqual(a, b)
}
