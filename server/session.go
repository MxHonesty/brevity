package server

import (
	"brevity/repository"
	"brevity/service"
)

// Stores session data for the user.
// A Session stores the services for the current user.
type Session struct {
	id uint64  // Session id
	scheduableSrv service.AbsScheduableService
	depSrv service.AbsDependencyService
}

// Create a new session.
// TODO: Mechanism for using any type of RepositoryFactory
func NewSession(id uint64) *Session {
	factory := repository.NewInMemoryRepositoryFactory()
	scheduableRepo := factory.CreateTaskRepository()
	depRepo := factory.CreateDependencyRepository()

	scheduableSrv := service.NewScheduableService(scheduableRepo)
	depSrv := service.NewDependencyService(depRepo, scheduableRepo)

	return &Session{id: id, depSrv: depSrv, scheduableSrv: scheduableSrv}
}

// Returns the id of the Session.
func (s *Session) GetId() uint64 {
	return s.id
}
