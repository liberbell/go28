package main

import "encoding/base64"

func main() {
	s := "Lately, I've been, I've been losin' sleep Dreamin' about the things that we could be But baby, I've been, I've been prayin' hard Said, No more countin' dollars, we'll be countin' stars Yeah, we'll be countin' stars"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
}
