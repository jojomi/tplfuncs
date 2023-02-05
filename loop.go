package tplfuncs

import (
	htmlTemplate "html/template"
	textTemplate "text/template"
)

// LoopHelpers returns a text template FuncMap with loop functions
func LoopHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"for": forFunction,
	}
}

// LoopHelpersHTML returns an HTML template FuncMap with loop functions
func LoopHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LoopHelpers())
}

// forFunction can be used like this: {{ range $i := for 0 14 }}
func forFunction(from, to int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := from; i <= to; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
