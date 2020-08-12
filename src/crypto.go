package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func getSha256(secret string, method string, path string, body string, timestamp string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	bs := []byte(fmt.Sprintf("%s\n%s\n%s\n%s", method, path, body, timestamp))
	mac.Write(bs)
	return hex.EncodeToString(mac.Sum(nil))
}

func getMd5(body []byte) string {
	fmt.Println("getMd5 - input", string(body))
	h := md5.New()
	io.WriteString(h, string(body))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
