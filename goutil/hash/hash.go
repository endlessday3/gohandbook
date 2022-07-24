package hash

import (
	"crypto/md5"
	"encoding/hex"
)

// StrToMd5 字符串转MD5字符串
func StrToMd5(s string) string {
	md5S := md5.New()
	md5S.Write([]byte(s))
	return hex.EncodeToString(md5S.Sum(nil))
}
