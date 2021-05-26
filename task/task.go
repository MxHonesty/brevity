package task

import "time"

type Task struct {
	title string
	description string
	id uint64

	// The end and start time of the task.
	start time.Time
	end time.Time
}

/* Returns a Task. */
func NewTask(id uint64, title, description string, start, end time.Time) Task {
	return Task{id: id, title: title, description: description,
		start: start, end: end}
}

/* Returns the title of the Task. */
func (t *Task) GetTitle() string {
	return t.title
}

/* Returns the description of the Task. */
func (t *Task) GetDescription() string {
	return t.description
}

// Returns the id of the Task.
func (t *Task) GetId() uint64 {
	return t.id
}

func (t *Task) GetStartTime() time.Time {
	return t.start
}

func (t *Task) GetEndTime() time.Time {
	return t.end
}

func (t *Task) GetDuration() time.Duration {
	return t.end.Sub(t.start)
}
