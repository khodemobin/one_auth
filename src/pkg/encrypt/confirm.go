package encrypt

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/khodemobin/pilo/auth/internal/model"
)

func GenerateConfirmCode(phone string) (*model.ConfirmCode, int, error) {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 900000
	code := rand.Intn(max-min+1) + min
	hash, err := Hash(fmt.Sprint(code))
	if err != nil {
		return nil, 0, err
	}
	time := 3 * time.Minute

	return &model.ConfirmCode{
		Phone:     phone,
		Hash:      hash,
		ExpiresIn: time,
	}, code, err
}
