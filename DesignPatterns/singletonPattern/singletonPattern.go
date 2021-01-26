package singletonPattern

import "sync"

var once sync.Once

type Singleton struct {
	Count int
}

var singleton *Singleton

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			Count: 0,
		}
	})
	return singleton
}

func GetInstanceNotOnce() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			Count: 0,
		}
	})
	return singleton
}