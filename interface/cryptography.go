package _interface

type Encryption interface {
	Encrypt(plaintext string) (ciphertext string)
}

type Decryption interface {
	Decrypt(ciphertext string) (plaintext string)
}
