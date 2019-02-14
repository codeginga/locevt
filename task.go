package locevt

import (
	"errors"
	"time"

	"github.com/sirupsen/logrus"
)

type task struct {
	fireOption FireOption
	event      Event
}

func (t *task) Data() interface{} {
	return t.fireOption.Data
}

func (t *task) Retry() error {
	if t.fireOption.MaxRetry <= 0 {
		return errors.New("max retry is <= 0")
	}

	go func() {
		time.Sleep(t.fireOption.RetryWait)
		if err := t.event.Fire(t.fireOption); err != nil {
			logrus.Warn("could not fire event from task.Retry, name="+t.fireOption.Name+", err=", err)
		}
	}()

	return nil
}
