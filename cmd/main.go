package main

import (
	"github.com/jfeng45/glogger/config"
	"github.com/jfeng45/glogger/factory"
	"log"
)

func main() {
	lc := config.Logging{config.ZAP, config.DEBUG, true}
	//lc := config.Logging{config.ZAP, config.INFO, true}
	//lc := config.Logging{config.ZAP, config.WARN, true}
	//lc := config.Logging{config.ZAP, config.ERROR, true}
	logger, err := factory.Build(&lc)
	if err != nil {
		log.Fatal("Init log failed")
	}
	logger.Infof("info")
	logger.Debugf("debug")
	logger.Warnf("error")
	//logger.Log.Warnf("warn")
}


