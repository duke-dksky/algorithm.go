package main

import (
	"fmt"
	"sync"
)

type Manager struct {
	name string
}

func newManager(name string) *Manager {
	return &Manager{name}
}

var once sync.Once
var manager *Manager

func GetManger() *Manager {
	once.Do(func() {
		manager = newManager("Singleton")
	})
	return manager
}

func main() {
	mgr := GetManger()
	fmt.Println(mgr.name)
	mgr1 := GetManger()
	fmt.Println(mgr1.name)
}
