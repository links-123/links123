package configuration

import (
	"flag"

	microConfig "github.com/micro/go-micro/v2/config"
	microEnv "github.com/micro/go-micro/v2/config/source/env"
	microFlag "github.com/micro/go-micro/v2/config/source/flag"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func init() {
	flag.String(KindFlagName, KindFlagDefaultValue, KindFlagUsage)               // to know which service we have to start
	flag.Bool(VersionFlagName, VersionFlagDefaultValue, VersionFlagUsage)        // path for the external file with env. variables
	flag.StringVar(&envFilePath, EnvFlagName, EnvFlagDefaultValue, EnvFlagUsage) // to show only version of the app and shut it down
	flag.Parse()
}

var (
	envFilePath string
)

func LoadSharedConfigurations() (Behavior, microConfig.Config, error) {
	var (
		behavior Behavior
		conf     microConfig.Config
		err      error
	)

	//
	// Load external environment variables
	//
	if envFilePath != "" {
		if err = godotenv.Load(envFilePath); err != nil {
			return behavior, conf, errors.Wrap(err, "failed to load env. file")
		}
	}

	//
	// Load complete configurations
	//
	conf, err = microConfig.NewConfig(
		microConfig.WithSource(
			microFlag.NewSource(
				microFlag.IncludeUnset(true),
			),
		),
		microConfig.WithSource(
			microEnv.NewSource(),
		),
	)
	if err != nil {
		return behavior, conf, errors.Wrap(err, "failed to init. configurations container")
	}

	if err = conf.Load(); err != nil {
		return behavior, conf, errors.Wrap(err, "failed to load configurations container")
	}

	behavior.ServiceKind = conf.Get("kind").String("undefined")
	behavior.ShowVersionOnly = conf.Get("version").Bool(false)

	return behavior, conf, nil
}
