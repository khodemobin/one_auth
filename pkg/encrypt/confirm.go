package encrypt

import (
	"math/rand"
	"time"
)

type ConfirmCode struct {
	Code int
	Hash string
	Time time.Time
}

func GenerateConfirmCode() (*ConfirmCode, error) {
	min := 100000
	max := 900000
	code := rand.Intn(max-min) + min
	hash, err := Hash(string(rune(code)))
	if err != nil {
		panic("internal error, can not create hashed token")
	}
	time := time.Now()

	return &ConfirmCode{
		Code: code,
		Hash: hash,
		Time: time,
	}, err
}
