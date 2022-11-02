package singleton

import (
	"sync"
)

type Singleton struct {
}

var s *Singleton

// GetIns 懒汉模式，非线程安全
func GetIns() *Singleton {
	if s == nil {
		s = &Singleton{}
	}
	return s
}

var mu sync.Mutex

// GetIns1 懒汉模式加锁，线程安全
func GetIns1() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if s == nil {
		s = &Singleton{}
	}
	return s
}

var once sync.Once

// GetIns2 sync.once实现
func GetIns2() *Singleton {
	once.Do(func() {
		s = &Singleton{}
	})
	return s
}

var s2 *Singleton = &Singleton{}

// GetIns3 饿汉
func GetIns3() *Singleton {
	return s2
}
