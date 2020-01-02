package utils

import (
	"crypto/md5"
	"fmt"
	"crypto/rand"
	"io"
	"encoding/hex"
)

// GenMd5 md5.
func GenMd5(str string) string {
	hash := md5.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

// NewUUID 生成uuid.
func NewUUID()(string , error)  {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader,uuid)
	if n !=len(uuid) || err !=nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8] &^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[8] &^0xcf | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x",uuid[0:4],uuid[4:6],uuid[6:8],uuid[8:10], uuid[10:]), nil

}
