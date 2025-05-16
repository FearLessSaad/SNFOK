package otp

import (
	"crypto/rand"
	"math/big"
)

func GenerateSixDigitOTP() (string, error) {
	otp := ""
	for i := 0; i < 6; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += n.String()
	}
	return otp, nil
}
