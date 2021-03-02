package core

import (
	"flag"
	"github.com/gin-gonic/gin"
	"goal-layout/global"
	"goal-layout/pkg/logger"
	"log"
)

func StartModule() {
	var e error
	if e = initFlag(); e != nil {
		log.Fatalf("initFlag e: %v", e)
	}

}

var (
	port    string
	runMode string
)

func initFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", gin.Mode(), "启动模式")
	flag.Parse()

	return nil
}

func initLogger() (e error) {

	logSet := global.LogSetting
	global.Logger, e = logger.NewLogger(logSet.Formatter, logSet.Level, logSet.ReportCaller, logSet.SavePath)
	if e != nil {
		return
	}

	return
}
