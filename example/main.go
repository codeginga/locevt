package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/codeginga/locevt"
)

// User holds user information
type User struct {
	Name  string
	Phone string
}

func worker(tsk locevt.Task) {
	usr, ok := tsk.Data().(User)
	if !ok {
		log.Println("could not convert to User")
		return
	}

	// do something with usr
	fmt.Println("name", usr.Name)
	fmt.Println("Phone", usr.Phone)

	// to retry
	// err := tsk.Retry()

	// taking time to finish task
	time.Sleep(time.Millisecond * 500)

}

func taskMiddleware(w locevt.Worker, wg *sync.WaitGroup) locevt.Worker {
	f := func(tsk locevt.Task) {
		w(tsk)
		wg.Done()
	}

	return locevt.Worker(f)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	evt := locevt.NewEvent()

	// registering worker
	evt.Register("user-notification", taskMiddleware(worker, &wg))

	// firing task
	evt.Fire(locevt.FireOption{
		Name:      "user-notification",
		Data:      User{Name: "test name", Phone: "test phone"},
		RetryWait: time.Microsecond * 50,
		MaxRetry:  2,
	})

	wg.Wait()
	fmt.Println("Done")
}
