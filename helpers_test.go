package tplfuncs

import (
	"bytes"
	htmlTemplate "html/template"
	"text/template"
)

func executeTemplateWithFuncMap(funcMap template.FuncMap, input string, data interface{}) (string, error) {
	var outputBuffer bytes.Buffer

	t, err := template.New("tmpl").Funcs(funcMap).Parse(input)
	if err != nil {
		return "", err
	}

	err = t.Execute(&outputBuffer, data)
	if err != nil {
		return "", err
	}

	output := outputBuffer.String()

	return output, nil
}

func executeTemplateWithHTMLFuncMap(funcMap htmlTemplate.FuncMap, input string, data interface{}) (string, error) {
	var outputBuffer bytes.Buffer

	t, err := htmlTemplate.New("tmpl").Funcs(funcMap).Parse(input)
	if err != nil {
		return "", err
	}

	err = t.Execute(&outputBuffer, data)
	if err != nil {
		return "", err
	}

	output := outputBuffer.String()

	return output, nil
}
