package manager

import (
	"example/secure-stash/manager/cacher"
	"example/secure-stash/manager/stasher"
)

var (
	c cacher.Cacher = cacher.Cacher{}
	st stasher.Stasher = stasher.Stasher{}
)

/*
 * Initialise cache connection and encryption settings
 */
func Init(password string) bool {
	c.InitCacher()
	st.InitStasher(stasher.RawToHash(password))
	if st.HasIntegrity(password) {
		return true
	}

	return false
}

/*
 * Encrypt then insert entry into redis cache
 */
func InsertEntry(key, value string) error {
	stashedValue := st.EncryptText(value)
	err := c.InsertKey(key, stashedValue)
	return err
}

/*
 * Decrypt then return an entry from redis cache
 */
func RetrieveEntry(key string) (string, error) {
	stashedValue, err := c.RetrieveKey(key)
	if err != nil {
		return "", err
	}

	unstashedValue, err := st.DecryptBytes(stashedValue)
	if err != nil {
		return "", err
	}

	return unstashedValue, nil
}