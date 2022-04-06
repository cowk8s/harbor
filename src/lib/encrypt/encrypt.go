package encrypt

var (
	defaultKeyPath = "/etc/core/key"
)

// Encryptor encrypts or decrypts a strings
type Encryptor interface {
	// Encrypt encrypts plaintext
	Encrypt(string) (string, error)
	// Decrypt decrypts ciphertext
	Decrypt(string) (string, error)
}

// AESEncryptor uses AES to encrypt or decrypt string
type AESEncryptor struct {
	keyProvider KeyProvider
}
