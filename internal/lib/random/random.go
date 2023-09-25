package random

import "golang.org/x/exp/rand"

func NewRandomString(aliasLength int) string {
	var letters = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lenLetter := len(letters)

	alias := make([]byte, aliasLength)
	for i := 0; i < aliasLength; i++ {
		alias[i] = letters[rand.Intn(lenLetter)]
	}

	aliasStr := string(alias)

	// TODO: validate to unique in DB

	return aliasStr
}
