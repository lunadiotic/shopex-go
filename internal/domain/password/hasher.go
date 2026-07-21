package password

type Hasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) (bool, error)
}