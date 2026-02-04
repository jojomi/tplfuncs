package tplfuncs

import (
	"bufio"
	"fmt"
	htmlTemplate "html/template"
	"regexp"
	"sort"
	"strings"
	textTemplate "text/template"
)

// LinesHelpers returns a text template FuncMap with functions related to line processing
func LinesHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"getLines":            getLinesFunc,
		"line":                lineFunc,
		"lineOrErr":           lineOrErrFunc,
		"head":                headFunc,
		"skipHead":            skipHeadFunc,
		"tail":                tailFunc,
		"skipTail":            skipTailFunc,
		"sortLines":           sortLinesFunc,
		"trimLines":           trimLinesFunc,
		"trimAll":             trimAllFunc,
		"wrapLines":           wrapLinesFunc,
		"withoutLineComments": withoutLineCommentsFunc,
		"withoutEmptyLines":   withoutEmptyLinesFunc,
		"match":               matchFunc,
		"notMatch":            notMatchFunc,
		"regexpReplaceLine":   regexpReplaceLineFunc,
		"joinLines":           asString,
		"indentSpaceLines":    indentSpaceLinesFunc,
		"indentTabLines":      indentTabLinesFunc,
		"prefixLines":         prefixLinesFunc,
		"isSingleLine":        isSingleLineFunc,
		"isMultiline":         isMultilineFunc,
		"filterLinesPrefix":   filterLinesPrefixFunc,
		"withoutLinesPrefix":  withoutLinesPrefixFunc,
	}
}

// LinesHelpersHTML returns an HTML template FuncMap with functions related to line processing
func LinesHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LinesHelpers())
}

// Doc: `lineOrErr` returns a single line from the multiline input. The index is 1-based. Returns an error, if the line does not exist.
func lineOrErrFunc(number int, input string) (string, error) {
	lines := getLines(input)
	if number < 1 || number > len(lines) {
		return "", fmt.Errorf("line ")
	}
	return lines[number-1], nil
}

// Doc: `line` returns a single line from the multiline input. The index is 1-based. Returns an empty string, if the line does not exist.
func lineFunc(number int, input string) string {
	lines := getLines(input)
	if number < 1 || number > len(lines) {
		return ""
	}
	return lines[number-1]
}

// Doc: Return the multiline input sorted alphabetically line by line.
func sortLinesFunc(input string) string {
	lines := getLines(input)
	sort.Strings(lines)
	return asString(lines)
}

// Doc: `head` returns the first n lines of a multiline string as one string, or all of it if there is less than n lines in total.
func headFunc(count int, input string) string {
	lines := getLines(input)
	if count < len(lines) {
		lines = lines[:count]
	}
	return asString(lines)
}

// Doc: `skipHead` returns the multiline string given without the first n lines or an empty string if there is less than n lines in total.
func skipHeadFunc(count int, input string) string {
	lines := getLines(input)
	if count < len(lines) {
		lines = lines[count:]
	}
	return asString(lines)
}

// Doc: `tail` returns the last n lines of a multiline string as one string, or all of it if there is less than n lines in total.
func tailFunc(count int, input string) string {
	lines := getLines(input)
	if count < len(lines) {
		lines = lines[len(lines)-count:]
	}
	return asString(lines)
}

// Doc: `skipTail` returns the multiline string given without the last n lines or an empty string if there is less than n lines in total.
func skipTailFunc(count int, input string) string {
	lines := getLines(input)
	if count < len(lines) {
		lines = lines[0 : len(lines)-count]
	}
	return asString(lines)
}

// Doc: `trimLines` returns the multiline string given without leading and trailing empty lines.
func trimLinesFunc(input string) string {
	lines := getLines(input)

	var (
		firstContentLine = 0
		lastContentLine  = len(lines)
	)

	// leading
	for i, l := range lines {
		if strings.TrimSpace(l) != "" {
			break
		}
		firstContentLine = i + 1
	}

	// trailing
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) != "" {
			break
		}
		lastContentLine = i
	}

	return asString(lines[firstContentLine:lastContentLine])
}

// Doc: `trimAll` returns the multiline string given with leading and trailing space removed for any line individually.
func trimAllFunc(input string) string {
	lines := getLines(input)
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return asString(lines)
}

// Doc: `notMatch` does return a string with all lines from the given multiline string that do not match the regexp given.
func notMatchFunc(regExp, input string) string {
	lines := getLines(input)
	r := regexp.MustCompile(regExp)
	result := make([]string, 0)

	for _, line := range lines {
		if r.MatchString(line) {
			continue
		}
		result = append(result, line)
	}

	return asString(result)
}

// Doc: `match` does return a string with all lines from the given multiline string that do match the regexp given.
func matchFunc(regExp, input string) string {
	lines := getLines(input)
	r := regexp.MustCompile(regExp)
	result := make([]string, 0)

	for _, line := range lines {
		if !r.MatchString(line) {
			continue
		}
		result = append(result, line)
	}

	return asString(result)
}

// Doc: `withoutEmptyLines` returns the multiline string given without empty lines.
func withoutEmptyLinesFunc(input string) string {
	return notMatchFunc(`^\s*$`, input)
}

// Doc: `withoutLineComments` returns the multiline string given without line comments (lines starting with optional whitespace and // or #).
func withoutLineCommentsFunc(input string) string {
	return notMatchFunc(`^\s*(//|#)`, input)
}

// Doc: `wrapLines` returns the multiline string with every single line wrapped with the given leading and trailing string.
func wrapLinesFunc(leading, trailing, input string) string {
	lines := getLines(input)
	for i, line := range lines {
		lines[i] = leading + line + trailing
	}
	return asString(lines)
}

// Doc: `indentSpaceLines` returns the multiline string given with every line indented by additional n spaces.
func indentSpaceLinesFunc(spaceCount int, input string) string {
	lines := getLines(input)
	for i, line := range lines {
		lines[i] = strings.Repeat(" ", spaceCount) + line
	}
	return asString(lines)
}

// Doc: `prefixLines` returns the multiline string given with every line prefixed with the string given.
func prefixLinesFunc(prefix string, input string) string {
	lines := getLines(input)
	for i, line := range lines {
		lines[i] = prefix + line
	}
	return asString(lines)
}

// Doc: `indentTabLines` returns the multiline string given with every line indented by additional n tab characters.
func indentTabLinesFunc(tabCount int, input string) string {
	lines := getLines(input)
	for i, line := range lines {
		lines[i] = strings.Repeat("\n", tabCount) + line
	}
	return asString(lines)
}

// Doc: `getLines` returns the individual lines of a multiline string.
func getLinesFunc(input string) []string {
	return getLines(input)
}

func getLines(input string) []string {
	var lines []string
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

// Doc: `asString` returns a string separated by newline characters from a string slice.
func asString(lines []string) string {
	return strings.Join(lines, "\n")
}

// Doc: `regexpReplaceLine` returns a string from a multiline string where the regexp given is executed on every single line and the replacement executed if there was one or more matches.
func regexpReplaceLineFunc(regExp, replacement, input string) string {
	lines := getLines(input)
	r := regexp.MustCompile(regExp)

	for i, line := range lines {
		lines[i] = r.ReplaceAllString(line, replacement)
	}

	return asString(lines)
}

// Doc: `filterLinesPrefix` returns only lines that have the given prefix
func filterLinesPrefixFunc(prefix, input string) string {
	lines := getLines(input)

	result := make([]string, 0, len(lines)/2)
	for _, line := range lines {
		if strings.HasPrefix(line, prefix) {
			result = append(result, line)
		}
	}

	return asString(result)
}

// Doc: `withoutLinesPrefix` removes all lines that have the given prefix
func withoutLinesPrefixFunc(prefix, input string) string {
	lines := getLines(input)

	result := make([]string, 0, len(lines)/2)
	for _, line := range lines {
		if !strings.HasPrefix(line, prefix) {
			result = append(result, line)
		}
	}

	return asString(result)
}
