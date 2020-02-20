package main

import (
	"fmt"
	"github.com/jfeng45/glogger"
)

func main() {
	lc := glogger.LogConfig{glogger.ZAP, "debug", true}
	//lc := config.LogConfig{config.LOGRUS, "debug", true}
	logger, err := glogger.InitLogger(lc)
	if err != nil {
		fmt.Errorf("Init log failed")
	}
	logger.Infof("info")
	logger.Debugf("debug")
	logger.Warnf("error")
	//logger.Log.Warnf("warn")
}


