package db

type MemDb struct {
    data map[string]int
}

func NewMemDb() Db {
    return &MemDb {
        data: make(map[string]int),
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
