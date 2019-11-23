package config

import (
	"github.com/pkg/errors"

	"github.com/ic2hrmk/links123/shared/env"
)

//
// All env. vars that are used by service
//
const (
	mongoURL           = "MONGO_URL"
	serviceLinkAddress = "SERVICE_LINK_ADDRESS"
)

//
// Reads all configuration from env, db etc.
//
func ResolveConfigurations() (*ConfigurationContainer, error) {
	c := &ConfigurationContainer{}

	if err := resolveEnvConfiguration(c); err != nil {
		return nil, err
	}

	return c, nil
}

//
// Reads all variables, stored in env. file
//
func resolveEnvConfiguration(config *ConfigurationContainer) error {
	var err error

	if config == nil {
		config = &ConfigurationContainer{}
	}

	if config.MongoURL, err = env.GetStringVar(mongoURL); err != nil {
		return errors.Errorf("failed to read env. var [%s]: %s", mongoURL, err)
	}

	if config.ServeAddress, err = env.GetStringVar(serviceLinkAddress); err != nil {
		return errors.Errorf("failed to read env. var [%s]: %s", serviceLinkAddress, err)
	}

	return nil
}
