package password

import (
	"crypto/sha256"
)

type password struct {
	hashFunc  HashFunc
	hash      []byte
	Plaintext string
}

func NewPassword(hf HashFunc) *password {
	if hf == nil {
		hf = defaultHashFunc
	}

	return &password{
		hashFunc: hf,
	}
}

func (p *password) Set() error {
	var err error

	if p.hash, err = p.hashFunc(p.Plaintext); err != nil {
		return err
	}

	return nil
}

func (p *password) Hash() []byte { return p.hash[:] }

type HashFunc func(Plaintext string) ([]byte, error)

func defaultHashFunc(Plaintext string) ([]byte, error) {
	hash := sha256.Sum256([]byte(Plaintext))
	return hash[:], nil
}
