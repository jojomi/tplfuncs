package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Join(t *testing.T) {
	type args struct {
		delim     string
		twoDelim  string
		lastDelim string
		input     []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1 elem",
			args: args{
				delim:     ", ",
				twoDelim:  " and ",
				lastDelim: ", and ",
				input:     []string{"a"},
			},
			want: "a",
		},
		{
			name: "2 elems",
			args: args{
				delim:     ", ",
				twoDelim:  " and ",
				lastDelim: ", and ",
				input:     []string{"a", "b"},
			},
			want: "a and b",
		},
		{
			name: "3 elems",
			args: args{
				delim:     ", ",
				twoDelim:  " and ",
				lastDelim: ", and ",
				input:     []string{"a", "b", "c"},
			},
			want: "a, b, and c",
		},
		{
			name: "4 elems",
			args: args{
				delim:     ", ",
				twoDelim:  " and ",
				lastDelim: ", and ",
				input:     []string{"a", "b", "c", "d"},
			},
			want: "a, b, c, and d",
		},
		{
			name: "different delimiters",
			args: args{
				delim:     " + ",
				twoDelim:  ", then",
				lastDelim: " and finally ",
				input:     []string{"a", "b", "c", "d"},
			},
			want: "a + b + c and finally d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Join(tt.args.delim, tt.args.twoDelim, tt.args.lastDelim, tt.args.input), "joinTextFunc(%v, %v, %v, %v)", tt.args.delim, tt.args.twoDelim, tt.args.lastDelim, tt.args.input)
		})
	}
}
