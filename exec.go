package tplfuncs

import (
	htmlTemplate "html/template"
	"os"
	"os/exec"
	textTemplate "text/template"

	"github.com/juju/errors"
	"github.com/kballard/go-shellquote"
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
	parts, err := shellquote.Split(command)
	if err != nil {
		return "", err
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = workingDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Annotatef(err, "Working Dir: %s, Output: %s", workingDir, string(out))
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
