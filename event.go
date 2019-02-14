package locevt

import (
	"sync"
)

// Event manages events
type event struct {
	register register
}

// Fire fires event
func (e *event) Fire(option FireOption) error {

	w, err := e.register.worker(option.Name)
	if err != nil {
		return err
	}

	option.MaxRetry--

	task := task{
		fireOption: option,
		event:      e,
	}

	go w(&task)

	return nil
}

// Register registers worker
func (e *event) Register(name string, w Worker) {
	e.register.add(name, w)
}

// NewEvent returns instance of Event
func NewEvent() Event {
	return &event{
		register: register{
			workerMap: map[string]Worker{},
			mu:        sync.RWMutex{},
		},
	}
}
