package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"
)

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	random := make([]rune, n)
	for i := range random {
		random[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(random)
}

func GenerateHashHMACSHA256(randomKey, uriPath, requestBody, secretKey string) string {
	data := randomKey + uriPath + requestBody
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	hash := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return hash
}
