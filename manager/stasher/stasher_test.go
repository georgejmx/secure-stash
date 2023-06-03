package stasher

import (
	"testing"
)

const TEST_PASSWORD = "something!secureY"

var validEntries = []string{"tootatabah, Testing123!", "somethingreallyreallyreallyreallyreallyreallylong", "{\"phrases\":[\"foo\",\"bar\"]}", "som3_rAndom_password"}

func TestCorrectNonceSize(t *testing.T) {
	if NONCE_SIZE <= 0 || NONCE_SIZE > 32 {
		t.Fatal("Incorrect nonce size constant set")
	}
}

// Performs encryption and decryption of a input string; throwing an error
// in the case of a mismatch
func stashEncryptsString(stash Stasher, inputString string) (bool, error) {
	result, err := stash.DecryptBytes(stash.EncryptText(inputString))
	if err != nil || result != inputString  {
		return false, err
	}

	return true, nil
}

func TestValidStashing(t *testing.T) {
	testStasher := Stasher{}
	testStasher.InitStasher(RawToHash(TEST_PASSWORD))

	for _, entry := range validEntries {
		ok, err := stashEncryptsString(testStasher, entry)
		if err != nil {
			t.Errorf("Error when decrypting entry %s: %v", entry, err)
		} else if !ok {
			t.Errorf("Inconsistent encryption of entry: %s", entry)
		}
	}
}