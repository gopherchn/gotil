package factory

import (
	"reflect"
	"testing"
)

func TestNewConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want ConfigParserFactory
	}{
		{
			name: "yaml factory",
			args: args{
				t: "yaml",
			},
			want: &YamlConfigParserFactory{},
		},
		{
			name: "json factory",
			args: args{
				t: "json",
			},
			want: &JSONConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigParserFactory(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigParserFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
