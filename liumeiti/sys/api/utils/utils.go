package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func GenMd5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}
