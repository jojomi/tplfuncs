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
	return htmlTemplate.FuncMap(LineHelpers())
}

func envFunc(key string) string {
	value, _ := os.LookupEnv(key)
	return value
}

func envIsSetFunc(key string) bool {
	_, isSet := os.LookupEnv(key)
	return isSet
}

func envEqFunc(key, value string) bool {
	return envFunc(key) == value
}
