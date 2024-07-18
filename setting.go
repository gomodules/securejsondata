package securejsondata

import (
	"sync"
)

var (
	secretKey     string
	keyConfigured bool
	mu            sync.Mutex
)

func InitSecretKey(key string) {
	mu.Lock()
	defer mu.Unlock()
	if keyConfigured {
		panic("secret key is already configured")
	}

	secretKey = key
	keyConfigured = true
}

func SecretKey() string {
	mu.Lock()
	defer mu.Unlock()
	if !keyConfigured {
		panic("secret key is not configured")
	}
	return secretKey
}
