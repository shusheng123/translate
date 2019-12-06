package main

import (
	"git.qfpay.net/server/goqfpay/logger"
	"github.com/translate/handler"
	"github.com/translate/rpcserver"
)

func main() {

	logger.SetConsole(true)
	logger.SetRollingDaily("/Users/mrshu/Documents/project/gopath/src/github.com/translate/log",
		"0.translate.log", "0.translate.error.log")
	loglevel, _ := logger.LoggerLevelIndex("DEBUG")
	logger.SetLevel(loglevel)

	handler := &handler.Handler{
		Memo: "Translate server",
		Addr: "0.0.0.0:6000",
	}
	server := rpcserver.NewRPCserver(handler)

	server.Start()
}
