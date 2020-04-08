package configuration

import "strings"

const (
	MongoURLEnvName         = "MONGO_URL"
	MongoURLEnvDefaultValue = ":27017"
)

const (
	KindFlagName         = "kind"
	KindFlagDefaultValue = "undefined"
	KindFlagUsage        = "Type of microservice which we have to run"

	VersionFlagName         = "version"
	VersionFlagDefaultValue = false
	VersionFlagUsage        = "Show application version"

	EnvFlagName         = "env"
	EnvFlagDefaultValue = ""
	EnvFlagUsage        = "Path to file with environment variables"
)

func EnvVarAsConfigurationPath(name string) []string {
	return strings.Split(strings.ToLower(name), "_")
}
