package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Lately, I've been, I've been losin' sleep Dreamin' about the things that we could be But baby, I've been, I've been prayin' hard Said, No more countin' dollars, we'll be countin' stars Yeah, we'll be countin' stars"

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("i`m giving her all she`s got Captain", err)
	}
	fmt.Println(string(bs))
}
