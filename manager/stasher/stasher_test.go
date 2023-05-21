package stasher

import "testing"

const IV = "12345678901234567890123456789012"

var validEntries = []string{"tootatabah, Testing123!", "somethingreallyreallyreallyreallyreallyreallylong", "{\"phrases\":[\"foo\",\"bar\"]}"}

func TestValidStashing(t *testing.T) {
	testStasher := Stasher{}
	testStasher.InitStasher(IV)
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