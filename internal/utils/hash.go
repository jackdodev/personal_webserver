package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

// Sha256Hex returns the full SHA256 hex digest (64 characters).
func Sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// Sha256HexN returns the first n characters of the SHA256 hex digest.
// If n <= 0 it returns an empty string. If n > 64 it returns the full digest.
// Note: truncating a 64-char hex (256-bit) digest to 16 chars reduces it to
// a 64-bit value (16 hex chars = 64 bits) which has collision risks.
func Sha256HexN(s string, n int) string {
	if n <= 0 {
		return ""
	}
	full := Sha256Hex(s)
	if n >= len(full) {
		return full
	}
	return full[:n]
}

// Sha256Hex16 is a convenience wrapper that returns the first 16 hex chars.
func Sha256Hex16(s string) string {
	return Sha256HexN(s, 16)
}

// (no examples in this file)
