package sessions

import (
	"brevity/repository"
	"brevity/service"
	"brevity/service_abstract"
)

// Stores session data for the user.
// A Session stores the services for the current user.
type Session struct {
	running       bool
	id            uint64  // Session id
	ScheduableSrv service_abstract.AbsScheduableService
	DepSrv        service_abstract.AbsDependencyService
}

// Create a new session.
// TODO: Mechanism for using any type of RepositoryFactory
func NewSession(id uint64) *Session {
	factory := service.NewLocalServiceFactory(repository.NewInMemoryRepositoryFactory())

	scheduableSrv := factory.ScheduableService()
	depSrv := factory.DependencyService()

	return &Session{id: id, DepSrv: depSrv, ScheduableSrv: scheduableSrv, running: true}
}

// Returns the id of the Session.
func (s *Session) GetId() uint64 {
	return s.id
}

// Stops the Session.
func (s *Session) Stop() {
	s.running = false
}
