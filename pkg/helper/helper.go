package helper

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/khodemobin/pilo/auth/internal/config"
)

func IsLocal(cfg *config.Config) bool {
	return cfg.App.Env == "local"
}

func ToMD5(s string) string {
	hasher := md5.New()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func HasString(list []string, find string) bool {
	for _, b := range list {
		if b == find {
			return true
		}
	}
	return false
}
