package database

import (
	"time"

	"codecrafters-redis/app/utils"
)

type storeValue struct {
	value  string
	expiry *time.Time
}

type RedisStore struct {
	kv map[string]*storeValue
}

func InitDatabase() *RedisStore {
	return &RedisStore{kv: make(map[string]*storeValue)}
}

func (db *RedisStore) Set(k string, v string) {
	utils.LogInfo("set key", k)
	sv := storeValue{value: v}
	db.kv[k] = &sv
}

func (db *RedisStore) SetPx(k string, v string, ex time.Duration) {
	utils.LogInfo("setpx key", k)
	utils.LogInfo("setpx duration", ex.Milliseconds())
	fut := time.Now().Add(ex)
	sv := storeValue{value: v, expiry: &fut}
	db.kv[k] = &sv
}

func (db *RedisStore) Get(k string) *string {
	utils.LogInfo("get key", k)
	sv := db.kv[k]
	if sv == nil {
		return nil
	} else if sv.expiry != nil && (*sv.expiry).Before(time.Now()) {
		utils.LogInfo("Get", sv.value)
		// Expired
		db.kv[k] = nil
		return nil
	}
	return &sv.value
}
