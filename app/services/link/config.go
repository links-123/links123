package link

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/pkg/errors"

	validation "github.com/go-ozzo/ozzo-validation"

	"github.com/links-123/links123/shared/configuration"
)

//
// All env. vars that are used by service
//
const (
	mongoURLParameter             = configuration.MongoURLEnvName
	mongoURLParameterDefaultValue = configuration.MongoURLEnvDefaultValue
)

//
// Configuration container
//
type serviceConfiguration struct {
	instanceID string
	mongoURL   string
}

func (rcv *serviceConfiguration) Validate() error {
	return validation.ValidateStruct(rcv,
		validation.Field(&rcv.mongoURL, validation.Required),
	)
}

func extractConfigurations(configSource config.Config) (*serviceConfiguration, error) {
	builder := newServiceConfigurationBuilder().
		SetInstanceID(configSource.
			Get(configuration.InstanceID).
			String(""),
		).
		SetMongoURL(configSource.
			Get(configuration.EnvVarAsConfigurationPath(mongoURLParameter)...).
			String(mongoURLParameterDefaultValue),
		)

	return builder.Build()
}

type serviceConfigurationBuilder struct {
	instanceID   string
	serveAddress string
}

func newServiceConfigurationBuilder() *serviceConfigurationBuilder {
	return &serviceConfigurationBuilder{
		serveAddress: mongoURLParameter,
	}
}

func (rcv *serviceConfigurationBuilder) SetInstanceID(id string) *serviceConfigurationBuilder {
	rcv.instanceID = id
	return rcv
}

func (rcv *serviceConfigurationBuilder) SetMongoURL(address string) *serviceConfigurationBuilder {
	rcv.serveAddress = address
	return rcv
}

func (rcv *serviceConfigurationBuilder) Build() (*serviceConfiguration, error) {
	conf := &serviceConfiguration{
		instanceID: rcv.instanceID,
		mongoURL:   rcv.serveAddress,
	}

	if err := conf.Validate(); err != nil {
		return nil, errors.Wrap(err, "configuration is invalid")
	}

	return conf, nil
}
