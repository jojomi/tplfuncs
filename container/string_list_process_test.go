package container

import (
	"reflect"
	"testing"
)

func TestStringList_WrapAll(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		before string
		after  string
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
				before: "'",
				after:  "'",
			},
			want: []string{"'abc'", "'def'"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.WrapAll(tt.args.before, tt.args.after).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_Map(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		mapper func(elem string) string
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
				mapper: func(elem string) string {
					return elem[0:1]
				},
			},
			want: []string{"a", "d"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.Map(tt.args.mapper).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_ReplaceAll(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		newValue string
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
				newValue: "-$0-$0-",
			},
			want: []string{"-abc-abc-", "-def-def-"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.ReplaceAll(tt.args.newValue).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReplaceAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_TrimAll(t *testing.T) {
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
				strings: []string{"  abc\n", "\tdef "},
			},
			want: []string{"abc", "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.TrimAll().All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TrimAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_UnindentAll(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		count int
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
				strings: []string{"  abc", "    def"},
			},
			args: args{
				count: 2,
			},
			want: []string{"abc", "  def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.UnindentAll(tt.args.count).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnindentAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_IndentSpaceAll(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		count int
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
				1,
			},
			want: []string{" abc", " def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.IndentSpaceAll(tt.args.count).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndentSpaceAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_IndentTabAll(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		count int
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
				strings: []string{"abc", "\tdef"},
			},
			args: args{
				3,
			},
			want: []string{"\t\t\tabc", "\t\t\t\tdef"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.IndentTabAll(tt.args.count).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IndentTabAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
