package config

import (
	"github.com/go-ozzo/ozzo-validation"
)

//
// All available configurations for the micro-service
//
type ConfigurationContainer struct {
	//
	// Addresses of all connected services
	//
	LinkDomainServiceAddress string

	//
	// Internal service configurations
	//
	ServeAddress string
}

func (c *ConfigurationContainer) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.LinkDomainServiceAddress, validation.Required),
		validation.Field(&c.ServeAddress, validation.Required),
	)
}
