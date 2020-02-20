// package loggerfactory handles creating concrete logger with factory method pattern
package glogger

// logger mapp to map logger code to logger builder
var logfactoryBuilderMap = map[string]logFbInterface{
	ZAP:    &ZapFactory{},
	LOGRUS: &LogrusFactory{},
}

// interface for logger factory
type logFbInterface interface {
	Build(*LogConfig) (Logger, error)
}

// accessors for factoryBuilderMap
func GetLogFactoryBuilder(key string) logFbInterface {
	return logfactoryBuilderMap[key]
}
