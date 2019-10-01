package db

type Db interface {
	Get(key string) (int, error)
	Set(key string, value int) error
	GetBytes(key string) ([]byte, error)
	SetBytes(key string, bytes []byte) error
	Name() string
}
