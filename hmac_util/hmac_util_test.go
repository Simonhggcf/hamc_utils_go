package hmacutil

import (
	"testing"
)

func TestGenerateHMAC(t *testing.T) {
	message := []byte("Hello, World!")
	key := []byte("supersecretkey")
	expectedHMAC := "66b1fe51c54d313a85d84dd57685f2e88a1adaa85417c78c83f88cba9bd7fbeb" // replace with your actual HMAC

	hmacValue := GenerateHMAC(message, key)
	if hmacValue != expectedHMAC {
		t.Errorf("Expected HMAC %s, got %s", expectedHMAC, hmacValue)
	}
}

func TestVerifyHMAC(t *testing.T) {
	message := []byte("Hello, World!")
	key := []byte("supersecretkey")
	validHMAC := GenerateHMAC(message, key)

	// Test successful verification
	isValid, err := VerifyHMAC(message, key, validHMAC)
	if !isValid || err != nil {
		t.Error("HMAC verification should succeed")
	}

	// Test failed verification
	invalidHMAC := "incorrecthmacvalue"
	isValid, err = VerifyHMAC(message, key, invalidHMAC)
	if isValid || err == nil {
		t.Error("HMAC verification should fail")
	}
}