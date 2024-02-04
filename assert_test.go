package tplfuncs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_assertStringFunc(t *testing.T) {
	type args struct {
		msg   string
		input interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "is a string",
			args: args{
				input: "yes, it is",
			},
			wantErr: assert.NoError,
		},
		{
			name: "is not a string",
			args: args{
				input: 124,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.wantErr(t, assertStringFunc(tt.args.input), fmt.Sprintf("assertStringFunc(%v, %v)", tt.args.msg, tt.args.input))
		})
	}
}
