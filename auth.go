package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

func createHMAC(message, key []byte) []byte {
	hash := hmac.New(sha256.New, key)
	hash.Write(message)

	messageWithMAC := hash.Sum(message)

	return messageWithMAC
}

func validateHMAC(message, messageMAC, key []byte) bool {
	hash := hmac.New(sha256.New, key)
	hash.Write(message)

	expectedMAC := hash.Sum(nil)

	return hmac.Equal(messageMAC, expectedMAC)
}
