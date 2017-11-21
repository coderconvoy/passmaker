package pswd

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/scrypt"
)

type Password struct {
	Hash, Salt []byte
}

func Equal(a, b Password) bool {
	if !bytes.Equal(a.Hash, b.Hash) {
		return false
	}
	if !bytes.Equal(a.Salt, b.Salt) {
		return false
	}
	return true
}

func New(pw string) (Password, error) {
	salt := make([]byte, 10)
	_, err := rand.Read(salt)
	if err != nil {
		return Password{}, err
	}
	dk, err := scrypt.Key([]byte(pw), salt, 16384, 8, 1, 32)
	return Password{dk, salt}, nil

}
func Parse(ps string) (Password, error) {
	ss := strings.Split(strings.TrimSpace(ps), "_")
	if len(ss) != 2 {
		return Password{}, errors.New("Password should have 2 parts")
	}

	salt, err := hex.DecodeString(ss[0])
	if err != nil {
		return Password{}, fmt.Errorf("Salt not decoded : %s", ss[0])
	}

	hash, err := hex.DecodeString(ss[1])
	if err != nil {
		return Password{}, fmt.Errorf("Hash Not decoded : %s", ss[1])
	}

	return Password{Salt: salt, Hash: hash}, nil
}

func (p Password) Check(pw string) bool {
	if pw == "" {
		return false
	}
	dk, err := scrypt.Key([]byte(pw), p.Salt, 16384, 8, 1, 32)
	if err != nil {
		return false
	}
	return string(dk) == string(p.Hash)
}

func (p Password) String() string {
	return fmt.Sprintf("%x_%x", p.Salt, p.Hash)
}
