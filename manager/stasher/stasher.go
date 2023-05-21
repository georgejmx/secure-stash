package stasher

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

type Stasher struct {
	galoisNonce []byte
	galoisCipher cipher.AEAD
}

const NONCE_SIZE = 12

func (st *Stasher) InitStasher(seed string) error {
	if len(seed) != 32 {
		return errors.New("initial seed not 32 characters long")
	} else if NONCE_SIZE > 32 {
		return errors.New("nonce size too large")
	}

	aesCipher, err := aes.NewCipher([]byte(seed))
	if err != nil {
		return err
	}

	st.galoisCipher, err = cipher.NewGCM(aesCipher)
	if err != nil {
		return err
	}
	st.galoisNonce = []byte(seed[:NONCE_SIZE]) // TODO: change to an independent value

	return nil
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