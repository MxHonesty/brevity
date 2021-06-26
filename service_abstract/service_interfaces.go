// This module defines common interfaces for Service implementations. It was
// created as a pure fabrication to eliminate cyclical imports.
package service_abstract

import (
	"brevity/dependency"
	"brevity/task"
	"time"
)

// Interface for ScheduableService. Declares all common service operations.
type AbsScheduableService interface {
	AddContainer(startYear int, startMonth time.Month, startDay, startHour, startMin int,
		endYear int, endMonth time.Month, endDay, endHour, endMin int)
	RemoveContainer(id uint64) bool
	GetAllContainers() []task.Container
}

// Interface for DependencyService. Declares all common service operations.
type AbsDependencyService interface {
	AddDependency(dependentId uint64, dependentOnId ...uint64) error
	RemoveDependency(id uint64) bool
	GetAllDependencies() []dependency.Dependency
}
