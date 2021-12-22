package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// ToHTMLFuncMap converts a text FuncMap to an HTML FuncMap
func ToHTMLFuncMap(funcMap textTemplate.FuncMap) htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(funcMap)
}

// ToTextFuncMap converts an HTML FuncMap to a text FuncMap
func ToTextFuncMap(funcMap htmlTemplate.FuncMap) textTemplate.FuncMap {
	return textTemplate.FuncMap(funcMap)
}

// MakeFuncMap creates a combined text FuncMap from a set of text FuncMaps
func MakeFuncMap(funcMaps ...textTemplate.FuncMap) textTemplate.FuncMap {
	result := textTemplate.FuncMap{}
	for _, funcMap := range funcMaps {
		for key, fun := range funcMap {
			result[key] = fun
		}
	}
	return result
}

// MakeHTMLFuncMap creates a combined HTML FuncMap from a set of HTML FuncMaps
func MakeHTMLFuncMap(funcMaps ...htmlTemplate.FuncMap) htmlTemplate.FuncMap {
	result := htmlTemplate.FuncMap{}
	for _, funcMap := range funcMaps {
		for key, fun := range funcMap {
			result[key] = fun
		}
	}
	return result
}
