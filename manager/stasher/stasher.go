package stasher

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/binary"
	"math/rand"
)

type Stasher struct {
	galoisNonce []byte
	galoisCipher cipher.AEAD
}

const NONCE_SIZE = 12

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

/*
 * Generates two byte arrays, that correspond to a seed and nonce, from an input
 * byte array. Note that this function is one-way reproducable and random
 */
 func hashToSecrets(hash [32]byte) ([NONCE_SIZE]byte, [32]byte) {
	copy := hash
	hashes := [2][32]byte{hash, copy}
	derivedSeeds := [2]int64{int64(binary.BigEndian.Uint64(hash[0:16])), int64(binary.BigEndian.Uint64(hash[16:32]))}
	
	for i := 0; i < 2; i++ {
		rand.Seed(int64(derivedSeeds[i]))
		rand.Shuffle(32, func(x, y int){
			hashes[i][x], hashes[i][y] = hashes[i][y], hashes[i][x]
		})
	}
	return ([NONCE_SIZE]byte)(hashes[0][0:NONCE_SIZE]), hashes[1]
}

/*
 * Converts a string password to hashed bytes
 */
 func RawToHash(raw string) [32]byte {
	return sha256.Sum256([]byte(raw))
}

func (st *Stasher) EncryptText(plaintext string) []byte {
	return st.galoisCipher.Seal(st.galoisNonce, st.galoisNonce, []byte(plaintext), nil)
}

func (st *Stasher) DecryptBytes(blob []byte) (string, error) {
	nonce, ciphertext := blob[:NONCE_SIZE], blob[NONCE_SIZE:]
	plaintext, err := st.galoisCipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}