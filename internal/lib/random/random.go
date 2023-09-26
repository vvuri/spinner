package random

import (
	"golang.org/x/exp/rand"
	"time"
)

func NewRandomString(aliasLength int) string {
	var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lenLetter := len(letters)

	rnd := rand.New(rand.NewSource(uint64(time.Now().UnixMilli())))

	alias := make([]byte, aliasLength)
	for i := 0; i < aliasLength; i++ {
		alias[i] = letters[rnd.Intn(lenLetter)]
	}

	aliasStr := string(alias)

	// TODO: validate to unique in DB

	return aliasStr
}
