package stasher

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"math/rand"
)

type Stasher struct {
	galoisNonce  []byte
	galoisCipher cipher.AEAD
}

const NONCE_SIZE = 12

// Initialise a stasher object derived from the input hash to perform later
// encryption and decryption using this hash
func (st *Stasher) InitStasher(inputHash [32]byte) error {
	nonce, seed := hashToSecrets(inputHash)

	aesCipher, err := aes.NewCipher(seed[:])
	if err != nil {
		return err
	}

	st.galoisCipher, err = cipher.NewGCM(aesCipher)
	if err != nil {
		return err
	}
	st.galoisNonce = nonce[:]
	return nil
}

// Generates two byte arrays, that correspond to a seed and nonce, from an input
// byte array. Note that this function is one-way reproducable and random
func hashToSecrets(hash [32]byte) ([NONCE_SIZE]byte, [32]byte) {
	copy := hash
	hashes := [2][32]byte{hash, copy}
	derivedSeeds := [2]int64{
		int64(binary.BigEndian.Uint64(hash[0:16])),
		int64(binary.BigEndian.Uint64(hash[16:32])),
	}

	for i := 0; i < 2; i++ {
		rand.Seed(int64(derivedSeeds[i]))
		rand.Shuffle(32, func(x, y int) {
			hashes[i][x], hashes[i][y] = hashes[i][y], hashes[i][x]
		})
	}
	fixedSeedReference := (*[NONCE_SIZE]byte)(hashes[0][0:NONCE_SIZE])
	return *fixedSeedReference, hashes[1]
}

// Perform encryption of plaintext to encrypted bytes
func (st *Stasher) EncryptText(plaintext string) []byte {
	return st.galoisCipher.Seal(st.galoisNonce, st.galoisNonce, []byte(plaintext), nil)
}

// Perform decryption of encrypted bytes to plaintext
func (st *Stasher) DecryptBytes(blob []byte) (string, error) {
	nonce, ciphertext := blob[:NONCE_SIZE], blob[NONCE_SIZE:]
	plaintext, err := st.galoisCipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
