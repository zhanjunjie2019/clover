package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5SaltString(src, salt string) string {
	srcCode := md5.Sum([]byte(src + "@" + salt))
	return fmt.Sprintf("%x", srcCode)
}
