[![Go Report Card](https://goreportcard.com/badge/github.com/codeginga/locevt)](https://goreportcard.com/report/github.com/codeginga/locevt)

# locevt
**locevt** manages the event locally. Register a function (worker) by name to listen to the event.  Fire event by name with data.

**Features:**
1. Retry mechanism
2. Retry with delay

## How to use:
```go
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

// define worker
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


func main() {
	// init event instance
	evt := locevt.NewEvent()

	// registering worker
	evt.Register("user-notification", worker)

	// example to fire event
	evt.Fire(locevt.FireOption{
		Name:      "user-notification",
		Data:      User{Name: "test name", Phone: "test phone"},
		RetryWait: time.Microsecond * 50,
		MaxRetry:  2,
	})

	
}

```