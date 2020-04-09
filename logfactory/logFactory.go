// package loggerfactory handles creating concrete logger with factory method pattern
package logfactory

import (
	"github.com/jfeng45/glogger"
	"github.com/jfeng45/glogger/logconfig"
	"github.com/jfeng45/glogger/logrus"
	"github.com/jfeng45/glogger/zap"
	"github.com/pkg/errors"
)
// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)

// logger mapp to map logger code to logger builder
var logfactoryBuilderMap = map[string]logFbInterface{
	ZAP:    &zap.ZapFactory{},
	LOGRUS: &logrus.LogrusFactory{},
}

// interface for logger factory
type logFbInterface interface {
	Build(*logconfig.LogConfig) (glogger.Logger, error)
}

// accessors for factoryBuilderMap
func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}

func InitLogger(lc logconfig.LogConfig) (glogger.Logger, error) {
	loggerType := lc.Code
	l, err := GetLogFactoryBuilder(loggerType).Build(&lc)
	if err != nil {
		return l, errors.Wrap(err, "")
	}
	return l, nil
}
