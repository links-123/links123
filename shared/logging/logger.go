package logging

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	promtailHook "github.com/ic2hrmk/lokigrus"
	microConfig "github.com/micro/go-micro/v2/config"

	"github.com/links-123/links123/shared/configuration"
)

func NewLogger(config microConfig.Config) (*logrus.Logger, error) {
	//
	// Extract configurations
	//
	logLevel, err := getLogLevel(config)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get log level from configurations")
	}

	logFormat := &logrus.TextFormatter{
		DisableColors:   true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	//
	// Init. basic logger
	//
	logger := logrus.StandardLogger()
	logger.SetLevel(logLevel)
	logger.SetFormatter(logFormat)

	//
	// Init. on-demand logger hooks
	//

	if isPromtailHookEnabled(config) {
		promtailHooker, err := initPromtailHook(config)
		if err != nil {
			return nil, errors.Wrap(err, "failed to initialize Promtail hook")
		}

		logger.AddHook(promtailHooker)
		logger.Debug("Promtail hook is enabled!")
	}

	return logger, nil
}

func getLogLevel(config microConfig.Config) (logrus.Level, error) {
	return logrus.ParseLevel(
		config.Get("link", "log", "level").String(logrus.InfoLevel.String()))
}

//
// Validates is Promtail client enabled via configuration
// (env. variable with default value - its parsed alternative)
// LINK_LOG_LOKI_ENABLED=false
//
func isPromtailHookEnabled(config microConfig.Config) bool {
	return config.Get("link", "log", "loki", "enabled").Bool(false)
}

//
// Initializes Promtail client for log submission to Loki server.
// Consumes configurations (env. variables with default values - their parsed alternatives):
//	LINK_LOG_LOKI_URL=unknown
//
func initPromtailHook(config microConfig.Config) (logrus.Hook, error) {
	var (
		lokiServerURL = config.Get("link", "log", "loki", "url").String("unknown")

		// Labels
		serviceKindLabel       = config.Get(configuration.KindFlagName).String("unknown")
		serviceInstanceIDLabel = config.Get(configuration.InstanceID).String("unknown")
	)

	return promtailHook.NewPromtailHook(
		lokiServerURL,
		map[string]string{
			"namespace":   "links-123.backend",
			"serviceKind": serviceKindLabel,
			"instanceId":  serviceInstanceIDLabel,
		},
	)
}
