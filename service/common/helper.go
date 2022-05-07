package common

import (
	"crypto/md5"
	"fmt"
)

func EncryptByMd5(param string) string {
	data := []byte(param)
	has := md5.Sum(data)
	res := fmt.Sprintf("%x", has)
	return res
}
