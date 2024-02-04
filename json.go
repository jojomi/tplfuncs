package tplfuncs

import (
	"context"
	"encoding/json"
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/PaesslerAG/gval"
	"github.com/PaesslerAG/jsonpath"
)

// JSONHelpers returns a text template FuncMap with json related functions
func JSONHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"toJSON":              toJSONFunc,
		"parseJSON":           parseJSONFunc,
		"jsonPath":            jsonPathFunc,
		"jsonPathWithDefault": jsonPathWithDefaultFunc,
	}
}

// JSONHelpersHTML returns an HTML template FuncMap with json related functions
func JSONHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(JSONHelpers())
}

func toJSONFunc(input interface{}) (string, error) {
	b, err := json.Marshal(input)
	return string(b), err
}

func parseJSONFunc(jsonString string) (interface{}, error) {
	v := interface{}(nil)
	err := json.Unmarshal([]byte(jsonString), &v)

	return v, err
}

func jsonPathFunc(expression string, jsonData interface{}) (interface{}, error) {
	builder := gval.Full(jsonpath.PlaceholderExtension())

	jsonpathExpr, err := builder.NewEvaluable(expression)
	if err != nil {
		return nil, err
	}

	return jsonpathExpr(context.Background(), jsonData)
}

func jsonPathWithDefaultFunc(expression string, defaultValue interface{}, jsonData interface{}) (interface{}, error) {
	result, err := jsonPathFunc(expression, jsonData)
	if err != nil {
		return defaultValue, nil
	}
	return result, nil
}
