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

// TestSpacingSpaceExample validates the documentation example given for the template function `space` from `spacing.go`
func TestSpacingSpaceExample(t *testing.T) {
	inputFile := "documentation/functions/spacing/space/input"
	outputFile := "documentation/functions/spacing/space/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		SpacingHelpers(),
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

// TestSpacingNewlineExample validates the documentation example given for the template function `newline` from `spacing.go`
func TestSpacingNewlineExample(t *testing.T) {
	inputFile := "documentation/functions/spacing/newline/input"
	outputFile := "documentation/functions/spacing/newline/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		SpacingHelpers(),
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

// TestSpacingNoopExample validates the documentation example given for the template function `noop` from `spacing.go`
func TestSpacingNoopExample(t *testing.T) {
	inputFile := "documentation/functions/spacing/noop/input"
	outputFile := "documentation/functions/spacing/noop/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		SpacingHelpers(),
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

// TestStringTrimExample validates the documentation example given for the template function `trim` from `string.go`
func TestStringTrimExample(t *testing.T) {
	inputFile := "documentation/functions/string/trim/input"
	outputFile := "documentation/functions/string/trim/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		StringHelpers(),
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

// TestStringEqFoldExample validates the documentation example given for the template function `eqFold` from `string.go`
func TestStringEqFoldExample(t *testing.T) {
	inputFile := "documentation/functions/string/eqFold/input"
	outputFile := "documentation/functions/string/eqFold/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		StringHelpers(),
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

// TestStringStringContainsExample validates the documentation example given for the template function `stringContains` from `string.go`
func TestStringStringContainsExample(t *testing.T) {
	inputFile := "documentation/functions/string/stringContains/input"
	outputFile := "documentation/functions/string/stringContains/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		StringHelpers(),
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

// TestStringToUpperCaseExample validates the documentation example given for the template function `toUpperCase` from `string.go`
func TestStringToUpperCaseExample(t *testing.T) {
	inputFile := "documentation/functions/string/toUpperCase/input"
	outputFile := "documentation/functions/string/toUpperCase/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		StringHelpers(),
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

// TestStringToLowerCaseExample validates the documentation example given for the template function `toLowerCase` from `string.go`
func TestStringToLowerCaseExample(t *testing.T) {
	inputFile := "documentation/functions/string/toLowerCase/input"
	outputFile := "documentation/functions/string/toLowerCase/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		StringHelpers(),
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

// TestStringTrimPrefixExample validates the documentation example given for the template function `trimPrefix` from `string.go`
func TestStringTrimPrefixExample(t *testing.T) {
	inputFile := "documentation/functions/string/trimPrefix/input"
	outputFile := "documentation/functions/string/trimPrefix/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		StringHelpers(),
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

// TestStringTrimSuffixExample validates the documentation example given for the template function `trimSuffix` from `string.go`
func TestStringTrimSuffixExample(t *testing.T) {
	inputFile := "documentation/functions/string/trimSuffix/input"
	outputFile := "documentation/functions/string/trimSuffix/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		StringHelpers(),
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

// TestLinesLineExample validates the documentation example given for the template function `line` from `lines.go`
func TestLinesLineExample(t *testing.T) {
	inputFile := "documentation/functions/lines/line/input"
	outputFile := "documentation/functions/lines/line/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		LinesHelpers(),
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

// TestLanguageJoinTextExample validates the documentation example given for the template function `joinText` from `language.go`
func TestLanguageJoinTextExample(t *testing.T) {
	inputFile := "documentation/functions/language/joinText/input"
	outputFile := "documentation/functions/language/joinText/output"

	result, err := renderTextTemplate(inputFile, "documentation/example-data.yml", MakeFuncMap(
		LanguageHelpers(),
		ContainerHelpers(),
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

