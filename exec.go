package tplfuncs

import (
	"github.com/juju/errors"
	htmlTemplate "html/template"
	"os"
	"os/exec"
	"strings"
	textTemplate "text/template"
)

// ExecHelpers returns a text template FuncMap with functions related to command execution
func ExecHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"exec":     execFunc,
		"execHome": execHomeFunc,
		"execTemp": execTempFunc,
		"execWd":   execWdFunc,
	}
}

// ExecHelpersHTML returns an HTML template FuncMap with functions related to command execution
func ExecHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(ExecHelpers())
}

func execWdFunc(command, workingDir string) (string, error) {
	parts := strings.Fields(command)

	var err error
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = workingDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Annotatef(err, "Output: "+string(out))
	}
	return string(out), nil
}

func execFunc(command string) (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return execWdFunc(command, workDir)
}

func execHomeFunc(command string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return execWdFunc(command, homeDir)
}

func execTempFunc(command string) (string, error) {
	return execWdFunc(command, os.TempDir())
}
