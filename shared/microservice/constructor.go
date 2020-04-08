package microservice

import (
	"net/http"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/web"

	"github.com/links-123/links123/shared/configuration"
	"github.com/links-123/links123/shared/version"
)

func getBuiltinFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  configuration.KindFlagName,
			Usage: configuration.KindFlagUsage,
		},
		&cli.BoolFlag{
			Name:  configuration.VersionFlagName,
			Usage: configuration.VersionFlagUsage,
		},
		&cli.StringFlag{
			Name:  configuration.EnvFlagName,
			Usage: configuration.EnvFlagUsage,
		},
	}
}

func NewBackendService(name string) micro.Service {
	microService := micro.NewService(
		micro.Name(name),
		micro.Version(version.Version),
		micro.Flags(getBuiltinFlags()...),
	)

	microService.Init()

	return microService
}

func NewWebService(name string, address string, handler http.Handler) (web.Service, error) {
	service := web.NewService(
		web.Name(name),
		web.Version(version.Version),
		web.Address(address),
		web.Handler(handler),
		web.Flags(getBuiltinFlags()...),
	)
	if err := service.Init(); err != nil {
		return nil, err
	}

	return service, nil
}
