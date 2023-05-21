package stasher

import (
	"testing"
)

const USER_PASSWORD = "something!secureY"

var validEntries = []string{"tootatabah, Testing123!", "somethingreallyreallyreallyreallyreallyreallylong", "{\"phrases\":[\"foo\",\"bar\"]}"}

func TestCorrectNonceSize(t *testing.T) {
	if NONCE_SIZE > 32 {
		t.Fatalf("Incorrect nonce size constant set")
	}
}

func TestValidStashing(t *testing.T) {
	testStasher := Stasher{}
	testStasher.InitStasher(RawToHash(USER_PASSWORD))

	if !testStasher.HasIntegrity(USER_PASSWORD) {
		t.Fatalf("stasher cannot encrypt or decrypt data")
	}

	for _, entry := range validEntries {
		encryptedEntry := testStasher.EncryptText(entry)
		decryptedEntry, err := testStasher.DecryptBytes(encryptedEntry)
		
		if err != nil {
			t.Errorf("error when decrypting entry %s: %v", entry, err)
		} else if entry != decryptedEntry {
			t.Errorf("Inconsistent encryption of entry: %s", entry)
		}
	}
}