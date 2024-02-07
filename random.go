package tplfuncs

import (
	htmlTemplate "html/template"
	"math/rand"
	textTemplate "text/template"
	"time"
)

// RandomHelpers returns a text template FuncMap with random related functions
func RandomHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"seededRandom": seededRandomFunc,
		"random":       randomFunc,
	}
}

// RandomHelpersHTML returns an HTML template FuncMap with random related functions
func RandomHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(RandomHelpers())
}

// Doc: `seededRandom` returns a rand.Source that is seeded with the given int value.
func seededRandomFunc(seed int) *rand.Rand {
	s := rand.NewSource(int64(seed))
	return rand.New(s)
}

// Doc: `random` returns a rand.Source that is seeded with the current time.
func randomFunc() *rand.Rand {
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s)
}
