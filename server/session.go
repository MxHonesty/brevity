package server

import (
	"brevity/repository"
	"brevity/service"
)

// TODO: Session builder.
// Stores session data for the user.
// A Session stores the services for the current user.
type Session struct {
	id uint64  // Session id
	scheduableSrv *service.ScheduableService
	depSrv *service.DependencyService
}

// Create a new session
func NewSession(id uint64) *Session {
	factory := repository.NewInMemoryRepositoryFactory()
	scheduableSrv := service.NewScheduableService(factory)
	depSrv := service.NewDependencyService(factory)

	return &Session{id: id, depSrv: depSrv, scheduableSrv: scheduableSrv}
}
