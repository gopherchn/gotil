package builder

import (
	"reflect"
	"testing"
)

func TestNewResourcePoolConfig(t *testing.T) {
	type args struct {
		options []func(*ResourcePoolConfig)
	}
	tests := []struct {
		name    string
		args    args
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				options: []func(*ResourcePoolConfig){
					WithMaxIdle(8),
				},
			},
			want: &ResourcePoolConfig{
				maxIdle: 8,
				minIdle: defaultMinIdle,
				maxTotal: defaultMaxTotal,
				name: "",
			},

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResourcePoolConfig(tt.args.options...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewResourcePoolConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourcePoolConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
