package manager

import "testing"

const TEST_PASSWORD = "something!secureY"

func TestInvalidPasswordFails(t *testing.T) {
	c = &TestCacher{fresh: false}

	ok, err := Init(TEST_PASSWORD)
	if ok {
		t.Fatal("Expected error error when initialising manager onto an existing encryption key")
	} else if err.Error() != "cipher: message authentication failed" {
		t.Error("Expected a decryption error however recieved different error")
	}
}

func TestValidPasswordWorks(t *testing.T) {
	c = &TestCacher{fresh: true}

	ok, err := Init(TEST_PASSWORD)
	if !ok {
		t.Fatalf("Error when initialising manager: %s", err.Error())
	}

	testKey := "testKey"
	testValue := "testVal"
	err = InsertEntry(testKey, testValue); if err != nil {
		t.Fatalf("Error when inserting key: %s", err.Error())
	}

	val, err := RetrieveEntry(testKey)
	if err != nil {
		t.Fatalf("Error when retrieving key: %s", err.Error())
	} else if val != testValue {
		t.Fatal("Decryption failed to match expected entry")
	}
}