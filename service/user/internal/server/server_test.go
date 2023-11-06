package server

import (
	"testing"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/registry"
)

func TestNewRegistrar(t *testing.T) {
	type args struct {
		conf *conf.Registry
	}
	tests := []struct {
		name string
		args args
		want registry.Registrar
	}{
		{
			args: args{
				conf: &conf.Registry{
					Consul: &conf.Registry_Consul{
						Address: "127.0.0.1:8500",
						Scheme:  "http",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRegistrar(tt.args.conf); got == nil {
				t.Errorf("NewRegistrar() = %v", got)
			} else {
				t.Log("111", got)
			}
		})
	}
}
