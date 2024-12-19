package mutex

import (
	"log"
	"sync"
	"time"
)

type Mutex struct{}

func (m *Mutex) Name() string {
	return "mutex"
}

func (m *Mutex) Run() {
	log.Println(m.Name(), "run")
	mutex := &sync.Mutex{}
	mutex.Lock()
	go func() {
		time.Sleep(1 * time.Second)
		mutex.Unlock()
	}()
	mutex.Lock() // 一定会造成死锁
}
