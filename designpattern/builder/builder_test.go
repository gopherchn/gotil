package builder

import (
	"reflect"
	"testing"
)

func TestResourcePoolConfigBuilder_Build(t *testing.T) {
	type fields struct {
		name     string
		maxTotal int
		maxIdle  int
		minIdle  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "test1",
			fields: fields{
				name:     "hi",
				maxTotal: 10,
			},
			want: &ResourcePoolConfig{
				name:     "hi",
				maxTotal: 10,
				maxIdle:  defaultMaxIdle,
				minIdle:  defaultMinIdle,
			},
		},
		{
			name: "test2",
			fields: fields{
				name:     "test2",
				maxTotal: 19,
				maxIdle:  1,
			},
			want: &ResourcePoolConfig{
				name:     "test2",
				maxTotal: 19,
				maxIdle:  1,
				minIdle:  defaultMinIdle,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ResourcePoolConfigBuilder{
				ResourcePoolConfig: ResourcePoolConfig{
					name:     tt.fields.name,
					maxTotal: tt.fields.maxTotal,
					maxIdle:  tt.fields.maxIdle,
					minIdle:  tt.fields.minIdle,
				},
			}
			t.Log("c: ", c)
			got, err := c.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourcePoolConfigBuilder.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResourcePoolConfigBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
