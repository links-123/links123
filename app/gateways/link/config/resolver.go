package config

import (
	"fmt"

	"github.com/ic2hrmk/links123/shared/env"
)

//
// All env. vars that are used by service
//
const (
	linkServiceAddress = "SERVICE_LINK_ADDRESS"
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

	if config.LinkDomainServiceAddress, err = env.GetStringVar(linkServiceAddress); err != nil {
		return fmt.Errorf("failed to read env. var [%s]: %s", linkServiceAddress, err)
	}

	return nil
}
