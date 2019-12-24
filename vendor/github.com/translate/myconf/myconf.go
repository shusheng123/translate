package myconf

import (
	"fmt"
	"git.qfpay.net/server/goqfpay/confparse"
	"git.qfpay.net/server/goqfpay/gconfig"
)

type Mcnf struct {
	IP   string `confpos:"rpc:addr" dtype:"base"`
	PORT string `confpos:"rpc:port" dtype:"base"`

	// RPC 配置
	RPCServerAddr    string `confpos:"rpc:addr" dtype:"base"`
	RPCServerPort    int    `confpos:"rpc:port" dtype:"base"`
	RPCServerTimeout int    `confpos:"rpc:timeout" dtype:"base"`
	EnableRPC        bool   `confpos:"rpc:enable_rpc" dtype:"base"`

	// DB 配置
	PaymentDbTk string `confpos:"db:qf_payment" dtype:"base"`
	TokenFile   string `confpos:"db:tk_file" dtype:"base"`

	// 日志配置
	LogFile    string `confpos:"log:logfile" dtype:"base"`
	LogFileErr string `confpos:"log:logfile_err" dtype:"base"`
	LogDir     string `confpos:"log:logdir" dtype:"base"`
	LogLevel   string `confpos:"log:loglevel" dtype:"base"`
	LogStdOut  bool   `confpos:"log:logstdout" dtype:"base"`

	// redis 配置
	RedisAddr        []string `confpos:"redis:redis_url" item_split:"," dtype:"base"`
	RedisTimeout     int      `confpos:"redis:timeout" dtype:"base"`
	RedisMaxConnAct  int      `confpos:"redis:redis_pool_maxact_conn" dtype:"base"`
	RedisMaxConnIdle int      `confpos:"redis:redis_pool_maxidle_conn" dtype:"base"`
	PrintRedisLog    bool     `confpos:"redis:print_redis_log" dtype:"base"`
	MsgStatTTL       int      `confpos:"redis:msg_stat_ttl" dtype:"base"`
}

var Scnf *Mcnf = new(Mcnf)

func Parseconf(filename string) error {
	cfg := gconfig.NewGconf(filename)
	err := cfg.GconfParse()
	if err != nil {
		fmt.Printf("parse %s %s", filename, err.Error())
		return err
	}
	cp := confparse.CpaseNew(filename)
	err = cp.CparseGo(Scnf, cfg)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("redis_url: %s\n", Scnf.RedisAddr[0])
	return err
}
