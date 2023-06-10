package manager

import (
	"errors"
	"example/secure-stash/manager/cacher"
	"example/secure-stash/manager/stasher"
	"example/secure-stash/manager/utils"
)

const (
	GENESIS_KEY = "$INIT"
	GENESIS_VAL = "kvasir"
)

var (
	c cacher.CacherTemplate = &cacher.Cacher{}
	st stasher.Stasher = stasher.Stasher{}
	isSetup bool = false
)

// Validate an established password using the genesis key value pair to test
// decryption
func isPasswordEstablished() (bool, error) {
	entry, err := RetrieveEntry(GENESIS_KEY)
	if err != nil || entry != GENESIS_VAL {
		return false, err
	}
	return true, nil
}

// Initialise cache connection and encryption settings and genesis. Returns an
// error if the integrity of the cache cannot be established e.g. the genesis
// pair cannot be retrieved or decryption of the gensis value fails
func Init(password string) (bool, error) {
	c.InitCacher()
	hashedPassword := utils.RawToHash(password)
	st.InitStasher(hashedPassword)

	isSetup, err := c.DetermineIfExists(GENESIS_KEY)
	if err != nil {
		return false, err
	} else if !isSetup {
		InsertEntry(GENESIS_KEY, GENESIS_VAL)
	}

	if ok, err := isPasswordEstablished(); !(ok && err == nil) {
		return false, err
	}

	return true, nil
}

// Retrieve all keys in the cache
func RetrieveEntries() ([]string, error) {
	keys, err := c.RetrieveEntries()
	if err != nil {
		return nil, err
	}

	return utils.Remove(keys, GENESIS_KEY), nil
}

// Encrypt then insert entry into redis cache
func InsertEntry(key, value string) error {
	if isSetup && key == GENESIS_KEY {
		return errors.New("Cannot use the reserved key: " + GENESIS_KEY)
	}

	stashedValue := st.EncryptText(value)
	err := c.InsertKey(key, stashedValue)
	return err
}

// Decrypt then return an entry from redis cache
func RetrieveEntry(key string) (string, error) {
	stashedValue, err := c.RetrieveEntry(key)
	if err != nil {
		return "", err
	}

	unstashedValue, err := st.DecryptBytes(stashedValue)
	if err != nil {
		return "", err
	}

	return unstashedValue, nil
}
