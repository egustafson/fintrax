package locker

type yubiKeyDecryptor struct {
	slot int
	pin  string
}

var _ Decryptor = (*yubiKeyDecryptor)(nil)

func MakeYubiKeyDecryptor(config *YubiKeyDecryptorConfig) (Decryptor, error) {
	return &yubiKeyDecryptor{
		slot: config.Slot,
		pin:  config.PIN,
	}, nil
}

func (y *yubiKeyDecryptor) Decrypt(encryptedValue string) (string, error) {
	//
	// TODO: implement YubiKey decryption logic
	//
	// Placeholder implementation for YubiKey decryption
	// In a real implementation, this would interact with the YubiKey hardware
	decryptedValue := "decrypted:" + encryptedValue // Mock decryption
	return decryptedValue, nil
}

func (y *yubiKeyDecryptor) Encrypt(value string) (string, error) {
	//
	// TODO: implement YubiKey encryption logic
	//
	// Placeholder implementation for YubiKey encryption
	// In a real implementation, this would interact with the YubiKey hardware
	encryptedValue := "encrypted:" + value // Mock encryption
	return encryptedValue, nil
}
