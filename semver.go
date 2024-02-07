package tplfuncs

import (
	"fmt"
	htmlTemplate "html/template"
	textTemplate "text/template"

	"github.com/Masterminds/semver"
	"github.com/juju/errors"
)

// SemverHelpers returns a text template FuncMap with semver related functions
func SemverHelpers() textTemplate.FuncMap {
	return textTemplate.FuncMap{
		"parseSemver":   parseSemverFunc,
		"semverToMinor": semverToMinorFunc,
		"semverToMajor": semverToMajorFunc,
	}
}

// SemverHelpersHTML returns an HTML template FuncMap with semver related functions
func SemverHelpersHTML() htmlTemplate.FuncMap {
	return htmlTemplate.FuncMap(SemverHelpers())
}

// Doc: `parseSemver` converts a string to a *semver.Version.
func parseSemverFunc(semverString string) (*semver.Version, error) {
	s, err := semver.NewVersion(semverString)
	if err != nil {
		return nil, errors.Annotatef(err, "invalid version: '%s'", semverString)
	}
	return s, nil
}

// Doc: `semverToMajor` converts a string to the major version part of a *semver.Version.
func semverToMajorFunc(semverString string) (string, error) {
	s, err := semver.NewVersion(semverString)
	if err != nil {
		return "", errors.Annotatef(err, "invalid version: '%s'", semverString)
	}
	return fmt.Sprintf("%d", s.Major()), nil
}

// Doc: `semverToMinor` converts a string to the major.minor version part of a *semver.Version.
func semverToMinorFunc(semverString string) (string, error) {
	s, err := semver.NewVersion(semverString)
	if err != nil {
		return "", errors.Annotatef(err, "invalid version: '%s'", semverString)
	}
	return fmt.Sprintf("%d.%d", s.Major(), s.Minor()), nil
}
