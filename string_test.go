package tplfuncs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stringCamelCaseFunc(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				input: "Camel Case test: This is good! But it can't be used as is...",
			},
			want: "CamelCaseTestThisIsGoodButItCantBeUsedAsIs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringCamelCaseFunc(tt.args.input), "stringCamelCaseFunc(%v)", tt.args.input)
		})
	}
}

func Test_stringContainsFunc(t *testing.T) {
	type args struct {
		needle   string
		haystack string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "contained",
			args: args{
				needle:   "and",
				haystack: "milk and honey",
			},
			want: true,
		},
		{
			name: "not contained",
			args: args{
				needle:   "or",
				haystack: "milk and honey",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringContainsFunc(tt.args.needle, tt.args.haystack), "stringContainsFunc(%v, %v)", tt.args.needle, tt.args.haystack)
		})
	}
}

func Test_stringKebabFunc(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				input: "Kebab Case test: This is good! But it can't be used as is...",
			},
			want: "kebab-case-test:-this-is-good!-but-it-can't-be-used-as-is---",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringKebabFunc(tt.args.input), "stringKebabFunc(%v)", tt.args.input)
		})
	}
}

func Test_stringLowerCamelCaseFunc(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				input: "Lower Camel Case test: This is good! But it can't be used as is...",
			},
			want: "lowerCamelCaseTestThisIsGoodButItCantBeUsedAsIs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringLowerCamelCaseFunc(tt.args.input), "stringLowerCamelCaseFunc(%v)", tt.args.input)
		})
	}
}

func Test_stringSnakeFunc(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				input: "Snake Case test: This is good! But it can't be used as is...",
			},
			want: "snake_case_test:_this_is_good!_but_it_can't_be_used_as_is___",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringSnakeFunc(tt.args.input), "stringSnakeFunc(%v)", tt.args.input)
		})
	}
}

func Test_stringToFilenameFunc(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				input: "To filename test: This is good! But it can't be used as is...",
			},
			want: "to_filename_test_this_is_good_but_it_can_t_be_used_as_is",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringToFilenameFunc(tt.args.input), "stringToFilenameFunc(%v)", tt.args.input)
		})
	}
}

func Test_stringToURLFunc(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				input: "ToURL test: This is good! But it can't be used as is...",
			},
			want: "to-url-test-this-is-good-but-it-can-t-be-used-as-is",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringToURLFunc(tt.args.input), "stringToURLFunc(%v)", tt.args.input)
		})
	}
}

func Test_stringCleanFunc(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				input: "Clean test: This is good!   But it can't be used as is...",
			},
			want: "Clean_test_This_is_good_But_it_can_t_be_used_as_is",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, stringCleanFunc(tt.args.input), "stringCleanFunc(%v)", tt.args.input)
		})
	}
}
