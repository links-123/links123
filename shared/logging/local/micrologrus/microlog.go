package micrologrus

import (
	"context"
	"fmt"
	"os"
	"sync"

	microLogger "github.com/micro/go-micro/v2/logger"
	logrusLogger "github.com/sirupsen/logrus"
)

//
// Adapter of Logrus for GoMicro which has in-memory buffer for reading purposes
//
type microLogrus struct {
	logger logrusLogger.FieldLogger

	sync.RWMutex
	options microLogger.Options
}

//
// Creates a middleware for connecting Logrus logger with Go-Micro framework.
// NOTE: default log level is taken from Logrus instance
// NOTE: it could receive options, but not all of them are used such as at Go-Micro default logger (skipped: context, output)
//
func NewMicroLogrus(logger *logrusLogger.Logger, options ...microLogger.Option) (microLogger.Logger, error) {
	if logger == nil {
		return nil, fmt.Errorf("logrus logger is nil")
	}

	// Derive log level from logrus
	microLogLevel, err := logrusLevelToMicroLevel(logger.Level)
	if err != nil {
		return nil, err
	}

	microOptions := microLogger.Options{
		Level:   microLogLevel,
		Fields:  make(map[string]interface{}),
		Out:     os.Stderr,
		Context: context.Background(), // Yes, yes, context as a struct field,
		// nothing to do cause native go-micro logger implementation requires it :(
	}

	//
	// Build middleware
	//
	middleware := &microLogrus{
		logger:  logger,
		options: microOptions,
	}

	if err = middleware.Init(options...); err != nil {
		return nil, fmt.Errorf("failed to intialize options: %s", err.Error())
	}

	return middleware, nil
}

func (rcv *microLogrus) Init(options ...microLogger.Option) error {
	for _, o := range options {
		o(&rcv.options)
	}
	return nil
}

func (rcv *microLogrus) Options() microLogger.Options {
	return rcv.options
}

func (rcv *microLogrus) Fields(fields map[string]interface{}) microLogger.Logger {
	rcv.Lock()
	rcv.options.Fields = copyFields(fields)
	rcv.Unlock()
	return rcv
}

func (rcv *microLogrus) Log(level microLogger.Level, v ...interface{}) {
	if rcv.options.Level < level {
		return
	}

	rcv.RLock()
	defer rcv.RUnlock()

	logLevel, err := microLevelToLogrusLevel(level)
	if err != nil { // Threat unknown log level as a Print
		rcv.logger.
			WithFields(rcv.options.Fields).
			Print(v...)
		return
	}

	rcv.logger.
		WithFields(rcv.options.Fields).
		Log(logLevel, v...)
}

func (rcv *microLogrus) Logf(level microLogger.Level, format string, v ...interface{}) {
	if rcv.options.Level < level {
		return
	}

	rcv.RLock()
	defer rcv.RUnlock()

	logLevel, err := microLevelToLogrusLevel(level)
	if err != nil { // Threat unknown log level as a Print
		rcv.logger.
			WithFields(rcv.options.Fields).
			Printf(format, v...)
		return
	}

	rcv.logger.
		WithFields(rcv.options.Fields).
		Logf(logLevel, format, v...)
}

func (rcv *microLogrus) String() string {
	return "microlog"
}

func logrusLevelToMicroLevel(logrusLvl logrusLogger.Level) (microLogger.Level, error) {
	var (
		microLvl microLogger.Level
		err      error
	)

	switch logrusLvl {
	case logrusLogger.TraceLevel:
		microLvl = microLogger.TraceLevel
	case logrusLogger.DebugLevel:
		microLvl = microLogger.DebugLevel
	case logrusLogger.InfoLevel:
		microLvl = microLogger.InfoLevel
	case logrusLogger.WarnLevel:
		microLvl = microLogger.WarnLevel
	case logrusLogger.ErrorLevel:
		microLvl = microLogger.ErrorLevel
	case logrusLogger.FatalLevel, logrusLogger.PanicLevel:
		microLvl = microLogger.FatalLevel
	default:
		err = fmt.Errorf("logrus level [%s] has no matching level with "+
			"go-micro log levels", logrusLvl.String())
	}

	if err != nil {
		return 0, err
	}

	return microLvl, nil
}

func microLevelToLogrusLevel(microLvl microLogger.Level) (logrusLogger.Level, error) {
	var (
		logrusLvl logrusLogger.Level
		err       error
	)

	switch microLvl {
	case microLogger.TraceLevel:
		logrusLvl = logrusLogger.TraceLevel
	case microLogger.DebugLevel:
		logrusLvl = logrusLogger.DebugLevel
	case microLogger.InfoLevel:
		logrusLvl = logrusLogger.InfoLevel
	case microLogger.WarnLevel:
		logrusLvl = logrusLogger.WarnLevel
	case microLogger.ErrorLevel:
		logrusLvl = logrusLogger.ErrorLevel
	case microLogger.FatalLevel:
		logrusLvl = logrusLogger.FatalLevel
	default:
		err = fmt.Errorf("go-micro level [%s] has no matching level with "+
			"logrus log levels", logrusLvl.String())
	}

	if err != nil {
		return 0, err
	}

	return logrusLvl, nil
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
