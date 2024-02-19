package tplfuncs

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func Test_writeFileWithPermsFunc(t *testing.T) {
	Fs = afero.NewMemMapFs()
	asrt := assert.New(t)

	testFilename := "file.log"
	fileContent := "test file content"

	err := writeFileWithPermsFunc(testFilename, 0o600, fileContent)
	asrt.Nil(err)

	// check
	f, err := Fs.Open(testFilename)
	asrt.Nil(err)
	data, err := afero.ReadAll(f)
	asrt.Nil(err)
	asrt.Equal(fileContent, string(data))
}
