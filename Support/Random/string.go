package Random

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	Uppercase     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Lowercase     = "abcdefghijklmnopqrstuvwxyz"
	Alphabetic    = Uppercase + Lowercase
	Numeric       = "0123456789"
	letterBytes   = Alphabetic + Numeric
	Hex           = Numeric + "abcdef"
	symbols       = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~" // 32
	characters    = Alphabetic + Numeric + symbols       // 94
)

func (r *_Random) randSrc() int64 {
	r.mr.Lock()
	defer r.mr.Unlock()
	return r.pseudo.Int63()
}

// GenerateSecureRandomKey method generates the random bytes for given length using
// `crypto/rand`.
func (r *_Random) SecureKey(length int) []byte {
	k := make([]byte, length)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		// fallback to math based random key generater
		return r.Key(length)
	}
	return k
}

// SecureRandomString method generates the random string for given length using
// `crypto/rand`.
func (r *_Random) SecureRandomString(length int) string {
	return hex.EncodeToString(r.SecureKey(length / 2))
}

// GenerateRandomKey method generates the random bytes for given length using
// `math/rand.Source` and byte mask.
func (r *_Random) Key(length int, charsets ...string) []byte {
	charset := strings.Join(charsets, "")
	if charset == "" {
		charset = letterBytes
	}

	b := make([]byte, length)
	// A randSrc() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, r.randSrc(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = r.randSrc(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(charset) {
			b[i] = charset[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return b
}

// RandomString method generates the random string for given length using
// `math/rand.Source` and byte mask.
func (r *_Random) String(length int, charsets ...string) string {
	return string(r.Key(length, charsets...))
}
