package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/khodemobin/pilo/auth/internal/config"
)

func IsLocal() bool {
	return config.GetConfig().App.Env == "local"
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

func ToJson(v any) (string, error) {
	j, err := json.Marshal(v)

	fmt.Println(string(j))
	return string(j), err
}
