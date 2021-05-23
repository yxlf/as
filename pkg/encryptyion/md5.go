package encryptyion

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 把字符串通过md5加密
func MD5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	sum := hash.Sum(nil)

	return hex.EncodeToString(sum)
}
