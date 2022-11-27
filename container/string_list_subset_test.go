package container

import (
	"reflect"
	"testing"
)

func TestStringList_First(t *testing.T) {
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
			want: []string{"abc"},
		},
		{
			name: "overflow",
			fields: fields{
				strings: []string{"abc", "def"},
			},
			args: args{
				3,
			},
			want: []string{"abc", "def"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			if got := x.First(tt.args.count).All(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringList_SubsetRegexp(t *testing.T) {
	type fields struct {
		strings []string
	}
	type args struct {
		startRegexp string
		endRegexp   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "basic",
			fields: fields{
				strings: []string{"abc", "def", "ghi", "jkl"},
			},
			args: args{
				startRegexp: "d.*f",
				endRegexp:   "g?hi$",
			},
			want:    []string{"def", "ghi"},
			wantErr: false,
		},
		{
			name: "start not found",
			fields: fields{
				strings: []string{"abc", "def", "ghi", "jkl"},
			},
			args: args{
				startRegexp: "nomatch",
				endRegexp:   "",
			},
			want:    []string{"abc", "def", "ghi", "jkl"},
			wantErr: true,
		},
		{
			name: "end not found",
			fields: fields{
				strings: []string{"abc", "def", "ghi", "jkl"},
			},
			args: args{
				startRegexp: "",
				endRegexp:   "nomatch",
			},
			want:    []string{"abc", "def", "ghi", "jkl"},
			wantErr: true,
		},
		{
			name: "everything",
			fields: fields{
				strings: []string{"abc", "def", "ghi", "jkl"},
			},
			args: args{
				startRegexp: "",
				endRegexp:   "",
			},
			want:    []string{"abc", "def", "ghi", "jkl"},
			wantErr: false,
		},
		{
			name: "open start",
			fields: fields{
				strings: []string{"abc", "def", "ghi", "jkl"},
			},
			args: args{
				startRegexp: "",
				endRegexp:   "def",
			},
			want:    []string{"abc", "def"},
			wantErr: false,
		},
		{
			name: "open end",
			fields: fields{
				strings: []string{"abc", "def", "ghi", "jkl"},
			},
			args: args{
				startRegexp: "gh.",
				endRegexp:   "",
			},
			want:    []string{"ghi", "jkl"},
			wantErr: false,
		},
		{
			name: "pick the right end",
			fields: fields{
				strings: []string{"abc", "def", "abc", "jkl"},
			},
			args: args{
				startRegexp: "def",
				endRegexp:   "abc",
			},
			want:    []string{"def", "abc"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &StringList{
				strings: tt.fields.strings,
			}
			got, err := x.SubsetRegexp(tt.args.startRegexp, tt.args.endRegexp)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubsetRegexp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.All(), tt.want) {
				t.Errorf("SubsetRegexp() got = %v, want %v", got, tt.want)
			}
		})
	}
}
