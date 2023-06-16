package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

func StrLen(s string) int {
	return utf8.RuneCountInString(s)
}

func Now() int64 {
	return time.Now().Unix()
}

func Date(v int64) string {
	if v == 0 {
		v = Now()
	}

	return time.Unix(v, 0).Format("2006-01-02")
}

func EmailSuffix(email string) string {
	v := strings.Split(email, "@")
	if len(v) == 2 {
		return v[1]
	}

	return email
}

func GenMD5(b []byte) string {
	return fmt.Sprintf("%x", md5.Sum(b))
}
