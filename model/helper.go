package model

import (
	"encoding/json"
	"math/rand"
	"time"
	"yac-go/log"
)

const IdLength = 64

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

func GenerateId() string {
	n := IdLength
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// Marshal converts map to interface
func Marshal(m map[string]interface{}, v interface{}) {
	byteNote, err := json.Marshal(m)
	if err != nil {
		log.Panic("Failed to marshal json:", err)
	}

	if err := json.Unmarshal(byteNote, v); err != nil {
		log.Panic("Failed to unmarshal json:", err)
	}
}

// Unmarshal converts interface to map
func Unmarshal(i interface{}) map[string]interface{} {
	var conv map[string]interface{}

	byteNote, err := json.Marshal(i)
	if err != nil {
		log.Panic("Failed to marshal json:", err)
	}

	if err := json.Unmarshal(byteNote, &conv); err != nil {
		log.Panic("Failed to unmarshal json:", err)
	}
	return conv
}
