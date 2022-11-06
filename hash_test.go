package tplfuncs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sha1Func(t *testing.T) {
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
				input: "my secret data",
			},
			want: "9300dfe02f230d49a5f06d12254f1eca01b3ff06",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sha1Func(tt.args.input), "sha1Func(%v)", tt.args.input)
		})
	}
}

func Test_sha256Func(t *testing.T) {
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
				input: "my secret data",
			},
			want: "26e70c20d95e066c4d1af0fa051843dcf4497ac9df3922e722f13211e0798649",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sha256Func(tt.args.input), "sha256Func(%v)", tt.args.input)
		})
	}
}
