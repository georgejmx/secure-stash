package utils

import "crypto/sha256"

// Remove an element from string slice
func Remove(s []string, r string) []string {
    for i, v := range s {
        if v == r {
            return append(s[:i], s[i+1:]...)
        }
    }
    return s
}

// Converts a string password to hashed bytes
func RawToHash(raw string) [32]byte {
	return sha256.Sum256([]byte(raw))
}