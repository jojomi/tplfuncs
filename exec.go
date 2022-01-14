package tplfuncs

import (
	htmlTemplate "html/template"
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
	out, err := exec.Command(parts[0], parts[1:]...).CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
