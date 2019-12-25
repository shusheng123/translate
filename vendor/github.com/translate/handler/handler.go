package handler

import (
	"fmt"
	"git.qfpay.net/server/goqfpay/logger"
	"github.com/translate/sruntime"
)

type Handler struct{}

func (tanslate *Handler) Translate(src_word string, lang string) (string, error) {

	defer func() {
		if err := recover(); err != nil {
			logger.Infof("translate server error: %s", err)
		}
	}()

	qf_descript := get_descript()

	// 判断是否有描述
	v, ok := qf_descript[src_word]
	if ok {
		src_word = v
	}

	dst_word := get_translate(src_word, lang)

	return dst_word, nil
}

// 检查是否在钱方描述中
func get_descript() map[string]string {

	descript := srunning.Gsvr.Cache.Get("descript")

	value, ok := descript.(map[string]string)
	if !ok {
		//return nil, false
		panic("get descript error!!")
	}

	return value

}

// 获取翻译
func get_translate(src_word string, lang string) string {

	translate_info := srunning.Gsvr.Cache.Get("translate")
	logger.Infof("translate_info:%s", translate_info)

	value, ok := translate_info.(map[string]string)
	if !ok {
		panic("get translate_info error!!")
	}
	src_lang := fmt.Sprintf(src_word + "_" + lang)
	logger.Infof("get translate info by key:%s", src_lang)
	if val, ok := value[src_lang]; ok {
		return val
	}

	return src_word
}
