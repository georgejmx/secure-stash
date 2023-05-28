package manager

import "testing"

const TEST_PASSWORD = "something!secureY"


/*
 * Tests that initialising a new cache onto a redis instance that has already
 * been setup with different encryption keys i.e. a different value of
 * GENESIS_VAL will fail authentication. This corresponds to an invalid password
 * login
 */
func TestInvalidPasswordFails(t *testing.T) {
	c = &TestCacher{fresh: false}

	ok, err := Init(TEST_PASSWORD)
	if ok {
		t.Fatal("Success when initialising manager onto an existing encryption key")
	} else if err.Error() != "cipher: message authentication failed" {
		t.Error("Expected a decryption error however recieved different error")
	}
}

/*
 * Tests that a valid password login allows encryption and decryption of data
 */
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