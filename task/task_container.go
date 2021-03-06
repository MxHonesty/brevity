package task

import (
	"time"
)

// Structure responsible for holding related tasks
// It is responsible for holding time information.
// such as the start and end time.
// It is compatible with the Scheduable interface.
type Container struct {
	// Unique id of the Container.
	id uint64

	// Start time of this Container.
	start time.Time

	// End time of this Container.
	end time.Time

	// A Slice of Tasks that are appended to this Container.
	tasks []Task
}

// Creates Container.
func NewContainer(id uint64, start, end time.Time) *Container {
	return &Container{id: id, start: start, end: end}
}

// Returns a deep copy of a Container.
func (c *Container) Copy() Scheduable {
	newEl := NewContainer(c.id, c.start, c.end)
	sliceCopy := make([]Task, len(c.GetTasks()))
	copy(sliceCopy, c.GetTasks())
	newEl.SetTasks(sliceCopy)
	return newEl
}

// Get the id of the Container.
func (c *Container) GetId() uint64 {
	return c.id
}

// Get the duration of the Container as a time.Duration.
func (c *Container) GetDuration() time.Duration {
	return c.end.Sub(c.start)
}

// Get the start time.
func (c *Container) GetStartTime() time.Time {
	return c.start
}

// Get the end time.
func (c *Container) GetEndTime() time.Time {
	return c.end
}

// Return a copy of the slice of Task instances
// appended to this Container.
func (c *Container) GetTasks() []Task {
	temp := make([]Task, len(c.tasks))
	copy(temp, c.tasks)
	return temp
}

func (c *Container) SetTasks(newTasks []Task) {
	temp := make([]Task, len(newTasks))  // Copy elements of the slice passed
	copy(temp, newTasks)
	c.tasks = temp
}

// Set the start time as a time.Time.
func (c *Container) SetStartTime(newStart time.Time) {
	c.start = newStart
}

// Set the end time of the Container as a time.Time.
func (c *Container) SetEndTime(newEnd time.Time) {
	c.end = newEnd
}
