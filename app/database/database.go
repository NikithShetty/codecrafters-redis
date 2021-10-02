package database

type RedisStore struct {
	kv map[string]string
}

func InitDatabase() *RedisStore {
	return &RedisStore{kv: make(map[string]string)}
}

func (db *RedisStore) Set(k string, v string) {
	db.kv[k] = v
}

func (db *RedisStore) Get(k string) string {
	return db.kv[k]
}
