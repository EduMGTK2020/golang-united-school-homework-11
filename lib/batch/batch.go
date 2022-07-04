package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {
	sem := make(chan struct{}, pool)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var i int64

	for i = 0; i < n; i++ {
		sem <- struct{}{}
		wg.Add(1)
		go func(j int64) {
			user := getOne(j)
			mu.Lock()
			res = append(res, user)
			mu.Unlock()
			wg.Done()
			<-sem
		}(i)
	}
	wg.Wait()
	return
}
