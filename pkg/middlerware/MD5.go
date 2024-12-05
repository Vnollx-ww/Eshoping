package middlerware

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5Encrypt(password string) string {
	hash := md5.New()
	hash.Write([]byte(password))
	return hex.EncodeToString(hash.Sum(nil))
}
