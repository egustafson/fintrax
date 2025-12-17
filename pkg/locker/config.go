package locker

type LockerConfig struct {
	Decryptor *DecryptorConfig `yaml:"decryptor"`
	Store     *StoreConfig     `yaml:"store"`
}
