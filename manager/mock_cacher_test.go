package manager

import "errors"

type TestCacher struct {
	fresh bool
}

var insertedKey string
var insertedVal []byte

func (tc *TestCacher) InitCacher() {}

func (tc *TestCacher) InsertKey(key string, val []byte) error {
	insertedKey = key
	insertedVal = val
	return nil
}

func (tc *TestCacher) DetermineIfExists(key string) (bool, error) {
	if !tc.fresh {
		if key == GENESIS_KEY {
			return true, nil
		}
	}

	if key == insertedKey {
		return true, nil
	}
	return false, nil
}

func (tc *TestCacher) RetrieveEntry(key string) ([]byte, error) {
	if !tc.fresh {
		if key == GENESIS_KEY {
			return []byte(GENESIS_VAL + " encrypted differently"), nil
		}
	}

	if key == insertedKey {
		return insertedVal, nil
	} else {
		return []byte(""), errors.New("Key does not exist")
	}
}