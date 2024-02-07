package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/spf13/cast"
)

// CastHelpers returns a text template FuncMap with cast functions
func CastHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"toBool":        toBoolFunc,
		"toString":      toStringFunc,
		"toInt":         toIntFunc,
		"toFloat":       toFloatFunc,
		"toStringSlice": toStringSliceFunc,
	}
}

// CastHelpersHTML returns an HTML template FuncMap with cast functions
func CastHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(CastHelpers())
}

// Doc: `toBool` converts any given input to a bool.
func toBoolFunc(input interface{}) bool {
	return cast.ToBool(input)
}

// Doc: `toString` converts any given input to a string.
func toStringFunc(input interface{}) string {
	return cast.ToString(input)
}

// Doc: `toInt` converts any given input to an int.
func toIntFunc(input interface{}) int {
	return cast.ToInt(input)
}

// Doc: `toFloat` converts any given content to a float64.
func toFloatFunc(input interface{}) float64 {
	return cast.ToFloat64(input)
}

// Doc: `toStringSlice` converts any given input to a string slice.
func toStringSliceFunc(input interface{}) []string {
	return cast.ToStringSlice(input)
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
