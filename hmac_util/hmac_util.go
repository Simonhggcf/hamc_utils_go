package hmacutil

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// GenerateHMAC 生成消息的 HMAC
func GenerateHMAC(message, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC 验证消息的 HMAC 是否匹配
func VerifyHMAC(message, key []byte, receivedHMAC string) (bool, error) {
	expectedHMAC := GenerateHMAC(message, key)
	if expectedHMAC == receivedHMAC {
		return true, nil
	}
	return false, errors.New("HMAC verification failed")
}