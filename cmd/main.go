package main

import (
	"github.com/jfeng45/glogger/logconfig"
	"github.com/jfeng45/glogger/logfactory"
	"log"
)

func main() {
	lc := logconfig.LogConfig{logfactory.ZAP, "debug", true}
	//lc := config.LogConfig{config.LOGRUS, "debug", true}
	logger, err := logfactory.InitLogger(&lc)
	if err != nil {
		log.Fatal("Init log failed")
	}
	logger.Infof("info")
	logger.Debugf("debug")
	logger.Warnf("error")
	//logger.Log.Warnf("warn")
}


