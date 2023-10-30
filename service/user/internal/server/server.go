package server

import (
	"user/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"

	consulApi "github.com/hashicorp/consul/api"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	c := consulApi.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme

	cli, err := consulApi.NewClient(c)
	if err != nil {
		panic(err)
	}

	r := consul.New(cli, consul.WithHealthCheck(false))

	return r
}
