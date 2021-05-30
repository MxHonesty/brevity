package repository

// Abstract factory for creating Repositories.
type Factory interface {
	CreateTaskRepository() TaskRepository
	CreateDependencyRepository() DependencyRepository
}

// Implements the Abstract Factory.
// Creates in memory versions of the repositories.
type InMemoryRepositoryFactory struct {}

// Creates a new InMemoryRepositoryFactory.
func NewInMemoryRepositoryFactory() *InMemoryRepositoryFactory {
	return &InMemoryRepositoryFactory{}
}

// Create an in memory version of the TaskRepository.
func (fac *InMemoryRepositoryFactory) CreateTaskRepository() TaskRepository {
	repo := NewScheduleRepository()
	return repo
}

// Create a in memory version of the DependencyRepository.
func (fac *InMemoryRepositoryFactory) CreateDependencyRepository() DependencyRepository {
	repo := NewDepRepository()
	return repo
}
