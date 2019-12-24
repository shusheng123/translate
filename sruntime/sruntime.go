package srunning

import (
	"git.qfpay.net/server/goqfpay/dbenc"
	"git.qfpay.net/server/goqfpay/dbpool"
	"git.qfpay.net/server/goqfpay/logger"
	"github.com/translate/cache"
	"github.com/translate/myconf"
)

type SvrRunning struct {
	Dbs   *dbpool.Dbpool
	Cache *cache.CacheDict
}

var Gsvr *SvrRunning

func CreateRuntime() {
	Gsvr = new(SvrRunning)
	Qfdb := dbenc.DbConfNew(myconf.Scnf.TokenFile)
	Gsvr.Dbs = dbpool.DbpoolNew(Qfdb)
	Gsvr.Cache = cache.NewCache()
}

func (g_rt *SvrRunning) LoadDb() error {
	logger.Debugf("url: %s", myconf.Scnf.PaymentDbTk)
	err := g_rt.Dbs.Add("qf_payment", myconf.Scnf.PaymentDbTk)
	if err != nil {
		logger.Warnf("db add %s %s", myconf.Scnf.PaymentDbTk, err.Error())
		return err
	}
	return nil
}
