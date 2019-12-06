package handler

import (
	"fmt"
)

type Handler struct {
	Memo string
	Addr string
}

func (tanslate Handler) Translate(t string) (string, error) {
	dst_str := "hello word!"
	fmt.Printf("src: %s dst:%s", t, dst_str)

	return dst_str, nil
}
