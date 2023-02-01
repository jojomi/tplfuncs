package tplfuncs

import (
	"github.com/spf13/cast"
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// CastHelpers returns a text template FuncMap with cast functions
func CastHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"toString": cast.ToString,
		// toStringAnyMapFunc,
	}
}

// CastHelpersHTML returns an HTML template FuncMap with cast functions
func CastHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(CastHelpers())
}

/*
func toStringAnyMapFunc(input interface{}) (container.StringAnyMap, error) {
	var m = map[string]interface{}{}

	switch v := i.(type) {
	case map[interface{}]interface{}:
		for k, val := range v {
			m[ToString(k)] = val
		}
		return m, nil
	case map[string]interface{}:
		return v, nil
	case string:
		err := jsonStringToObject(v, &m)
		return m, err
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to map[string]interface{}", i, i)
	}
}
*/
