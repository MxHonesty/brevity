package repository

import "brevity/task"

// Common interface for all Repositories
// Responsible for storing the Scheduable date
// TODO: add remove item by id.
type Repository interface {
	Retrieve(id uint64) (*task.Scheduable, error)
	Add(scheduable task.Scheduable)
	Find(id uint64) bool
	RemoveAll()
	GetAll() []task.Scheduable

}
