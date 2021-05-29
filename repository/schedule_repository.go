package repository

import (
	"brevity/task"
	"errors"
	"github.com/emirpasic/gods/lists/arraylist"
)

type ScheduleRepository struct {
	elements arraylist.List
}

// Creates an empty ScheduleRepository
func NewScheduleRepository() *ScheduleRepository {
	return &ScheduleRepository{elements: *arraylist.New()}
}

func (repo *ScheduleRepository) Retrieve(id uint64) (task.Scheduable, error) {
	index, el := repo.elements.Find(func(index int, value interface{}) bool {
		corValue := value.(task.Scheduable)
		return corValue.GetId() == id
	})
	if index == -1 { // Not Found
		return nil, errors.New("element with id not found")
	} else {
		return el.(task.Scheduable), nil
	}
}

// Adds the Scheduable item to the repository.
func (repo *ScheduleRepository) Add(scheduable task.Scheduable) {
	copyVal := scheduable.Copy()
	repo.elements.Add(copyVal) // Adds copy of element to list
}

// Returns true if the item with the given id is found.
func (repo *ScheduleRepository) Find(id uint64) bool {
	index, _ := repo.elements.Find(func(index int, value interface{}) bool {
		corValue := value.(task.Scheduable)
		return corValue.GetId() == id
	})

	return index != -1
}

// Removes all the elements.
func (repo *ScheduleRepository) RemoveAll() {
	repo.elements.Clear()
}

// Returns a copy of the elements
func (repo *ScheduleRepository) GetAll() *arraylist.List {
	return &repo.elements
}

func (repo *ScheduleRepository) Remove(id uint64) error {
	index, _ := repo.elements.Find(func (index int, value interface{}) bool {
		corValue := value.(task.Scheduable)
		return corValue.GetId() == id
	})

	if index == -1 {
		return errors.New("element not found")
	} else {
		repo.elements.Remove(index)
	}
	return nil
}

func (repo *ScheduleRepository) Size() int {
	return repo.elements.Size()
}

func (repo *ScheduleRepository) Update(id uint64, newItem task.Scheduable) error {
	index, el := repo.elements.Find(func(index int, value interface{}) bool {
		corValue := value.(task.Scheduable)
		return corValue.GetId() == id
	})

	if index == -1 {
		return errors.New("element not found")
	} else {
		element := el.(task.Scheduable)
		element.SetTasks(newItem.GetTasks())
		element.SetStartTime(newItem.GetStartTime())
		element.SetEndTime(newItem.GetEndTime())
		repo.elements.Set(index, element)
		return nil
	}
}
