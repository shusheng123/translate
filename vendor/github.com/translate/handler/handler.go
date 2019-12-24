package handler

import (
	"git.qfpay.net/server/goqfpay/logger"
	"github.com/translate/sruntime"
)

type Handler struct{}

func (tanslate *Handler) Translate(t string) (string, error) {

	qf_descript, check := check(t)
	if !check {
		logger.Info("not in descript")
	} else {
		t = qf_descript[t]
	}

	return t, nil
}

// 返回钱方描述中所有的数据
func get_descript() (map[string]string, error) {
	_sql := "SELECT name, descript from qf_descript"
	db, ok := srunning.Gsvr.Dbs.Pools["qf_payment"]
	if !ok {
		return nil, nil
	}

	rows, err := db.Query(_sql)
	if err != nil {
		return nil, nil
	}

	columns, _ := rows.Columns()
	var s map[string]string
	s = make(map[string]string)

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, nil
		}
		record := make(map[string]string)

		for i, col := range values {
			if col != nil {
				//字段名 = 字段信息
				record[columns[i]] = string(col.([]byte))
				logger.Info(record)
			}
		}
		s[record["name"]] = record["descript"]
	}

	return s, nil
}

// 检查是否在钱方描述中
func check(t string) (map[string]string, bool) {

	descript, err := get_descript()
	if err != nil {
		return descript, false
	}
	if _, ok := descript[t]; ok {
		return descript, true
	}
	return descript, false
}
