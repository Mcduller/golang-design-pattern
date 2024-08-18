package singleton_szl

import (
	"fmt"
	"sync"
)

type Singleton interface {
	Foo() string
}

type singleton struct {
}

func (s singleton) Foo() string {
	return fmt.Sprintf("Singleton Szl")
}

var (
	instance *singleton
	once     sync.Once
)

func NewSingleton() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
