package IEncryption

type EncryptorFactory interface {
	Encryptor

	// Extend  encryption factory with given key and encryptor.
	Extend(key string, encryptor Encryptor)

	// Driver Get the encryptor by the given key.
	Driver(key string) Encryptor
}

type Encryptor interface {

	// Encode Encrypt the given value.
	Encode(value string) string

	// Decode Decrypt the given value.
	Decode(encrypted string) (string, error)
}
