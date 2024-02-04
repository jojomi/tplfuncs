package tplfuncs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_envEqFunc(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				key:   "TEST_ENV",
				value: "true",
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				key:   "TEST_ENV",
				value: "checkValue",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		os.Setenv("TEST_ENV", "true")

		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, envEqFunc(tt.args.key, tt.args.value), "envEqFunc(%v, %v)", tt.args.key, tt.args.value)
		})
	}
}

func Test_envFunc(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{
				key: "TEST_ENV",
			},
			want: "true",
		},
	}
	for _, tt := range tests {
		os.Setenv("TEST_ENV", "true")

		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, envFunc(tt.args.key), "envFunc(%v)", tt.args.key)
		})
	}
}

func Test_envIsSetFunc(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found",
			args: args{
				key: "TEST_ENV",
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				key: "NON_EXISTING_ENV",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		os.Setenv("TEST_ENV", "true")

		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, envIsSetFunc(tt.args.key), "envIsSetFunc(%v)", tt.args.key)
		})
	}
}
