package task

// Holds the data for a single Task.
type Task struct {
	title string
	description string
}

// Returns a Task.
func NewTask(title, description string) Task {
	return Task{title: title, description: description}
}

// Returns the title of the Task.
func (t *Task) GetTitle() string {
	return t.title
}

// Returns the description of the Task.
func (t *Task) GetDescription() string {
	return t.description
}

// Set the description of the Task.
func (t *Task) SetDescription(newDescription string) {
	t.description = newDescription
}

// Set the title of the Task.
func (t *Task) SetTitle(newTitle string) {
	t.title = newTitle
}
