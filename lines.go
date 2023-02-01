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

// LineHelpers returns a text template FuncMap with functions related to line processing
func LineHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"getLines":            getLines,
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
	}
}

// LineHelpersHTML returns an HTML template FuncMap with functions related to line processing
func LineHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LineHelpers())
}

func lineOrErrFunc(number int, input string) (string, error) {
	lines := getLines(input)
	if number < 1 || number > len(lines) {
		return "", fmt.Errorf("line ")
	}
	return lines[number-1], nil
}

// 1-based
func lineFunc(number int, input string) string {
	lines := getLines(input)
	if number < 1 || number > len(lines) {
		return ""
	}
	return lines[number-1]
}

func sortLinesFunc(input string) string {
	lines := getLines(input)
	sort.Strings(lines)
	return asString(lines)
}

func headFunc(count int, input string) string {
	lines := getLines(input)
	if count < len(lines) {
		lines = lines[:count]
	}
	return asString(lines)
}

func skipHeadFunc(count int, input string) string {
	lines := getLines(input)
	if count < len(lines) {
		lines = lines[count:]
	}
	return asString(lines)
}

func tailFunc(count int, input string) string {
	lines := getLines(input)
	if count < len(lines) {
		lines = lines[len(lines)-count:]
	}
	return asString(lines)
}

func skipTailFunc(count int, input string) string {
	lines := getLines(input)
	if count < len(lines) {
		lines = lines[0 : len(lines)-count]
	}
	return asString(lines)
}

// remove leading and trailing empty lines
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

func trimAllFunc(input string) string {
	lines := getLines(input)
	for i, l := range lines {
		lines[i] = strings.TrimSpace(l)
	}
	return asString(lines)
}

func notMatchFunc(regExp, input string) string {
	lines := getLines(input)
	r := regexp.MustCompile(regExp)
	var result = make([]string, 0)

	for _, line := range lines {
		if r.MatchString(line) {
			continue
		}
		result = append(result, line)
	}

	return asString(result)
}

func matchFunc(regExp, input string) string {
	lines := getLines(input)
	r := regexp.MustCompile(regExp)
	var result = make([]string, 0)

	for _, line := range lines {
		if !r.MatchString(line) {
			continue
		}
		result = append(result, line)
	}

	return asString(result)
}

func withoutEmptyLinesFunc(input string) string {
	return notMatchFunc(`^\s*$`, input)
}

func withoutLineCommentsFunc(input string) string {
	return notMatchFunc(`^\s*(//|#)`, input)
}

func wrapLinesFunc(leading, trailing, input string) string {
	lines := getLines(input)
	for i, line := range lines {
		lines[i] = leading + line + trailing
	}
	return asString(lines)
}

func indentSpaceLinesFunc(spaceCount int, input string) string {
	lines := getLines(input)
	for i, line := range lines {
		lines[i] = strings.Repeat(" ", spaceCount) + line
	}
	return asString(lines)
}

func indentTabLinesFunc(tabCount int, input string) string {
	lines := getLines(input)
	for i, line := range lines {
		lines[i] = strings.Repeat("\n", tabCount) + line
	}
	return asString(lines)
}

func getLines(input string) []string {
	var lines []string
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func asString(lines []string) string {
	return strings.Join(lines, "\n")
}

func regexpReplaceLineFunc(regExp, replacement, input string) string {
	lines := getLines(input)
	r := regexp.MustCompile(regExp)

	for i, line := range lines {
		lines[i] = r.ReplaceAllString(line, replacement)
	}

	return asString(lines)
}
