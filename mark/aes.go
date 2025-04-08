package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"time"
)

func BenchmarkAes(cfg *Config, countFn func(int)) {
	cfg = cfg.New()
	start := time.Now()

	key := make([]byte, 32)
	nonce := make([]byte, 12)
	_, _ = rand.Read(key)
	_, _ = rand.Read(nonce)
	plain := make([]byte, 1<<16)

	block, _ := aes.NewCipher(key)
	aesgcm, _ := cipher.NewGCM(block)

	it := 0
	for {
		_ = aesgcm.Seal(nil, nonce, plain, nil)
		it++
		if cfg.ShouldStop(start, it) {
			break
		}
	}

	if countFn != nil {
		countFn(it)
	}
}
