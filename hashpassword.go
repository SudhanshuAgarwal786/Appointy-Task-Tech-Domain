package main

import "crypto/sha256"

func HashPassword(password string) string {
	temp := sha256.Sum256([]byte(password))
	return string(temp[:])
}
