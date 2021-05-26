package repository

import (
	"brevity/task"
	"errors"
)

type ScheduleRepository struct {
	elements []task.Scheduable
}

// Creates an empty ScheduleRepository
func NewScheduleRepository() ScheduleRepository {
	return ScheduleRepository{[]task.Scheduable{}}
}

func (repo *ScheduleRepository) Retrieve(id uint64) (*task.Scheduable, error) {
	for i, el := range repo.elements {
		if el.GetId() == id {
			return &repo.elements[i], nil // Return pointer to element
		}
	}
	return nil, errors.New("no element with id found")
}

// Adds the Scheduable item to the repository.
func (repo *ScheduleRepository) Add(scheduable task.Scheduable) {
	repo.elements = append(repo.elements, scheduable)  // Adds copy of element to list
}

// Returns true if the item with the given id is found.
func (repo *ScheduleRepository) Find(id uint64) bool {
	for _, el := range repo.elements {
		if el.GetId() == id {
			return true
		}
	}
	return false
}

// Removes all the elements.
func (repo *ScheduleRepository) RemoveAll() {
	repo.elements = nil
}

// Returns the array of elements.
func (repo *ScheduleRepository) GetAll() []task.Scheduable {
	return repo.elements
}
