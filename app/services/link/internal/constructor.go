package internal

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/ic2hrmk/links123/app/services/link/persistence/repository"
	"github.com/pkg/errors"
	"log"
	"net"

	linkPb "github.com/ic2hrmk/links123/app/services/link/pb/link"

	"github.com/ic2hrmk/links123/app"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type linkDomainService struct {
	configurations *linkDomainServiceConfigurations
	linkRepository repository.LinkRepository
}

func NewLinkDomainService(
	configurations *linkDomainServiceConfigurations,
	linkRepository repository.LinkRepository,
) app.MicroService {
	return &linkDomainService{
		configurations: configurations,
		linkRepository: linkRepository,
	}
}

func (rcv *linkDomainService) Serve() error {
	return rcv.serveUnsecured()
}

func (rcv *linkDomainService) serveUnsecured() error {
	lis, err := net.Listen("tcp", rcv.configurations.serveAddress)
	if err != nil {
		return errors.Errorf("failed to acquire address [%s]: %s",
			rcv.configurations.serveAddress, err)
	}

	s := grpc.NewServer()
	linkPb.RegisterLinkDomainServiceServer(s, rcv)

	reflection.Register(s)

	log.Printf("service is available via [unsecure grpc] at address [%s]",
		rcv.configurations.serveAddress)

	if err := s.Serve(lis); err != nil {
		return errors.Errorf("failed to serve: %v", err)
	}

	return nil
}

//
// =============== Additional configuration section ================
//

//
// Configuration container
//
type linkDomainServiceConfigurations struct {
	mongoURL     string
	serveAddress string
}

//
// Configuration container builder
//
type linkDomainGatewayConfigBuilder struct {
	mongoURL     string
	serveAddress string
}

func NewLinkDomainServiceConfigBuilder() *linkDomainGatewayConfigBuilder {
	return &linkDomainGatewayConfigBuilder{}
}

func (rcv *linkDomainGatewayConfigBuilder) SetServeAddress(address string) *linkDomainGatewayConfigBuilder {
	rcv.serveAddress = address
	return rcv
}

func (rcv *linkDomainGatewayConfigBuilder) SetMongoURL(address string) *linkDomainGatewayConfigBuilder {
	rcv.mongoURL = address
	return rcv
}

//
// Builds gateway configuration object. It's the only way to initialize it's settings
//
func (rcv *linkDomainGatewayConfigBuilder) Build() (*linkDomainServiceConfigurations, error) {
	if err := rcv.Validate(); err != nil {
		return nil, errors.Wrap(err, "configuration is invalid")
	}

	return &linkDomainServiceConfigurations{
		mongoURL:     rcv.mongoURL,
		serveAddress: rcv.serveAddress,
	}, nil
}

//
// Validates acquired settings
//
func (rcv *linkDomainGatewayConfigBuilder) Validate() error {
	return validation.ValidateStruct(rcv,
		validation.Field(&rcv.mongoURL, validation.Required),
		validation.Field(&rcv.serveAddress, validation.Required),
	)
}
