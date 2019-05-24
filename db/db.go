package db

type Db interface {
	Get(key string) (int, error)
    Set(key string, value int) error
    Name() string
}
