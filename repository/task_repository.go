package repository

import (
	"brevity/task"
	"github.com/emirpasic/gods/lists/arraylist"
)

// Common interface for all task.Task Repositories
// Responsible for storing the Scheduable date
// TODO: update method.
type TaskRepository interface {
	Retrieve(id uint64) (task.Scheduable, error)
	Add(scheduable task.Scheduable)
	Remove(id uint64) error
	Find(id uint64) bool
	RemoveAll()
	GetAll() *arraylist.List  // Get the list of all the items.
	Size() int
}
