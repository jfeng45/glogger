// package factory handles creating concrete logger with factory method pattern
package factory

import (
	"github.com/jfeng45/glogger"
	"github.com/jfeng45/glogger/config"
	"github.com/jfeng45/glogger/logrus"
	"github.com/jfeng45/glogger/zap"
	"github.com/pkg/errors"
)

// logger mapp to map logger code to logger builder
var logfactoryBuilderMap = map[string]logFbInterface{
	config.ZAP:    &zap.ZapFactory{},
	config.LOGRUS: &logrus.LogrusFactory{},
}

// interface for logger factory
type logFbInterface interface {
	Build(*config.Logging) (glogger.Logger, error)
}

// accessors for factoryBuilderMap
func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}

func Build(lc *config.Logging) (glogger.Logger, error) {
	loggerType := lc.Code
	l, err := GetLogFactoryBuilder(loggerType).Build(lc)
	if err != nil {
		return l, errors.Wrap(err, "")
	}
	return l, nil
}
