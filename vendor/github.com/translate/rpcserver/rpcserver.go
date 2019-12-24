package rpcserver

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"git.qfpay.net/server/goqfpay/logger"
	"github.com/translate/gen-go/translate"
	"github.com/translate/handler"
	"github.com/translate/myconf"
	"os"
	"sync"
	"time"
)

type RPCserver struct {
	server  thrift.TServer
	handler *handler.Handler
	wg      sync.WaitGroup
}

func NewRPCserver(handler *handler.Handler) *RPCserver {

	return &RPCserver{handler: handler}

}

func (rpc *RPCserver) Start() {
	rpcAddr := fmt.Sprintf("%s:%d", myconf.Scnf.RPCServerAddr, myconf.Scnf.RPCServerPort)
	logger.Infof("rpc server start %s", rpcAddr)
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	serverTransport, err := thrift.NewTServerSocketTimeout(rpcAddr, time.Duration(myconf.Scnf.RPCServerTimeout)*time.Second)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}
	rpc_h := &rpcImpl{handler: rpc.handler}
	processor := translate.NewTranslateProcessor(rpc_h)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	rpc.server = server
	defer rpc.wg.Done()
	server.Serve()
}

func (rpc *RPCserver) stop() {
	rpc.server.Stop()
}

type rpcImpl struct {
	handler *handler.Handler
}

func (imp *rpcImpl) Translate(t string) (string, error) {

	return imp.handler.Translate(t)
}
