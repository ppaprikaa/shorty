package cryptor

type Cryptor interface {
	Encrypt(data []byte) ([]byte, error)
	Decrypt(data []byte) ([]byte, error)
	Verify(encrypted []byte, toVerify []byte) error
}
