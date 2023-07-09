package singleton

import (
	"fmt"
	"sync"
)

type SingletonStruct struct{}

var uniqueInstance *SingletonStruct

type SingletonIface interface {
	GetInstance() *SingletonStruct
}

func NewSingleTonThreadUnsafe() SingletonIface {
	return &SingletonThreadUnsafe{}
}

func NewSingleTonThreadSafe() SingletonIface {
	return &SingletonThreadSafe{}
}

type SingletonThreadUnsafe struct {
}

func (ins *SingletonThreadUnsafe) GetInstance() *SingletonStruct {
	if uniqueInstance == nil {
		fmt.Println("Instantiating singleton struct")
		uniqueInstance = &SingletonStruct{}
	}
	return uniqueInstance
}

type SingletonThreadSafe struct{}

var lock = &sync.Mutex{}

func (ins *SingletonThreadSafe) GetInstance() *SingletonStruct {
	lock.Lock()
	defer lock.Unlock()
	if uniqueInstance == nil {
		fmt.Println("Instantiating singleton struct")
		uniqueInstance = &SingletonStruct{}
	}
	return uniqueInstance
}
