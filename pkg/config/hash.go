package config

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

// hashing token

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	return hex.EncodeToString(hash[:])
}

// random token

func RandomToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}

// random otp

func RandomOTP() string {
	const chars = "0123456789ABCDEFGHJKMNPQRSTUVWXYZ"
	b := make([]byte, 6)
	rand.Read(b)
	for i := range b {
		b[i] = chars[int(b[i])%len(chars)]
	}
	return string(b)
}
