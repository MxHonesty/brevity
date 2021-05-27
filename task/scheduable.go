package task

import "time"

// Interface for all items that can be added to the Schedule.
// It needs a to be able to return a duration and also
// have access to the start time and end time in the time
// package format.
// Every Scheduable item needs an unique id.
type Scheduable interface {
	GetId() uint64
	GetDuration() time.Duration
	GetStartTime() time.Time
	GetEndTime() time.Time
	GetTasks() []Task  // Get the tasks of this item.
	SetTasks(newTasks []Task)
	Copy() Scheduable
}
