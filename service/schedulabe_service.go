package service

import (
	"brevity/repository"
	"brevity/task"
	"fmt"
	"time"
)

type ScheduableService struct {
	repo      repository.TaskRepository
	currentId uint64
}

// Creates a service.
// Receives a repository.Factory as a parameter.
// The Constructor uses it as a way to create the needed repositories.
func NewScheduableService(factory repository.Factory) *ScheduableService {
	rep := factory.CreateTaskRepository()
	srv := ScheduableService{repo: rep, currentId: 0}
	return &srv
}

// Adds the created Container to the repository.
func (srv *ScheduableService) AddContainer(startYear int, startMonth time.Month, startDay, startHour, startMin int,
	endYear int, endMonth time.Month, endDay, endHour, endMin int) {
	startDate := time.Date(startYear, startMonth, startDay, startHour, startMin, 0, 0, time.Local)
	endDate := time.Date(endYear, endMonth, endDay, endHour, endMin,0 , 0, time.Local)

	item := task.NewContainer(srv.currentId, startDate, endDate)
	srv.currentId++  // Increment id.
	srv.repo.Add(item)
}

// Removes the container with the given id.
// Returns true if the item was removed.
// Returns false if the item was not found.
func (srv *ScheduableService) RemoveContainer(id uint64) bool {
	err := srv.repo.Remove(id)
	if err != nil {
		return false
	}
	return true
}

// Iterates over the elements and prints them
func (srv *ScheduableService) PrintContainers() {
	vec := srv.repo.GetAll()
	it := vec.Iterator()
	for it.Next() {
		el := it.Value().(*task.Container)  // task.Container type check
		fmt.Println(el)
	}
}

func (srv *ScheduableService) GetAllContainers() []task.Container {
	vec := srv.repo.GetAll()
	slc := vec.Values()
	temp := make([]task.Container, 0)  // Temporary slice for copy

	for _, el := range slc {
		newEl := el.(task.Scheduable).Copy()  // Type asserting to Container
		// We copy the element.
		containerEl := newEl.(*task.Container)
		temp = append(temp, *containerEl)
	}

	return temp
}
