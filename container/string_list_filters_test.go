package container

import (
	"reflect"
	"strings"
	"testing"
)

func TestStringList_Filter(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		keep func(elem string) (keep bool)
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"abc", "ade", "dfg"},
			},
			args: args{
				keep: func(elem string) (keep bool) {
					return strings.HasPrefix(elem, "a")
				},
			},
			want: []string{"abc", "ade"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Filter(tt.args.keep).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_FilterRegexp(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		regExp string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"abc", "bbc", "cbc"},
			},
			args: args{
				regExp: `[a|c]bc$`,
			},
			want: []string{"abc", "cbc"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.FilterRegexp(tt.args.regExp).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_FilterContains(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		filter string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"abc", "def", "add"},
			},
			args: args{
				filter: "a",
			},
			want: []string{"abc", "add"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.FilterContains(tt.args.filter).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_FilterContainsRegexp(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		regExp string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"abc", "def", "ghi"},
			},
			args: args{
				regExp: `[d|i]`,
			},
			want: []string{"def", "ghi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.FilterContainsRegexp(tt.args.regExp).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterContainsRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_WithoutEmpty(t *testing.T) {
	type fields struct {
		strings []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"abc", "   ", "def", "\t", ""},
			},
			want: []string{"abc", "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.WithoutEmpty().All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithoutEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_WithoutLineComments(t *testing.T) {
	type fields struct {
		strings []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"// abc", "def // keep this", "# see this", "  # or this", "- # but keep that one"},
			},
			want: []string{"def // keep this", "- # but keep that one"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.WithoutLineComments().All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithoutLineComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_WithoutEmptyStartEnd(t *testing.T) {
	type fields struct {
		strings []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"", "abc", "", "def", "", "   "},
			},
			want: []string{"abc", "", "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.WithoutEmptyStartEnd().All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithoutEmptyStartEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
