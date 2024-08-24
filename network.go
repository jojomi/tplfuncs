package tplfuncs

import (
	htmlTemplate "html/template"
	"io"
	"net/http"
	"os"
	textTemplate "text/template"
)

// NetworkHelpers returns a text template FuncMap with network functions
func NetworkHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"download":   downloadFunc,
		"includeUrl": includeUrlFunc,
	}
}

// NetworkHelpersHTML returns an HTML template FuncMap with network functions
func NetworkHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(LinesHelpers())
}

// Doc: `download` executes an HTTP GET request to a given URL and stores the result to a file.
func downloadFunc(srcURL, filename string) error {
	body, err := includeUrlFunc(srcURL)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, []byte(body), 0x640)
}

// Doc: `includeUrl` executes an HTTP GET request to a given URL and returns the result.
func includeUrlFunc(srcURL string) (string, error) {
	resp, err := http.Get(srcURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
