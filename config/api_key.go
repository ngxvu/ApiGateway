package config

import (
	"os"
	"sync"
)

var (
	aServiceApiKey string
	onceApiKey     sync.Once
)

func GetServiceAKey() string {
	onceApiKey.Do(func() {
		aServiceApiKey = os.Getenv("SERVICE_A_API_KEY")
		if aServiceApiKey == "" {
			panic("failed to get service a key")
		}
	})
	return aServiceApiKey
}
