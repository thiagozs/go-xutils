package md5

import (
	cmd5 "crypto/md5"
	"encoding/hex"
)

type Md5 struct{}

func New() *Md5 {
	return &Md5{}
}

func (m *Md5) MD5Hash(str string) string {
	s := cmd5.New()
	_, _ = s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
