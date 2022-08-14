package nthclient

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

func CalculateAPIHostname(seed, tld string) string {
	t := time.Now().Truncate(0).UTC().Format("2006-01-02")
	digest := md5.Sum([]byte(seed + t))
	return fmt.Sprintf("www.%s.%s",
		hex.EncodeToString(digest[0:6]),
		tld)
}

func VerifyResponse(response string, pubkey *rsa.PublicKey) ([]byte, error) {
	parts := strings.SplitN(response, "*-*", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("data was not found in the response. parts found: %d", len(parts))
	}

	signature, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("signature decoding failed: %w", err)
	}
	hashed := sha256.Sum256([]byte(parts[1]))

	err = rsa.VerifyPKCS1v15(pubkey, crypto.SHA256, hashed[:], signature)
	if err != nil {
		return nil, fmt.Errorf("signature verification failed: %w", err)
	}
	return []byte(parts[1]), nil
}

func Decrypt(ciphertext, key string) ([]byte, error) {
	bytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return nil, fmt.Errorf("ciphertexttext hex decoding failed: %w", err)
	}

	if len(bytes) < 16 {
		return nil, fmt.Errorf("too short ciphertext bytes len: %d", len(bytes))
	}

	iv := bytes[0:16]
	data := bytes[16:]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("can't create block ciphertext: %w", err)
	}

	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(data))
	stream.XORKeyStream(plaintext, data)

	return plaintext, nil
}
