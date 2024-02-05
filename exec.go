package tplfuncs

import (
	"fmt"
	htmlTemplate "html/template"
	"os"
	"os/exec"
	textTemplate "text/template"

	"github.com/jojomi/gorun"

	"github.com/juju/errors"
	"github.com/kballard/go-shellquote"
)

// ExecHelpers returns a text template FuncMap with functions related to command execution
func ExecHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"exec":       execFunc,
		"execNoFail": execNoFailFunc,
		// "execSilent": and all others
		"execHome": execHomeFunc,
		"execTemp": execTempFunc,
		"execWd":   execWdFunc,

		//"commandExists": commandExistsFunc,
		"run":              runFunc,
		"runSSH":           runSSHFunc,
		"runner":           runnerFunc,
		"localCommandFrom": localCommandFromFunc,
	}
}

// ExecHelpersHTML returns an HTML template FuncMap with functions related to command execution
func ExecHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(ExecHelpers())
}

func execWdFunc(command, workingDir string) (string, error) {
	return execInternalFunc(command, workingDir, false)
}

func execInternalFunc(command, workingDir string, noFail bool) (string, error) {
	parts, err := shellquote.Split(command)
	if err != nil {
		return "", err
	}

	envVal, _ := os.LookupEnv("TPLFUNCS_DEBUG_EXEC")
	debug := envVal != ""

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Dir = workingDir

	if debug {
		fmt.Printf("executing: %s...\n", cmd.String())
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		if noFail && cmd.ProcessState.Exited() && cmd.ProcessState.ExitCode() > 0 {
			return string(out), nil
		}
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

func execNoFailFunc(command string) (string, error) {
	workDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return execInternalFunc(command, workDir, true)
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

func runFunc(command string) (string, error) {
	r := getRunner().WithCommand(gorun.LocalCommandFrom(command))

	if os.Getenv("IO_LOG_RUN_FUNC") == "true" {
		r = r.LogCommand(true)
	}

	rr, err := r.Exec()
	if err != nil {
		return "", err
	}
	return rr.CombinedOutput(), nil
}

func getRunner() *gorun.Runner {
	r := gorun.New().WithoutStdout().WithoutStderr()

	if os.Getenv("IO_LOG_RUN_FUNC") == "true" {
		r = r.LogCommand(true)
	}

	return r
}

func runSSHFunc(sshAlias, command string) (string, error) {
	c := gorun.NewSSHCommandFrom(sshAlias, gorun.LocalCommandFrom(command))
	r := getRunner().WithCommand(c)

	rr, err := r.Exec()
	if err != nil {
		return "", err
	}
	return rr.CombinedOutput(), nil
}

func runnerFunc() *gorun.Runner {
	return getRunner()
}

func localCommandFromFunc(command string) *gorun.LocalCommand {
	return gorun.LocalCommandFrom(command)
}
