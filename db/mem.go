package db

type MemDb struct {
	data  map[string]int
	bytes map[string][]byte
}

func NewMemDb() Db {
	return &MemDb{
		data:  make(map[string]int),
		bytes: make(map[string][]byte),
	}
}

func (db *MemDb) Get(key string) (int, error) {
	return db.data[key], nil
}

func (db *MemDb) Set(key string, value int) error {
	db.data[key] = value
	return nil
}

func (db *MemDb) Name() string {
	return "mem"
}
func (db *MemDb) GetBytes(key string) ([]byte, error) {
	return db.bytes[key], nil
}

func (db *MemDb) SetBytes(key string, bytes []byte) error {
	db.bytes[key] = bytes
	return nil
}
