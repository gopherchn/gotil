package factory

import (
	"reflect"
	"testing"
)

func TestAbstractJSONConfigParserFactory_CreateSystemParser(t *testing.T) {
	tests := []struct {
		name string
		a    *AbstractJSONConfigParserFactory
		want SystemConfigParser
	}{
		{
			name: "system",
			a: &AbstractJSONConfigParserFactory{},
			want: &JSONSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractJSONConfigParserFactory{}
			if got := a.CreateSystemParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AbstractJSONConfigParserFactory.CreateSystemParser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbstractJSONConfigParserFactory_CreateConfigParser(t *testing.T) {
	tests := []struct {
		name string
		a    *AbstractJSONConfigParserFactory
		want ConfigParser
	}{

		{
			name: "json",
			a: &AbstractJSONConfigParserFactory{},
			want: &JSONConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AbstractJSONConfigParserFactory{}
			if got := a.CreateConfigParser(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AbstractJSONConfigParserFactory.CreateConfigParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
