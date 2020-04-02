package registry

import (
	"github.com/links-123/links123/app"
	"github.com/pkg/errors"
)

type RegistryContainer map[string]FactoryMethod

func NewRegistryContainer() RegistryContainer {
	return make(RegistryContainer)
}

type FactoryMethod func() (app.MicroService, error)

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
