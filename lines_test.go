package tplfuncs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sortLinesFunc(t *testing.T) {
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
				input: "a\nc\nd\nb",
			},
			want: "a\nb\nc\nd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, sortLinesFunc(tt.args.input), "sortLinesFunc(%v)", tt.args.input)
		})
	}
}
