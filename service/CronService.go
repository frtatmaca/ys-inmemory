package service

import (
	"sync"
	"time"
)

type CronService struct {
	inMemoryService InMemoryService
}

func NewCronService(inMemoryService InMemoryService) CronService {
	return CronService{inMemoryService}
}

func (p *CronService) StartChrone() {
	var lock sync.Mutex
	timer1 := time.NewTicker(time.Second * 10)
	defer timer1.Stop()
	for {
		select {
		case <-timer1.C:
			go func() {
				lock.Lock()
				defer lock.Unlock()
				p.inMemoryService.SetMemoryToStore()
			}()
		}
	}
}
