package tplfuncs

// DO NOT EDIT, auto-generated tests (see generate/generate_tests.sh)

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"path"
	"strings"
	"testing"
	"text/template"
)

func renderTextTemplate(templateFilename, dataFilename string, funcMap template.FuncMap) (string, error) {
	tmpl, err := template.
		New(path.Base(templateFilename)).
		// Set functions
		Funcs(funcMap).
		// Read template file
		ParseFiles(templateFilename)
	if err != nil {
		return "", err
	}

	// Read data file
	dataFile, err := os.ReadFile(dataFilename)
	if err != nil {
		return "", err
	}

	// Unmarshal YAML data into struct
	var data interface{}
	if err := yaml.Unmarshal(dataFile, &data); err != nil {
		return "", err
	}

	// Execute template with the unmarshalled data
	var b strings.Builder
	if err := tmpl.Execute(&b, data); err != nil {
		return "", err
	}

	// Return the rendered template as a string
	return b.String(), nil
}

{{- newline 2 -}}

{{- /* list all categories */ -}}
{{ $config := readFile "../documentation/data.yml" | parseYAML -}}
{{ $dirs := $config.categories -}}
{{ range $dirs -}}
	{{ $dir := . -}}

	{{ $categoryName := $dir | toCamelCase -}}
	{{ if eq $categoryName "Io" -}}
		{{ $categoryName = "IO" -}}
	{{ end -}}
	{{ if eq $categoryName "Json" -}}
		{{ $categoryName = "JSON" -}}
	{{ end -}}
	{{ if eq $categoryName "Yaml" -}}
		{{ $categoryName = "YAML" -}}
	{{ end -}}

	{{- /* read metadata */ -}}
	{{ $ymlPath := printf "../documentation/functions/%s/metadata.yml" $dir -}}
	{{ if not (fileAt $ymlPath).Exists -}}
		{{ continue -}}
	{{ end -}}
	{{ $metadata := readFile $ymlPath | parseYAML -}}

	{{- /* read funcs per category */ -}}
	{{- $command := printf "sh -c '\"${SOURCE_CHECKER_PATH}\" list_funcs ../%s'" $metadata.source -}}
	{{ $funcs := run $command | parseJSON -}}

	{{ range $funcs -}}
		{{ $func := . -}}

		{{- /* check if testfiles are available for this function */ -}}
		{{ $inputPath := printf "../documentation/functions/%s/%s/input" $dir $func.TemplateName -}}
		{{ if not (fileAt $inputPath).Exists -}}
			{{ continue -}}
		{{ end -}}

		{{ $outputPath := printf "../documentation/functions/%s/%s/output" $dir $func.TemplateName -}}
		{{ if not (fileAt $outputPath).Exists -}}
			{{ continue -}}
		{{ end -}}

// Test {{- $dir | toCamelCase -}} {{- $func.TemplateName | toCamelCase -}} Example validates the documentation example given for the template function `{{- $func.TemplateName -}}` from `{{- $metadata.source -}}`
func Test {{- $categoryName -}} {{- $func.TemplateName | toCamelCase -}} Example(t *testing.T) {
	inputFile := "documentation/functions/ {{- $dir -}} / {{- $func.TemplateName -}} /input"
	outputFile := "documentation/functions/ {{- $dir -}} / {{- $func.TemplateName -}} /output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		{{ $categoryName -}} Helpers(),

		{{- $ymlPath := printf "../documentation/functions/%s/%s/metadata.yml" $dir $func.TemplateName -}}
		{{ $ymlFile := fileAt $ymlPath -}}
		{{ if $ymlFile.Exists -}}
			{{ range ($ymlFile.ContentString | parseYAML).additionalHelpers }}
		{{ . -}} Helpers(),
			{{- end -}}
		{{ end }}
	))
	if err != nil {
		fmt.Printf("could not render template: %s\n", err.Error())
		t.Fail()
	}

	expected, err := os.ReadFile(outputFile)
	if err != nil {
		fmt.Printf("could not read output file at %s\n", outputFile)
		t.Fail()
	}

	assert.Equalf(t, string(expected), result, "template evaluation result is unexpected for input file %s", inputFile)
}

		{{- newline 2 -}}
	{{ end -}}
{{ end -}}