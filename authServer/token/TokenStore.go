package token

import (
	"log"
	"time"

	"github.com/partha-sen/ostd/authServer/config"
)

type TokenStore map[string]int64

var GlobalTokenStore TokenStore

func (store *TokenStore) Add(key string, value int64) {
	GlobalTokenStore[key] = value
}

func (store *TokenStore) Contains(key string) bool {
	_, ok := (*store)[key]
	return ok
}

func (store *TokenStore) Remove(key string) {
	delete(GlobalTokenStore, key)
}

func RemoveExpiredToken() {
	log.Println("Removing expired token..")
	now := time.Now().Unix()
	for key, expiry := range GlobalTokenStore {
		if now > expiry {
			delete(GlobalTokenStore, key)
		}
	}
	time.Sleep(config.TIME_INTERVAL * time.Second)
	RemoveExpiredToken()
}
