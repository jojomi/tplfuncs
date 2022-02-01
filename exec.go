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
		"exec": execFunc,
	}
}

// ExecHelpersHTML returns an HTML template FuncMap with functions related to command execution
func ExecHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(ExecHelpers())
}

func execFunc(command string) (string, error) {
	parts := strings.Fields(command)

	var err error
	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir, err = os.UserHomeDir()
	if err != nil {
		return "", err
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.Annotatef(err, "Output: "+string(out))
	}
	return string(out), nil
}