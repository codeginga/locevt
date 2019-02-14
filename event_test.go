package locevt_test

import (
	"sync"
	"testing"
	"time"

	"github.com/codeginga/locevt"
)

type customData struct {
	id int
}

func TestBus(t *testing.T) {
	e := locevt.NewEvent()

	runTest := 200
	retry := [200]int{}
	mxRetry := 5

	wg := sync.WaitGroup{}

	f := func(task locevt.Task) {
		time.Sleep(time.Millisecond * 100)

		dt := task.Data().(customData)
		retry[dt.id]++

		if retry[dt.id] == mxRetry {
			wg.Done()
			return
		}

		task.Retry()

	}

	e.Register("test1", f)
	e.Register("test", f)

	for i := 0; i < runTest; i++ {
		wg.Add(1)
		if i%2 == 0 {
			e.Fire(locevt.FireOption{
				Name:      "test",
				Data:      customData{id: i},
				MaxRetry:  mxRetry,
				RetryWait: time.Millisecond * 100,
			})
			continue
		}
		e.Fire(locevt.FireOption{
			Name:      "test1",
			Data:      customData{id: i},
			MaxRetry:  mxRetry,
			RetryWait: time.Millisecond * 100,
		})

	}

	wg.Wait()

	for i := 0; i < runTest; i++ {
		if retry[i] != mxRetry {
			t.Errorf("i=%d: retry is not mxretry, retry=%d ", i, retry[i])
		}
	}
}
