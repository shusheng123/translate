package main

import (
	"flag"
	"fmt"
	"git.qfpay.net/server/goqfpay/logger"
	"github.com/translate/handler"
	"github.com/translate/myconf"
	"github.com/translate/rpcserver"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	conf_file := flag.String("conf_file", "conf_file", "please add conf file")
	flag.Parse()
	conf, err := myconf.ParseConf(*conf_file)
	if err != nil {
		fmt.Println("parse conf error")
		return
	}

	logger.SetConsole(true)
	logger.SetRollingDaily(conf.Gcf["log"]["logdir"],
		conf.Gcf["log"]["logfile"],
		conf.Gcf["log"]["logfile_err"])
	loglevel, _ := logger.LoggerLevelIndex(conf.Gcf["log"]["loglevel"])
	logger.SetLevel(loglevel)

	handler := &handler.Handler{
		Memo: "Translate server",
		Addr: "0.0.0.0:6000",
	}
	server := rpcserver.NewRPCserver(handler)
	server.Start()
}
