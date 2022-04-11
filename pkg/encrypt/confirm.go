package encrypt

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/khodemobin/pilo/auth/internal/domain"
)

func GenerateConfirmCode(phone string) (*domain.ConfirmCode, int, error) {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 900000
	code := rand.Intn(max-min+1) + min
	hash, err := Hash(fmt.Sprint(code))
	if err != nil {
		panic("internal error, can not create hashed token")
	}
	time := 3 * time.Minute

	return &domain.ConfirmCode{
		Phone:     phone,
		Hash:      hash,
		ExpiresIn: time,
	}, code, err
}
