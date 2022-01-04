package singleton

import "sync"

// 饿汉式
type Singleton struct{}

var singleton *Singleton
var lazySingleton *Singleton
var lonce = &sync.Once{}

func init() {
	singleton = &Singleton{}
}

func Instance() *Singleton {
	return singleton
}

// 懒汉式
func LazyInstance() *Singleton {
	if lazySingleton == nil {
		lonce.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
