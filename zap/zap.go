// package zap handles creating zap logger
package zap

import (
	"encoding/json"
	"github.com/jfeng45/glogger"
	"github.com/jfeng45/glogger/logconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

//type loggerWrapper struct {
//	lw *zap.SugaredLogger
//}
//
//
//func (logger *loggerWrapper) Errorf(format string, args ...interface{}) {
//	logger.lw.Errorf(format, args)
//}
//func (logger *loggerWrapper) Fatalf(format string, args ...interface{}) {
//	logger.lw.Fatalf(format, args)
//}
//func (logger *loggerWrapper) Fatal(args ...interface{}) {
//	logger.lw.Fatal(args)
//}
//func (logger *loggerWrapper) Infof(format string, args ...interface{}) {
//	logger.lw.Infof(format, args)
//}
//func (logger *loggerWrapper) Warnf(format string, args ...interface{}) {
//	logger.lw.Warnf(format, args)
//}
//func (logger *loggerWrapper) Debugf(format string, args ...interface{}) {
//	logger.lw.Debugf(format, args)
//}
//func (logger *loggerWrapper) Printf(format string, args ...interface{}) {
//	logger.lw.Infof(format, args)
//}
//func (logger *loggerWrapper) Println(args ...interface{}) {
//	logger.lw.Info(args, "\n")
//}

func RegisterLog(lc logconfig.LogConfig) (glogger.Logger, error) {
	zLogger, err := initLog(lc)
	if err != nil {
		return nil, errors.Wrap(err, "RegisterLogrusLog")
	}
	defer zLogger.Sync()
	zSugarlog := zLogger.Sugar()
	zSugarlog.Info()

	//This is for loggerWrapper implementation
	//appLogger.SetLogger(&loggerWrapper{zaplog})

	//SetLogger(zSugarlog)
	return zSugarlog, nil
}

// initLog create logger
func initLog(lc logconfig.LogConfig) (zap.Logger, error) {
	rawJSON := []byte(`{
	 "level": "info",
     "Development": true,
      "DisableCaller": false,
	 "encoding": "console",
	 "outputPaths": ["stdout", "../../../demo.log"],
	 "errorOutputPaths": ["stderr"],
	 "encoderConfig": {
		"timeKey":        "ts",
		"levelKey":       "level",
		"messageKey":     "msg",
         "nameKey":        "name",
		"stacktraceKey":  "stacktrace",
         "callerKey":      "caller",
		"lineEnding":     "\n\t",
        "timeEncoder":     "time",
		"levelEncoder":    "lowercaseLevel",
        "durationEncoder": "stringDuration",
         "callerEncoder":   "shortCaller"
	 }
	}`)

	var cfg zap.Config
	var zLogger *zap.Logger
	//standard configuration
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		return *zLogger, errors.Wrap(err, "Unmarshal")
	}
	//customize it from configuration file
	err := customizeLogFromConfig(&cfg, lc)
	if err != nil {
		return *zLogger, errors.Wrap(err, "cfg.Build()")
	}
	zLogger, err = cfg.Build()
	if err != nil {
		return *zLogger, errors.Wrap(err, "cfg.Build()")
	}

	zLogger.Debug("logger construction succeeded")
	return *zLogger, nil
}

// customizeLogrusLogFromConfig customize log based on parameters from configuration file
func customizeLogFromConfig(cfg *zap.Config, lc logconfig.LogConfig) error {
	cfg.DisableCaller = !lc.EnableCaller

	// set log level
	l := zap.NewAtomicLevel().Level()
	err := l.Set(lc.Level)
	if err != nil {
		return errors.Wrap(err, "")
	}
	cfg.Level.SetLevel(l)

	return nil
}
