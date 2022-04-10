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
	hashes := md5.New()
	hashes.Write([]byte(s))
	return hex.EncodeToString(hashes.Sum(nil))
}

func HasString(list []string, find string) bool {
	for _, b := range list {
		if b == find {
			return true
		}
	}
	return false
}

func DefaultResponse(data interface{}, message string, code int) struct {
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
} {
	r := struct {
		Data    interface{} `json:"data"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
	}{
		Data:    data,
		Code:    code,
		Message: message,
	}

	return r
}
