package tplfuncs

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_readFileFunc(t *testing.T) {
	Fs = afero.NewMemMapFs()
	asrt := assert.New(t)

	testFilename := "file.log"
	fileContent := "test file content"

	// setup
	f, err := Fs.Create(testFilename)
	asrt.Nil(err)
	_, err = f.WriteString(fileContent)
	asrt.Nil(err)

	out, err := readFileFunc(testFilename)
	asrt.Nil(err)
	asrt.Equal(fileContent, out)
}

func Test_writeFileWithPermsFunc(t *testing.T) {
	Fs = afero.NewMemMapFs()
	asrt := assert.New(t)

	testFilename := "file.log"
	fileContent := "test file content"

	err := writeFileWithPermsFunc(testFilename, 0600, fileContent)
	asrt.Nil(err)

	// check
	f, err := Fs.Open(testFilename)
	asrt.Nil(err)
	data, err := afero.ReadAll(f)
	asrt.Nil(err)
	asrt.Equal(fileContent, string(data))
}
