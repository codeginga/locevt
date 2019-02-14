package locevt

import "time"

// Worker defines Worker func
type Worker func(Task)

// Task wraps task functionality
type Task interface {
	Data() interface{}
	Retry() error
}

// FireOption holds data to fire event
type FireOption struct {
	Name      string
	Data      interface{}
	MaxRetry  int
	RetryWait time.Duration
}

// Event wraps evnet functionality
type Event interface {
	Fire(option FireOption) error
	Register(name string, w Worker)
}
