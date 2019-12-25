package handler

import (
	"fmt"
	"github.com/translate/sruntime"
)

// 返回钱方描述中所有的数据
func Get_descript() interface{} {
	_sql := "SELECT name, descript from qf_descript"
	db, ok := srunning.Gsvr.Dbs.Pools["qf_payment"]
	if !ok {
		return nil
	}

	rows, err := db.Query(_sql)
	if err != nil {
		panic(fmt.Sprintf("db query error:%s", err))
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
			return nil
		}
		record := make(map[string]string)

		for i, col := range values {
			if col != nil {
				//字段名 = 字段信息
				record[columns[i]] = string(col.([]byte))
			}
		}
		s[record["name"]] = record["descript"]
	}

	return s
}

// 缓存钱方翻译信息
func Get_translate() interface{} {
	fmt.Println("select again")
	_sql := "SELECT src_word, lang, dst_word from translate"
	db, ok := srunning.Gsvr.Dbs.Pools["qf_payment"]
	if !ok {
		return nil
	}

	rows, err := db.Query(_sql)
	if err != nil {
		panic(fmt.Sprintf("db query error:%s", err))
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
			return nil
		}
		record := make(map[string]string)

		for i, col := range values {
			if col != nil {
				//字段名 = 字段信息
				record[columns[i]] = string(col.([]byte))
			}
		}
		src_lang := fmt.Sprintf(record["src_word"] + "_" + record["lang"])
		s[src_lang] = record["dst_word"]
	}

	fmt.Println("translate info", s)
	return s
}
