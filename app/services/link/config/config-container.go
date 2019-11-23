package config

import (
	"github.com/go-ozzo/ozzo-validation"
)

//
// All available configurations for the micro-service
//
type ConfigurationContainer struct {
	MongoURL     string
	ServeAddress string
}

func (c *ConfigurationContainer) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.MongoURL, validation.Required),
		validation.Field(&c.ServeAddress, validation.Required),
	)
}
