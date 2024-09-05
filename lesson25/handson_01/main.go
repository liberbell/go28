package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%s", h.Sum(nil))
}
