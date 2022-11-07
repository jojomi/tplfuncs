package tplfuncs

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_FileExistsFunc(t *testing.T) {
	Fs = afero.NewMemMapFs()
	asrt := assert.New(t)

	testFilename := "file.log"

	asrt.False(fileExistsFunc(testFilename))

	// create file
	_, err := Fs.Create(testFilename)
	asrt.Nil(err)
	asrt.True(fileExistsFunc(testFilename))

	// create dir
	err = Fs.Remove(testFilename)
	asrt.Nil(err)
	err = Fs.MkdirAll(testFilename, 0700)
	asrt.Nil(err)
	asrt.False(fileExistsFunc(testFilename))
}

func Test_DirExistsFunc(t *testing.T) {
	Fs = afero.NewMemMapFs()
	asrt := assert.New(t)

	testDirname := "log-dir"

	asrt.False(dirExistsFunc(testDirname))

	// create dir
	err := Fs.MkdirAll(testDirname, 0700)
	asrt.Nil(err)
	asrt.True(dirExistsFunc(testDirname))

	// create file
	err = Fs.RemoveAll(testDirname)
	asrt.Nil(err)
	_, err = Fs.Create(testDirname)
	asrt.Nil(err)
	asrt.False(dirExistsFunc(testDirname))
}

func Test_GlobFunc(t *testing.T) {
	Fs = afero.NewMemMapFs()
	asrt := assert.New(t)

	testFilenames := []string{"file.log", "another.log", "some.bin", "last.log"}
	for _, t := range testFilenames {
		_, err := Fs.Create(t)
		asrt.Nil(err)
	}

	globResult, err := globFunc("*.log")
	asrt.Nil(err)
	asrt.Equal([]string{"another.log", "file.log", "last.log"}, globResult)
}

func Test_EnsureDirFunc(t *testing.T) {
	Fs = afero.NewMemMapFs()
	asrt := assert.New(t)

	testDirname := "log-dir"

	asrt.False(dirExistsFunc(testDirname))

	// ensure dir
	err := ensureDirFunc(testDirname)
	asrt.Nil(err)
	asrt.True(dirExistsFunc(testDirname))

	err = ensureDirFunc(testDirname)
	asrt.Nil(err)
	asrt.True(dirExistsFunc(testDirname))
}

func Test_IsMinFileSizeFunc(t *testing.T) {
	Fs = afero.NewMemMapFs()
	asrt := assert.New(t)

	testFilename := "file.log"

	f, err := Fs.Create(testFilename)
	asrt.Nil(err)
	asrt.False(isMinFileSizeFunc(testFilename, 10))

	f.WriteString("sample file content")
	asrt.True(isMinFileSizeFunc(testFilename, 10))
}
