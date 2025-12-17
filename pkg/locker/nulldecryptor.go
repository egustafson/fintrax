package locker

type NullDecryptor struct{}

func MakeNullDecryptor(_ *NullDecryptorConfig) *NullDecryptor {
	return &NullDecryptor{}
}

func (n NullDecryptor) Decrypt(data string) (string, error) {
	return data, nil
}

func (n NullDecryptor) Encrypt(data string) (string, error) {
	return data, nil
}
