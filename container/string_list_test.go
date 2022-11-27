package container

import (
	"reflect"
	"testing"
)

func TestNewStringList(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "basic",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringList(); !got.Empty() {
				t.Errorf("NewStringList() = %v, want empty StringList", got)
			}
		})
	}
}

func TestNewStringListFromList(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want *StringList
	}{
		{
			name: "basic",
			args: args{
				input: []string{"abc", "def"},
			},
			want: NewStringList().AddAll("abc", "def"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringListFromList(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringListFromList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_Add(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		input string
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
				strings: []string{"abc"},
			},
			args: args{
				"def",
			},
			want: []string{"abc", "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Add(tt.args.input).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_AddAll(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		input []string
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
				strings: []string{},
			},
			args: args{
				input: []string{"abc", "def"},
			},
			want: []string{"abc", "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.AddAll(tt.args.input...).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_All(t *testing.T) {
	type fields struct {
		strings []string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "empty",
			fields: fields{
				strings: []string{},
			},
			want: []string{},
		},
		{
			name: "non-empty",
			fields: fields{
				strings: []string{"abc", "def"},
			},
			want: []string{"abc", "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_AsText(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		delim     string
		twoDelim  string
		lastDelim string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "one elem",
			fields: fields{
				strings: []string{"a"},
			},
			args: args{
				delim:     ", ",
				twoDelim:  " and ",
				lastDelim: ", and ",
			},
			want: "a",
		},
		{
			name: "three elems",
			fields: fields{
				strings: []string{"a", "b"},
			},
			args: args{
				delim:     ", ",
				twoDelim:  " and ",
				lastDelim: ", and ",
			},
			want: "a and b",
		},
		{
			name: "three elems",
			fields: fields{
				strings: []string{"a", "b", "c"},
			},
			args: args{
				delim:     ", ",
				twoDelim:  " and ",
				lastDelim: ", and ",
			},
			want: "a, b, and c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.AsText(tt.args.delim, tt.args.twoDelim, tt.args.lastDelim); got != tt.want {
				t.Errorf("AsText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_Clear(t *testing.T) {
	type fields struct {
		strings []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "clear",
			fields: fields{
				strings: []string{"abc"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Clear(); !reflect.DeepEqual(got.Len(), tt.want) {
				t.Errorf("Clear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_Has(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		query string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "contained",
			fields: fields{
				strings: []string{"abc", "def", "ghi"},
			},
			args: args{"def"},
			want: true,
		},
		{
			name: "not contained",
			fields: fields{
				strings: []string{"abc", "def", "ghi"},
			},
			args: args{"uvw"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Has(tt.args.query); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_Joined(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		delim string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"abc", "def"},
			},
			args: args{
				delim: "|",
			},
			want: "abc|def",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Joined(tt.args.delim); got != tt.want {
				t.Errorf("Joined() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_Len(t *testing.T) {
	type fields struct {
		strings []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "empty",
			fields: fields{
				strings: []string{},
			},
			want: 0,
		},
		{
			name: "filled",
			fields: fields{
				strings: []string{"abc"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_Remove(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		input string
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
				strings: []string{"abc", "def"},
			},
			args: args{
				input: "abc",
			},
			want: []string{"def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Remove(tt.args.input).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_RemoveAll(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		input []string
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
				input: []string{"def", "abc"},
			},
			want: []string{"ghi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.RemoveAll(tt.args.input...).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_RemoveList(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		input *StringList
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
				input: NewStringList().AddAll("def", "abc"),
			},
			want: []string{"ghi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.RemoveList(tt.args.input).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_Empty(t *testing.T) {
	type fields struct {
		strings []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "empty",
			fields: fields{
				strings: []string{},
			},
			want: true,
		},
		{
			name: "non-empty",
			fields: fields{
				strings: []string{"abc"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_AddList(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		input []string
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
				strings: []string{"abc"},
			},
			args: args{
				input: []string{"def", "ghi"},
			},
			want: []string{"abc", "def", "ghi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.AddList(tt.args.input).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewStringListFromMultilineString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "basic",
			args: args{
				input: "first\nand second\nand third\n",
			},
			want: []string{"first", "and second", "and third", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStringListFromMultilineString(tt.args.input).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStringListFromMultilineString() = %v, want %v", got, tt.want)
			}
		})
	}
}
