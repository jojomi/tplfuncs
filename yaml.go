package tplfuncs

import (
	"gopkg.in/yaml.v3"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// YAMLHelpers returns a text template FuncMap with yaml related functions
func YAMLHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"toYAML":    toYAMLFunc,
		"parseYAML": parseYAMLFunc,
	}
}

// YAMLHelpersHTML returns an HTML template FuncMap with yaml related functions
func YAMLHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(YAMLHelpers())
}

func toYAMLFunc(input interface{}) (string, error) {
	yamlData, err := yaml.Marshal(input)
	if err != nil {
		return "", err
	}

	return string(yamlData), nil
}

func parseYAMLFunc(yamlString string) (interface{}, error) {
	// unmarshal YAML into a generic interface{}
	var data interface{}
	err := yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
