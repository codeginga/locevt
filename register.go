package locevt

import (
	"errors"
	"sync"
)

type register struct {
	workerMap map[string]Worker
	mu        sync.RWMutex
}

func (r *register) add(name string, w Worker) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.workerMap[name] = w
}

func (r *register) worker(name string) (Worker, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	w, ok := r.workerMap[name]
	if !ok {
		return nil, errors.New("no worker found for event " + name)
	}

	return w, nil
}
