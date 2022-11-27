package text

import (
	"fmt"
	"strings"
)

func Join(delim, twoDelim, lastDelim string, input []string) string {
	var result strings.Builder
	for i, elem := range input {
		result.WriteString(fmt.Sprintf("%v", elem))

		// delims
		if i == 0 && len(input) == 2 {
			result.WriteString(twoDelim)
			continue
		}
		if i < len(input)-2 {
			result.WriteString(delim)
			continue
		}
		if i == len(input)-2 {
			result.WriteString(lastDelim)
		}
	}
	return result.String()
}

func JoinStringer(delim, twoDelim, lastDelim string, input []fmt.Stringer) string {
	r := make([]string, len(input))
	for i, in := range input {
		r[i] = in.String()
	}
	return Join(delim, twoDelim, lastDelim, r)
}
