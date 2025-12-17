package locker

type Locker interface {
	Get(name string) (string, error)
	GetRaw(name string) (string, error)
	List() ([]string, error)
}

type LockerWr interface {
	Locker
	Put(name string, value string) error
	Delete(name string) error
}

type Decryptor interface {
	Decrypt(data string) (string, error)
	Encrypt(data string) (string, error)
}
