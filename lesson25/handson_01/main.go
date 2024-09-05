package main

import (
	"crypto/hmac"
	"crypto/sha256"
)

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
}
