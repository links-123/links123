package link_rest_gtw

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/micro/go-micro/v2/config"
	"github.com/pkg/errors"

	"github.com/links-123/links123/shared/configuration"
)

//
// All env. vars that are used by service
//
const (
	restGatewayAddressParameter    = "LINK_REST_GATEWAY_ADDRESS"
	restGatewayAddressDefaultValue = ":80"
)

//
// Configuration container
//
type gatewayConfiguration struct {
	instanceID   string
	serveAddress string
}

func (rcv *gatewayConfiguration) Validate() error {
	return validation.ValidateStruct(rcv,
		validation.Field(&rcv.serveAddress, validation.Required),
	)
}

func ResolveConfigurations(configSource config.Config) (*gatewayConfiguration, error) {
	builder := newGatewayConfigurationBuilder().
		SetInstanceID(configSource.
			Get(configuration.InstanceID).
			String("unknown"),
		).
		SetServeAddress(configSource.
			Get(configuration.EnvVarAsConfigurationPath(restGatewayAddressParameter)...).
			String(restGatewayAddressDefaultValue),
		)

	return builder.Build()
}

type gatewayConfigurationBuilder struct {
	instanceID   string
	serveAddress string
}

func newGatewayConfigurationBuilder() *gatewayConfigurationBuilder {
	return &gatewayConfigurationBuilder{
		serveAddress: restGatewayAddressDefaultValue,
	}
}

func (rcv *gatewayConfigurationBuilder) SetInstanceID(instanceID string) *gatewayConfigurationBuilder {
	rcv.instanceID = instanceID
	return rcv
}

func (rcv *gatewayConfigurationBuilder) SetServeAddress(address string) *gatewayConfigurationBuilder {
	rcv.serveAddress = address
	return rcv
}

func (rcv *gatewayConfigurationBuilder) Build() (*gatewayConfiguration, error) {
	conf := &gatewayConfiguration{
		instanceID:   rcv.instanceID,
		serveAddress: rcv.serveAddress,
	}

	if err := conf.Validate(); err != nil {
		return nil, errors.Wrap(err, "configuration is invalid")
	}

	return conf, nil
}
