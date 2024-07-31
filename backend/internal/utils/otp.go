package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Token struct {
	Secret string
	Hash   string
}

func GenerateOTP() (*Token, error) {
	bigInt, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return nil, err
	}
	sixDigitNum := bigInt.Int64() + 100000

	// Convert the integer to a string and get the first 6 characters
	sixDigitStr := fmt.Sprintf("%06d", sixDigitNum)

	token := Token{
		Secret: sixDigitStr,
	}

	hash := sha256.Sum256([]byte(token.Secret))

	token.Hash = fmt.Sprintf("%x\n", hash)

	return &token, nil
}
