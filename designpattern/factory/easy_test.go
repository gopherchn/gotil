package factory

import (
	"reflect"
	"testing"
)

func TestNewConfigParser(t *testing.T) {
	type args struct {
		typ string
	}
	tests := []struct {
		name string
		args args
		want ConfigParser
	}{
		{
			name: "yaml",
			args: args{
				typ: "yaml",
			},
			want: &YamlConfigParser{},
		},
		{
			name: "json",
			args: args{
				typ: "json",
			},
			want: &JSONConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigParser(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
