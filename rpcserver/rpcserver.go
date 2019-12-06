package rpcserver

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"git.qfpay.net/server/goqfpay/logger"
	"github.com/translate/gen-go/translate"
	"os"
	"reflect"
	"sync"
)

type RPCserver struct {
	server  thrift.TServer
	handler interface{}
	wg      sync.WaitGroup
}

func NewRPCserver(handler interface{}) *RPCserver {

	return &RPCserver{handler: handler}

}

func (rpc *RPCserver) Start() {
	hand_fie := reflect.ValueOf(rpc.handler).Elem()
	addrs := hand_fie.FieldByName("Addr").String()
	handle_ := hand_fie.Interface().(translate.Translate)

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	serverTransport, err := thrift.NewTServerSocket(addrs)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	processor := translate.NewTranslateProcessor(handle_)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	rpc.server = server
	defer rpc.wg.Done()
	logger.Info("thrift server start: ", addrs)
	server.Serve()
}

func (rpc *RPCserver) stop() {
	rpc.server.Stop()
}
