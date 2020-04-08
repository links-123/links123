package registry

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/pkg/errors"

	"github.com/links-123/links123/app"
)

type RegistryContainer map[string]FactoryMethod

func NewRegistryContainer() RegistryContainer {
	return make(RegistryContainer)
}

type FactoryMethod func(container config.Config) (app.MicroService, error)

func (c RegistryContainer) Add(serviceName string, fabric FactoryMethod) {
	c[serviceName] = fabric
}

func (c RegistryContainer) Get(serviceName string) (FactoryMethod, error) {
	entry, ok := c[serviceName]
	if !ok {
		return nil, errors.Errorf("service [%s] isn't registered", serviceName)
	}

	return entry, nil
}
