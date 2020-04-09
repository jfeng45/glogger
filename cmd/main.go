package main

import (
	"fmt"
	"github.com/jfeng45/glogger/logconfig"
	"github.com/jfeng45/glogger/logfactory"
)

func main() {
	lc := logconfig.LogConfig{logfactory.ZAP, "debug", true}
	//lc := config.LogConfig{config.LOGRUS, "debug", true}
	logger, err := logfactory.InitLogger(lc)
	if err != nil {
		fmt.Errorf("Init log failed")
	}
	logger.Infof("info")
	logger.Debugf("debug")
	logger.Warnf("error")
	//logger.Log.Warnf("warn")
}


