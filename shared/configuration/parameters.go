package configuration

import "strings"

// Environment variables (shared for all microservices)
const (
	MongoURLEnvName         = "MONGO_URL"
	MongoURLEnvDefaultValue = ":27017"
)

// Flags (shared for all microservices)
const (
	KindFlagName         = "kind"
	KindFlagDefaultValue = "undefined"
	KindFlagUsage        = "Type of microservice which we have to run"

	EnvFlagName         = "env"
	EnvFlagDefaultValue = ""
	EnvFlagUsage        = "Path to file with environment variables"

	VersionFlagName         = "version"
	VersionFlagDefaultValue = false
	VersionFlagUsage        = "Show application version"
)

// Common tags, generated at boot-up
const (
	InstanceID = "instanceId"
)

func EnvVarAsConfigurationPath(name string) []string {
	return strings.Split(strings.ToLower(name), "_")
}
