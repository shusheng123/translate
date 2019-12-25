package main

import (
	"git.qfpay.net/server/goqfpay/logger"
	"github.com/translate/handler"
	"github.com/translate/myconf"
	"github.com/translate/rpcserver"
	"github.com/translate/sruntime"
	"os"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	arg_num := len(os.Args)
	if arg_num != 2 {
		logger.Warn("input param error")
		return
	}
	var filename = os.Args[1]
	err := myconf.Parseconf(filename)
	if err != nil {
		return
	}

	logger.SetConsole(myconf.Scnf.LogStdOut)
	logger.SetRollingDaily(myconf.Scnf.LogDir, myconf.Scnf.LogFile, myconf.Scnf.LogFileErr)
	loglevel, _ := logger.LoggerLevelIndex(myconf.Scnf.LogLevel)
	logger.SetLevel(loglevel)

	srunning.CreateRuntime()
	logger.Infof("server start")

	logger.Debugf("db token file: %s", myconf.Scnf.TokenFile)
	err = srunning.Gsvr.LoadDb()

	if err != nil {
		os.Exit(1)
	}

	srunning.Gsvr.Cache.Put("descript", handler.Get_descript, 20)
	srunning.Gsvr.Cache.Put("translate", handler.Get_translate, 20)

	handler := &handler.Handler{}
	server := rpcserver.NewRPCserver(handler)
	server.Start()
}
