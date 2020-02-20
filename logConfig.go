package glogger

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	LOGRUS string = "logrus"
	ZAP    string = "zap"
)
// LogConfig represents logger handler
// Logger has many parameters can be set or changed. Currently, only three are listed here. Can add more into it to
// fits your needs.
type LogConfig struct {
	// log library name
	Code string `yaml:"code"`
	// log level
	Level string `yaml:"level"`
	// show caller in log message
	EnableCaller bool `yaml:"enableCaller"`
}

//func validateLogger(logConfig LogConfig) error {
//	key := logConfig.Code
//	zcMsg := " in validateLogger doesn't match key = "
//	if ZAP != key {
//		errMsg := ZAP + zcMsg + key
//		return errors.New(errMsg)
//	}
//	if LOGRUS != lc.Code {
//		errMsg := LOGRUS + zcMsg + key
//		return errors.New(errMsg)
//	}
//	return nil
//}

