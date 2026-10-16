package helper

import (
	"crypto/rand"
	"math/big"
)

const codeCharset = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"

func Generate4CharRoomCode() (string, error) {
	b := make([]byte, 4)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(codeCharset))))
		if err != nil {
			return "", err
		}
		b[i] = codeCharset[num.Int64()]
	}
	return string(b), nil
}
