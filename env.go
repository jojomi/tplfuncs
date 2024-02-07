package tplfuncs

import (
	htmlTemplate "html/template"
	"os"
	textTemplate "text/template"
)

// EnvHelpers returns a text template FuncMap with env functions
func EnvHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"env":      envFunc,
		"envIsSet": envIsSetFunc,
		"envEq":    envEqFunc,
	}
}

// EnvHelpersHTML returns an HTML template FuncMap with env functions
func EnvHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LinesHelpers())
}

// Doc: `env` returns the value of an ENV variable by name.
func envFunc(key string) string {
	value, _ := os.LookupEnv(key)
	return value
}

// Doc: `envIsSet` checks if an ENV variable is set by its name.
func envIsSetFunc(key string) bool {
	_, isSet := os.LookupEnv(key)
	return isSet
}

// Doc: `envEq` checks if an ENV variable of a given name has the given value.
func envEqFunc(key, value string) bool {
	return envFunc(key) == value
}
