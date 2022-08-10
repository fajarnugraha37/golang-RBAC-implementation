package util

import (
	"encoding/json"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

func Singleton[T any](instance *T, initialize func() *T) *T {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = initialize()
		}
	}
	return instance
}

func LogAny(result interface{}) {
	s, _ := json.MarshalIndent(result, "", "\t")
	log.Println(string(s))
}

func Ternary[T any](condition bool, trueItem T, falseItem T) T {
	return map[bool]T{true: trueItem, false: falseItem}[condition]
}
