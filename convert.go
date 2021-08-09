package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"
)

func ToHTMLFuncMap(funcMap textTemplate.FuncMap) htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(funcMap)
}

func ToTextFuncMap(funcMap htmlTemplate.FuncMap) textTemplate.FuncMap {
	return textTemplate.FuncMap(funcMap)
}

func MakeFuncMap(funcMaps ...textTemplate.FuncMap) textTemplate.FuncMap {
	result := textTemplate.FuncMap{}
	for _, funcMap := range funcMaps {
		for key, fun := range funcMap {
			result[key] = fun
		}
	}
	return result
}

func MakeHTMLFuncMap(funcMaps ...htmlTemplate.FuncMap) htmlTemplate.FuncMap {
	result := htmlTemplate.FuncMap{}
	for _, funcMap := range funcMaps {
		for key, fun := range funcMap {
			result[key] = fun
		}
	}
	return result
}
