package stasher

import (
	"example/secure-stash/manager/utils"
	"testing"
)

const TEST_PASSWORD = "something!secureY"

var validEntries = []string{"tootatabah, Testing123!", "somethingreallyreallyreallyreallyreallyreallylong", "{\"phrases\":[\"foo\",\"bar\"]}", "som3_rAndom_password"}

// Asserts that the nonce size specified is valid
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

// Validates AES integrity on test cases
func TestValidStashing(t *testing.T) {
	testStasher := Stasher{}
	testStasher.InitStasher(utils.RawToHash(TEST_PASSWORD))

	for _, entry := range validEntries {
		ok, err := stashEncryptsString(testStasher, entry)
		if err != nil {
			t.Errorf("Error when decrypting entry %s: %v", entry, err)
		} else if !ok {
			t.Errorf("Inconsistent encryption of entry: %s", entry)
		}
	}
}

// Checks that hash to secrets reproducibly produces outputs
func TestReproducibleInitialisation(t *testing.T) {
	testCases := [2][32]byte{ utils.RawToHash("test1"), utils.RawToHash("test2")}

	testSeed, testNonce := hashToSecrets(testCases[0])
	if len(testSeed) != NONCE_SIZE || len(testNonce) != 32 {
		t.Fatal("Unexpected output when generating seed")
	}


	testSeed2, testNonce2 := hashToSecrets(testCases[0])
	if testSeed != testSeed2 || testNonce != testNonce2 {
		t.Error("Generating encryption seeds is not reproducible")
	}

	testDifferentSeed, testDifferentNonce := hashToSecrets(testCases[1])
	if testSeed == testDifferentSeed || testNonce == testDifferentNonce {
		t.Error("Generating encryption seeds produces constant output")
	}
}