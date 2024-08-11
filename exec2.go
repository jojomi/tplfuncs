package tplfuncs

import (
	"github.com/jojomi/gorun"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// Exec2Helpers returns a text template FuncMap with functions related to command execution
func Exec2Helpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"runLocal": runLocal,
		"run":      run,
		//"runShell": runShell,
	}
}

func run() *gorun.Runner {
	return gorun.New()
}

func runLocal(command string) (string, error) {
	c := gorun.LocalCommandFrom(command)
	res, err := gorun.NewWithCommand(c).Exec()
	if err != nil {
		return res.CombinedOutput(), err
	}
	return res.CombinedOutput(), res.CombinedError()
}

// Exec2HelpersHTML returns an HTML template FuncMap with functions related to command execution
func Exec2HelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(Exec2Helpers())
}
